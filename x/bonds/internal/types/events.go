package types

const (
	EventTypeCreateBond   = "create_bond"
	EventTypeEditBond     = "edit_bond"
	EventTypeInitSwapper  = "init_swapper"
	EventTypeBuy          = "buy"
	EventTypeSell         = "sell"
	EventTypeSwap         = "swap"
	EventTypeOrderCancel  = "order_cancel"
	EventTypeOrderFulfill = "order_fulfill"

	AttributeKeyBondDid                = "bond_did"
	AttributeKeyToken                  = "token"
	AttributeKeyName                   = "name"
	AttributeKeyDescription            = "description"
	AttributeKeyFunctionType           = "function_type"
	AttributeKeyFunctionParameters     = "function_parameters"
	AttributeKeyReserveTokens          = "reserve_tokens"
	AttributeKeyReserveAddress         = "reserve_address"
	AttributeKeyTxFeePercentage        = "tx_fee_percentage"
	AttributeKeyExitFeePercentage      = "exit_fee_percentage"
	AttributeKeyFeeAddress             = "fee_address"
	AttributeKeyMaxSupply              = "max_supply"
	AttributeKeyOrderQuantityLimits    = "order_quantity_limits"
	AttributeKeySanityRate             = "sanity_rate"
	AttributeKeySanityMarginPercentage = "sanity_margin_percentage"
	AttributeKeyAllowSells             = "allow_sells"
	AttributeKeyBatchBlocks            = "batch_blocks"
	AttributeKeyMaxPrices              = "max_prices"
	AttributeKeySwapFromToken          = "from_token"
	AttributeKeySwapToToken            = "to_token"
	AttributeKeyOrderType              = "order_type"
	AttributeKeyAddress                = "address"
	AttributeKeyCancelReason           = "cancel_reason"
	AttributeKeyTokensMinted           = "tokens_minted"
	AttributeKeyTokensBurned           = "tokens_burned"
	AttributeKeyTokensSwapped          = "tokens_swapped"
	AttributeKeyChargedPrices          = "charged_prices"
	AttributeKeyChargedFees            = "charged_fees"
	AttributeKeyReturnedToAddress      = "returned_to_address"
	AttributeKeyNewBondTokenBalance    = "new_bond_token_balance"

	AttributeValueBuyOrder  = "buy"
	AttributeValueSellOrder = "sell"
	AttributeValueSwapOrder = "swap"
	AttributeValueCategory  = ModuleName
)
