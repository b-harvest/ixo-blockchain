package fees

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/fees/internal/keeper"
	abci "github.com/tendermint/tendermint/abci/types"
)

func EndBlocker(ctx sdk.Context, keeper keeper.Keeper) []abci.ValidatorUpdate {

	iterator := keeper.GetSubscriptionIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		subscription := keeper.MustGetSubscriptionByKey(ctx, iterator.Key())
		subContent := subscription.Content

		// Continue if hasn't started or no charge necessary
		if !subContent.Started(ctx) || !subContent.ShouldCharge(ctx) {
			continue
		}

		// Charge fee
		success := true // TODO: Charge fee
		subContent.NextPeriod(success)

		// Note: if fee can be re-charged immediately, this should be done in
		// the next block to prevent spending too much time charging fees

		// Delete subscription if it has ended and no more charges
		if subContent.Ended() && !subContent.ShouldCharge(ctx) {
			// TODO: delete subscription
		}

		// Save subscription
		keeper.SetSubscription(ctx, subscription)
	}
	return []abci.ValidatorUpdate{}
}