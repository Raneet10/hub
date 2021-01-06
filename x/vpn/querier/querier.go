package querier

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/deposit"
	"github.com/sentinel-official/hub/x/node"
	"github.com/sentinel-official/hub/x/plan"
	"github.com/sentinel-official/hub/x/provider"
	"github.com/sentinel-official/hub/x/session"
	"github.com/sentinel-official/hub/x/subscription"
	"github.com/sentinel-official/hub/x/vpn/keeper"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func NewQuerier(k keeper.Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case deposit.ModuleName:
			return deposit.Querier(ctx, path[1:], req, k.Deposit)
		case provider.ModuleName:
			return provider.Querier(ctx, path[1:], req, k.Provider)
		case node.ModuleName:
			return node.Querier(ctx, path[1:], req, k.Node)
		case plan.ModuleName:
			return plan.Querier(ctx, path[1:], req, k.Plan)
		case subscription.ModuleName:
			return subscription.Querier(ctx, path[1:], req, k.Subscription)
		case session.ModuleName:
			return session.Querier(ctx, path[1:], req, k.Session)
		default:
			return nil, errors.Wrapf(types.ErrorUnknownQueryType, "%s", path[0])
		}
	}
}
