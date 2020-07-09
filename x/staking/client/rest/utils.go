package rest

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/okex/okchain/x/common"
	"net/http"

	"github.com/gorilla/mux"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/okex/okchain/x/staking/types"
)

// contains checks if the a given query contains one of the tx types
func contains(stringSlice []string, txType string) bool {
	for _, word := range stringSlice {
		if word == txType {
			return true
		}
	}
	return false
}

// queries staking txs
func queryTxs(cliCtx client.Context, action string, delegatorAddr string) (*sdk.SearchTxsResult, error) {
	page := 1
	limit := 100
	events := []string{
		fmt.Sprintf("%s.%s='%s'", sdk.EventTypeMessage, sdk.AttributeKeyAction, action),
		fmt.Sprintf("%s.%s='%s'", sdk.EventTypeMessage, sdk.AttributeKeySender, delegatorAddr),
	}

	return authclient.QueryTxsByEvents(cliCtx, events, page, limit, "")
}

func queryDelegator(cliCtx client.Context, endpoint string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bech32DelAddr := mux.Vars(r)["delegatorAddr"]

		delegatorAddr, err := sdk.AccAddressFromBech32(bech32DelAddr)
		if err != nil {
			common.HandleErrorResponseV2(w, http.StatusBadRequest, common.ErrorInvalidDelegatorAddress)
			return
		}

		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		params := types.NewQueryDelegatorParams(delegatorAddr)

		bz, err := cliCtx.Codec.MarshalJSON(params)
		if err != nil {
			common.HandleErrorResponseV2(w, http.StatusBadRequest, common.ErrorCodecFails)
			return
		}

		res, height, err := cliCtx.QueryWithData(endpoint, bz)
		if err != nil {
			common.HandleErrorResponseV2(w, http.StatusInternalServerError, common.ErrorABCIQueryFails)
			return
		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func queryValidator(cliCtx client.Context, endpoint string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bech32ValAddr := mux.Vars(r)["validatorAddr"]

		validatorAddr, err := sdk.ValAddressFromBech32(bech32ValAddr)
		if err != nil {
			common.HandleErrorResponseV2(w, http.StatusBadRequest, common.ErrorInvalidValidatorAddress)
			return
		}

		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		params := types.NewQueryValidatorParams(validatorAddr)

		bz, err := cliCtx.Codec.MarshalJSON(params)
		if err != nil {
			common.HandleErrorResponseV2(w, http.StatusBadRequest, common.ErrorCodecFails)
			return
		}

		res, height, err := cliCtx.QueryWithData(endpoint, bz)
		if err != nil {
			common.HandleErrorResponseV2(w, http.StatusInternalServerError, common.ErrorABCIQueryFails)
			return
		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, res)
	}
}
