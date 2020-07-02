package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/okex/okchain/hexutil"
)

const (
	TypeMsgContract = "contract"
)

var (
	_ sdk.Msg = &MsgContract{}
)

type MsgContract struct {
	From    sdk.AccAddress `json:"from" yaml:"from"`
	To      sdk.AccAddress `json:"to" yaml:"to"`
	Payload hexutil.Bytes  `json:"payload" yaml:"payload"`
	Amount  sdk.Coin       `json:"amount" yaml:"amount"`
}

func (msg MsgContract) Route() string {
	return RouterKey
}

func (msg MsgContract) Type() string {
	return TypeMsgContract
}

func (msg MsgContract) ValidateBasic() sdk.Error {
	if msg.From.Empty() {
		return sdk.ErrInvalidAddress("failed to check send msg because miss sender address")
	}
	if !msg.Amount.IsValid() {
		return sdk.ErrInvalidCoins("failed to check msg because send amount is invalid: " + msg.Amount.String())
	}
	if msg.Amount.IsLT(sdk.ZeroFee()) {
		return sdk.ErrInsufficientCoins("failed to check msg because send amount must be positive")
	}
	if msg.Amount.Denom != sdk.DefaultBondDenom {
		return sdk.ErrInvalidCoins("coins of amount is invalid: " + msg.Amount.String())
	}
	if len(msg.Payload) == 0 {
		return ErrNoPayload("")
	}

	return nil
}

func (msg MsgContract) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgContract) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

func NewMsgContract(from, to sdk.AccAddress, payload []byte, amount sdk.Coin) MsgContract {
	return MsgContract{
		From:    from,
		To:      to,
		Payload: payload,
		Amount:  amount,
	}
}

type MsgContractQuery MsgContract

func NewMsgContractQuery(from, to sdk.AccAddress, payload []byte, amount sdk.Coin) MsgContractQuery {
	return MsgContractQuery{
		From:    from,
		To:      to,
		Payload: payload,
		Amount:  amount,
	}
}
