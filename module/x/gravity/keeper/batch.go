package keeper

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/althea-net/cosmos-gravity-bridge/module/x/gravity/types"
)

const OutgoingTxBatchSize = 100
const BatchCreationPeriod = uint64(120)

// BuildOutgoingTXBatch starts the following process chain:
// - find bridged denominator for given voucher type
// - determine if an unexecuted batch is already waiting for this token type, if so confirm the new batch would
//   have a higher total fees. If not exit without creating a batch
// - select available transactions from the outgoing transaction pool sorted by fee desc
// - persist an outgoing batch object with an incrementing ID = nonce
// - emit an event
func (k Keeper) BuildOutgoingTXBatch(
	ctx sdk.Context,
	contract types.EthAddress,
	maxElements uint) (*types.InternalOutgoingTxBatch, error) {
	if maxElements == 0 {
		return nil, sdkerrors.Wrap(types.ErrInvalid, "max elements value")
	}

	lastBatch := k.GetLastOutgoingBatchByTokenType(ctx, contract)

	// lastBatch may be nil if there are no existing batches, we only need
	// to perform this check if a previous batch exists
	if lastBatch != nil {
		// this traverses the current tx pool for this token type and determines what
		// fees a hypothetical batch would have if created
		currentFees := k.GetBatchFeeByTokenType(ctx, contract, maxElements)
		if currentFees == nil {
			return nil, sdkerrors.Wrap(types.ErrInvalid, "error getting fees from tx pool")
		}

		lastFees := lastBatch.ToExternal().GetFees()
		if lastFees.GT(currentFees.TotalFees) {
			return nil, sdkerrors.Wrap(types.ErrInvalid, "new batch would not be more profitable")
		}
	}

	selectedTx, err := k.pickUnbatchedTX(ctx, contract, maxElements)

	if err != nil {
		return nil, err
	}

	if len(selectedTx) == 0 {
		return nil, sdkerrors.Wrap(types.ErrEmpty, "there are no unbatched transactions for that denom")
	}

	nextID := k.autoIncrementID(ctx, types.KeyLastOutgoingBatchID)
	batch, err := types.NewInternalOutgingTxBatch(nextID, k.getBatchTimeoutHeight(ctx), selectedTx, contract, 0)
	if err != nil {
		panic(sdkerrors.Wrap(err, "unable to create batch"))
	}
	k.StoreBatch(ctx, batch)

	// Get the checkpoint and store it as a legit past batch
	checkpoint := batch.GetCheckpoint(k.GetGravityID(ctx))
	k.SetPastEthSignatureCheckpoint(ctx, checkpoint)

	batchEvent := sdk.NewEvent(
		types.EventTypeOutgoingBatch,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		sdk.NewAttribute(types.AttributeKeyContract, k.GetBridgeContractAddress(ctx).GetAddress()),
		sdk.NewAttribute(types.AttributeKeyBridgeChainID, strconv.Itoa(int(k.GetBridgeChainID(ctx)))),
		sdk.NewAttribute(types.AttributeKeyOutgoingBatchID, fmt.Sprint(nextID)),
		sdk.NewAttribute(types.AttributeKeyNonce, fmt.Sprint(nextID)),
	)
	ctx.EventManager().EmitEvent(batchEvent)
	return batch, nil
}

// This gets the batch timeout height in Ethereum blocks.
func (k Keeper) getBatchTimeoutHeight(ctx sdk.Context) uint64 {
	params := k.GetParams(ctx)
	// currentCosmosHeight := ctx.BlockHeight()
	currentCosmosTimeMs := uint64(ctx.BlockTime().UnixNano() / 1000000)
	// we store the last observed Cosmos and Ethereum heights, we do not concern ourselves if these values are zero because
	// no batch can be produced if the last Ethereum block height is not first populated by a deposit event.
	heights := k.GetLastObservedEthereumBlockHeight(ctx)
	if heights.CosmosBlockHeight == 0 || heights.EthereumBlockHeight == 0 || heights.CosmosBlockTimeMs == 0 {
		return 0
	}
	// we project how long it has been in milliseconds since the last Ethereum block height was observed
	// projectedMillis := (uint64(currentCosmosHeight) - heights.CosmosBlockHeight) * params.AverageBlockTime

	// directlly calculate the difference between current time and the last cosmos block when Ethereum update occured
	realMillis := currentCosmosTimeMs - heights.CosmosBlockTimeMs
	// There is a delay between actual Ethereum event and its observation.
	// This results to the fact that heights.CosmosBlockHeight and heights.EthereumBlockHeight are not in the same time.
	// The difference between them is equal to the delay, which is usually less than 5 minutes.
	// Even if it is more than 5 minutes, the logic is save as long as the delay is lower than batch_timeout parameter, which by default is 12h.
	// In order to make it as precise as possible, we set 300000ms as a contants which purpose is to cover the usual delay between the two events.
	ethereumOffset := 300000 / params.AverageEthereumBlockTime
	// we convert that projection into the current Ethereum height using the average Ethereum block time in millis
	// projectedCurrentEthereumHeight := (projectedMillis / params.AverageEthereumBlockTime) + heights.EthereumBlockHeight
	projectedCurrentEthereumHeight := (realMillis / params.AverageEthereumBlockTime) + heights.EthereumBlockHeight + ethereumOffset
	// we convert our target time for block timeouts (lets say 12 hours) into a number of blocks to
	// place on top of our projection of the current Ethereum block height.
	blocksToAdd := params.TargetBatchTimeout / params.AverageEthereumBlockTime
	return projectedCurrentEthereumHeight + blocksToAdd
}

// OutgoingTxBatchExecuted is run when the Cosmos chain detects that a batch has been executed on Ethereum
// It frees all the transactions in the batch, then cancels all earlier batches, this function panics instead
// of returning errors because any failure will cause a double spend.
func (k Keeper) OutgoingTxBatchExecuted(ctx sdk.Context, tokenContract types.EthAddress, nonce uint64) {
	b := k.GetOutgoingTXBatch(ctx, tokenContract, nonce)
	if b == nil {
		panic(fmt.Sprintf("unknown batch nonce for outgoing tx batch %s %d", tokenContract, nonce))
	}

	// Iterate through remaining batches
	k.IterateOutgoingTXBatches(ctx, func(key []byte, iter_batch *types.InternalOutgoingTxBatch) bool {
		// If the iterated batches nonce is lower than the one that was just executed, cancel it
		if iter_batch.BatchNonce < b.BatchNonce && iter_batch.TokenContract.GetAddress() == tokenContract.GetAddress() {
			err := k.CancelOutgoingTXBatch(ctx, tokenContract, iter_batch.BatchNonce)
			if err != nil {
				panic(fmt.Sprintf("Failed cancel out batch %s %d while trying to execute %s %d with %s", tokenContract, iter_batch.BatchNonce, tokenContract, nonce, err))
			}
		}
		return false
	})

	// Delete batch since it is finished
	k.DeleteBatch(ctx, *b)

}

// StoreBatch stores a transaction batch
func (k Keeper) StoreBatch(ctx sdk.Context, batch *types.InternalOutgoingTxBatch) {
	if err := batch.ValidateBasic(); err != nil {
		panic(sdkerrors.Wrap(err, "attempted to store invalid batch"))
	}
	store := ctx.KVStore(k.storeKey)
	// set the current block height when storing the batch
	batch.Block = uint64(ctx.BlockHeight())
	key := types.GetOutgoingTxBatchKey(batch.TokenContract, batch.BatchNonce)
	store.Set(key, k.cdc.MustMarshal(batch.ToExternal()))

	blockKey := types.GetOutgoingTxBatchBlockKey(batch.Block)
	store.Set(blockKey, k.cdc.MustMarshal(batch.ToExternal()))
}

// StoreBatchUnsafe stores a transaction batch w/o setting the height
func (k Keeper) StoreBatchUnsafe(ctx sdk.Context, batch *types.InternalOutgoingTxBatch) {
	if err := batch.ValidateBasic(); err != nil {
		panic(sdkerrors.Wrap(err, "attempted to store invalid batch"))
	}
	batchExt := batch.ToExternal()
	store := ctx.KVStore(k.storeKey)
	key := types.GetOutgoingTxBatchKey(batch.TokenContract, batchExt.BatchNonce)
	store.Set(key, k.cdc.MustMarshal(batchExt))

	blockKey := types.GetOutgoingTxBatchBlockKey(batchExt.Block)
	store.Set(blockKey, k.cdc.MustMarshal(batchExt))
}

// DeleteBatch deletes an outgoing transaction batch
func (k Keeper) DeleteBatch(ctx sdk.Context, batch types.InternalOutgoingTxBatch) {
	if err := batch.ValidateBasic(); err != nil {
		panic(sdkerrors.Wrap(err, "attempted to delete invalid batch"))
	}
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetOutgoingTxBatchKey(batch.TokenContract, batch.BatchNonce))
	store.Delete(types.GetOutgoingTxBatchBlockKey(batch.Block))
}

// pickUnbatchedTX find TX in pool and remove from "available" second index
func (k Keeper) pickUnbatchedTX(
	ctx sdk.Context,
	contractAddress types.EthAddress,
	maxElements uint) ([]*types.InternalOutgoingTransferTx, error) {
	var selectedTx []*types.InternalOutgoingTransferTx
	var err error
	k.IterateUnbatchedTransactionsByContract(ctx, contractAddress, func(_ []byte, tx *types.InternalOutgoingTransferTx) bool {
		if tx != nil && tx.Erc20Fee != nil {
			selectedTx = append(selectedTx, tx)
			err = k.removeUnbatchedTX(ctx, *tx.Erc20Fee, tx.Id)
			oldTx, oldTxErr := k.GetUnbatchedTxByFeeAndId(ctx, *tx.Erc20Fee, tx.Id)
			if oldTx != nil || oldTxErr == nil {
				panic("picked a duplicate transaction from the pool, duplicates should never exist!")
			}
			return err != nil || uint(len(selectedTx)) == maxElements
		} else {
			panic("tx and fee should never be nil!")
		}
	})
	return selectedTx, err
}

// GetOutgoingTXBatch loads a batch object. Returns nil when not exists.
func (k Keeper) GetOutgoingTXBatch(ctx sdk.Context, tokenContract types.EthAddress, nonce uint64) *types.InternalOutgoingTxBatch {
	store := ctx.KVStore(k.storeKey)
	key := types.GetOutgoingTxBatchKey(tokenContract, nonce)
	bz := store.Get(key)
	if len(bz) == 0 {
		return nil
	}
	var b types.OutgoingTxBatch
	k.cdc.MustUnmarshal(bz, &b)
	for _, tx := range b.Transactions {
		tx.Erc20Token.Contract = tokenContract.GetAddress()
		tx.Erc20Fee.Contract = tokenContract.GetAddress()
	}
	ret, err := b.ToInternal()
	if err != nil {
		panic(sdkerrors.Wrap(err, "found invalid batch in store"))
	}
	return ret
}

// CancelOutgoingTXBatch releases all TX in the batch and deletes the batch
func (k Keeper) CancelOutgoingTXBatch(ctx sdk.Context, tokenContract types.EthAddress, nonce uint64) error {
	batch := k.GetOutgoingTXBatch(ctx, tokenContract, nonce)
	if batch == nil {
		return types.ErrUnknown
	}
	for _, tx := range batch.Transactions {
		err := k.addUnbatchedTX(ctx, tx)
		if err != nil {
			panic(sdkerrors.Wrapf(err, "unable to add batched transaction back into pool %v", tx))
		}
	}

	// Delete batch since it is finished
	k.DeleteBatch(ctx, *batch)

	batchEvent := sdk.NewEvent(
		types.EventTypeOutgoingBatchCanceled,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		sdk.NewAttribute(types.AttributeKeyContract, k.GetBridgeContractAddress(ctx).GetAddress()),
		sdk.NewAttribute(types.AttributeKeyBridgeChainID, strconv.Itoa(int(k.GetBridgeChainID(ctx)))),
		sdk.NewAttribute(types.AttributeKeyOutgoingBatchID, fmt.Sprint(nonce)),
		sdk.NewAttribute(types.AttributeKeyNonce, fmt.Sprint(nonce)),
	)
	ctx.EventManager().EmitEvent(batchEvent)
	return nil
}

// IterateOutgoingTXBatches iterates through all outgoing batches in DESC order.
func (k Keeper) IterateOutgoingTXBatches(ctx sdk.Context, cb func(key []byte, batch *types.InternalOutgoingTxBatch) bool) {
	prefixStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.OutgoingTXBatchKey)
	iter := prefixStore.ReverseIterator(nil, nil)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var batch types.OutgoingTxBatch
		k.cdc.MustUnmarshal(iter.Value(), &batch)
		intBatch, err := batch.ToInternal()
		if err != nil {
			panic(sdkerrors.Wrap(err, "found invalid batch in store"))
		}
		// cb returns true to stop early
		if cb(iter.Key(), intBatch) {
			break
		}
	}
}

// GetOutgoingTxBatches returns the outgoing tx batches
func (k Keeper) GetOutgoingTxBatches(ctx sdk.Context) (out []*types.InternalOutgoingTxBatch) {
	k.IterateOutgoingTXBatches(ctx, func(_ []byte, batch *types.InternalOutgoingTxBatch) bool {
		out = append(out, batch)
		return false
	})
	return
}

// GetLastOutgoingBatchByTokenType gets the latest outgoing tx batch by token type
func (k Keeper) GetLastOutgoingBatchByTokenType(ctx sdk.Context, token types.EthAddress) *types.InternalOutgoingTxBatch {
	batches := k.GetOutgoingTxBatches(ctx)
	var lastBatch *types.InternalOutgoingTxBatch = nil
	lastNonce := uint64(0)
	for _, batch := range batches {
		if batch.TokenContract.GetAddress() == token.GetAddress() && batch.BatchNonce > lastNonce {
			lastBatch = batch
			lastNonce = batch.BatchNonce
		}
	}
	return lastBatch
}

// SetLastSlashedBatchBlock sets the latest slashed Batch block height
func (k Keeper) SetLastSlashedBatchBlock(ctx sdk.Context, blockHeight uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.LastSlashedBatchBlock, types.UInt64Bytes(blockHeight))
}

// GetLastSlashedBatchBlock returns the latest slashed Batch block
func (k Keeper) GetLastSlashedBatchBlock(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bytes := store.Get(types.LastSlashedBatchBlock)

	if len(bytes) == 0 {
		return 0
	}
	return types.UInt64FromBytes(bytes)
}

// GetUnSlashedBatches returns all the unslashed batches in state
func (k Keeper) GetUnSlashedBatches(ctx sdk.Context, maxHeight uint64) (out []*types.InternalOutgoingTxBatch) {
	lastSlashedBatchBlock := k.GetLastSlashedBatchBlock(ctx)
	k.IterateBatchBySlashedBatchBlock(ctx,
		lastSlashedBatchBlock,
		maxHeight,
		func(_ []byte, batch *types.InternalOutgoingTxBatch) bool {
			if batch.Block > lastSlashedBatchBlock {
				out = append(out, batch)
			}
			return false
		})
	return
}

// IterateBatchBySlashedBatchBlock iterates through all Batch by last slashed Batch block in ASC order
func (k Keeper) IterateBatchBySlashedBatchBlock(
	ctx sdk.Context,
	lastSlashedBatchBlock uint64,
	maxHeight uint64,
	cb func([]byte, *types.InternalOutgoingTxBatch) bool) {
	prefixStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.OutgoingTXBatchBlockKey)
	iter := prefixStore.Iterator(types.UInt64Bytes(lastSlashedBatchBlock), types.UInt64Bytes(maxHeight))
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var batch types.OutgoingTxBatch
		k.cdc.MustUnmarshal(iter.Value(), &batch)
		intBatch, err := batch.ToInternal()
		if err != nil {
			panic(sdkerrors.Wrap(err, "found invalid batch in store"))
		}

		// cb returns true to stop early
		if cb(iter.Key(), intBatch) {
			break
		}
	}
}
