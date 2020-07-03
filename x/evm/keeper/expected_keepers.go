package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"
)

type AccountKeeper interface {
	NewAccountWithAddress(ctx sdk.Context, addr sdk.AccAddress) exported.Account
	RemoveAccount(ctx sdk.Context, acc exported.Account)
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) exported.Account
	SetAccount(ctx sdk.Context, acc exported.Account)
}

type BankKeeper interface {
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
}
