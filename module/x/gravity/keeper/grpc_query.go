package keeper

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/althea-net/cosmos-gravity-bridge/module/x/gravity/types"
)

// var _ types.QueryServer = Keeper{
// 	StakingKeeper:      nil,
// 	storeKey:           nil,
// 	paramSpace:         paramstypes.Subspace{},
// 	cdc:                nil,
// 	BankKeeper:         nil,
// 	SlashingKeeper:     nil,
// 	AccountKeeper:      nil,
// 	AttestationHandler: nil,
// }

var _ types.QueryServer = Querier{}

type Querier struct {
	Keeper
}

func NewQuerier(keeper Keeper) Querier {
	return Querier{Keeper: keeper}
}

const QUERY_ATTESTATIONS_LIMIT uint64 = 1000
const maxValsetRequestsReturned = 5
const MaxResults = 100 // todo: impl pagination

// Params queries the params of the gravity module
func (k Querier) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	var params types.Params
	k.paramSpace.GetParamSet(sdk.UnwrapSDKContext(c), &params)
	return &types.QueryParamsResponse{Params: params}, nil

}

// CurrentValset queries the CurrentValset of the gravity module
func (k Querier) CurrentValset(
	c context.Context,
	req *types.QueryCurrentValsetRequest) (*types.QueryCurrentValsetResponse, error) {
	return &types.QueryCurrentValsetResponse{Valset: k.GetCurrentValset(sdk.UnwrapSDKContext(c))}, nil
}

// ValsetRequest queries the ValsetRequest of the gravity module
func (k Querier) ValsetRequest(
	c context.Context,
	req *types.QueryValsetRequestRequest) (*types.QueryValsetRequestResponse, error) {
	return &types.QueryValsetRequestResponse{Valset: k.GetValset(sdk.UnwrapSDKContext(c), req.Nonce)}, nil
}

// ValsetConfirm queries the ValsetConfirm of the gravity module
func (k Querier) ValsetConfirm(
	c context.Context,
	req *types.QueryValsetConfirmRequest) (*types.QueryValsetConfirmResponse, error) {
	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "address invalid")
	}
	return &types.QueryValsetConfirmResponse{Confirm: k.GetValsetConfirm(sdk.UnwrapSDKContext(c), req.Nonce, addr)}, nil
}

// ValsetConfirmsByNonce queries the ValsetConfirmsByNonce of the gravity module
func (k Querier) ValsetConfirmsByNonce(
	c context.Context,
	req *types.QueryValsetConfirmsByNonceRequest) (*types.QueryValsetConfirmsByNonceResponse, error) {
	var confirms []*types.MsgValsetConfirm
	k.IterateValsetConfirmByNonce(sdk.UnwrapSDKContext(c), req.Nonce, func(_ []byte, c types.MsgValsetConfirm) bool {
		confirms = append(confirms, &c)
		return false
	})
	return &types.QueryValsetConfirmsByNonceResponse{Confirms: confirms}, nil
}

// LastValsetRequests queries the LastValsetRequests of the gravity module
func (k Querier) LastValsetRequests(
	c context.Context,
	req *types.QueryLastValsetRequestsRequest) (*types.QueryLastValsetRequestsResponse, error) {
	valReq := k.GetValsets(sdk.UnwrapSDKContext(c))
	valReqLen := len(valReq)
	retLen := 0
	if valReqLen < maxValsetRequestsReturned {
		retLen = valReqLen
	} else {
		retLen = maxValsetRequestsReturned
	}
	return &types.QueryLastValsetRequestsResponse{Valsets: valReq[0:retLen]}, nil
}

// LastPendingValsetRequestByAddr queries the LastPendingValsetRequestByAddr of the gravity module
func (k Querier) LastPendingValsetRequestByAddr(
	c context.Context,
	req *types.QueryLastPendingValsetRequestByAddrRequest) (*types.QueryLastPendingValsetRequestByAddrResponse, error) {
	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "address invalid")
	}

	var pendingValsetReq []*types.Valset
	k.IterateValsets(sdk.UnwrapSDKContext(c), func(_ []byte, val *types.Valset) bool {
		// foundConfirm is true if the operatorAddr has signed the valset we are currently looking at
		foundConfirm := k.GetValsetConfirm(sdk.UnwrapSDKContext(c), val.Nonce, addr) != nil
		// if this valset has NOT been signed by operatorAddr, store it in pendingValsetReq
		// and exit the loop
		if !foundConfirm {
			pendingValsetReq = append(pendingValsetReq, val)
		}
		// if we have more than 100 unconfirmed requests in
		// our array we should exit, TODO pagination
		if len(pendingValsetReq) > 100 {
			return true
		}
		// return false to continue the loop
		return false
	})
	return &types.QueryLastPendingValsetRequestByAddrResponse{Valsets: pendingValsetReq}, nil
}

// BatchFees queries the batch fees from unbatched pool
func (k Querier) BatchFees(
	c context.Context,
	req *types.QueryBatchFeeRequest) (*types.QueryBatchFeeResponse, error) {
	return &types.QueryBatchFeeResponse{BatchFees: k.GetAllBatchFees(sdk.UnwrapSDKContext(c), OutgoingTxBatchSize)}, nil
}

// LastPendingBatchRequestByAddr queries the LastPendingBatchRequestByAddr of the gravity module
func (k Querier) LastPendingBatchRequestByAddr(
	c context.Context,
	req *types.QueryLastPendingBatchRequestByAddrRequest) (*types.QueryLastPendingBatchRequestByAddrResponse, error) {
	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "address invalid")
	}

	var pendingBatchReq *types.InternalOutgoingTxBatch
	k.IterateOutgoingTXBatches(sdk.UnwrapSDKContext(c), func(_ []byte, batch *types.InternalOutgoingTxBatch) bool {
		foundConfirm := k.GetBatchConfirm(sdk.UnwrapSDKContext(c), batch.BatchNonce, batch.TokenContract, addr) != nil
		if !foundConfirm {
			pendingBatchReq = batch
			return true
		}
		return false
	})

	return &types.QueryLastPendingBatchRequestByAddrResponse{Batch: pendingBatchReq.ToExternal()}, nil
}

func (k Querier) LastPendingLogicCallByAddr(
	c context.Context,
	req *types.QueryLastPendingLogicCallByAddrRequest) (*types.QueryLastPendingLogicCallByAddrResponse, error) {
	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "address invalid")
	}

	var pendingLogicReq *types.OutgoingLogicCall
	k.IterateOutgoingLogicCalls(sdk.UnwrapSDKContext(c), func(_ []byte, logic *types.OutgoingLogicCall) bool {
		foundConfirm := k.GetLogicCallConfirm(sdk.UnwrapSDKContext(c),
			logic.InvalidationId, logic.InvalidationNonce, addr) != nil
		if !foundConfirm {
			pendingLogicReq = logic
			return true
		}
		return false
	})
	return &types.QueryLastPendingLogicCallByAddrResponse{Call: pendingLogicReq}, nil
}

// OutgoingTxBatches queries the OutgoingTxBatches of the gravity module
func (k Querier) OutgoingTxBatches(
	c context.Context,
	req *types.QueryOutgoingTxBatchesRequest) (*types.QueryOutgoingTxBatchesResponse, error) {
	var batches []*types.OutgoingTxBatch
	k.IterateOutgoingTXBatches(sdk.UnwrapSDKContext(c), func(_ []byte, batch *types.InternalOutgoingTxBatch) bool {
		batches = append(batches, batch.ToExternal())
		return len(batches) == MaxResults
	})
	return &types.QueryOutgoingTxBatchesResponse{Batches: batches}, nil
}

// OutgoingLogicCalls queries the OutgoingLogicCalls of the gravity module
func (k Querier) OutgoingLogicCalls(
	c context.Context,
	req *types.QueryOutgoingLogicCallsRequest) (*types.QueryOutgoingLogicCallsResponse, error) {
	var calls []*types.OutgoingLogicCall
	k.IterateOutgoingLogicCalls(sdk.UnwrapSDKContext(c), func(_ []byte, call *types.OutgoingLogicCall) bool {
		calls = append(calls, call)
		return len(calls) == MaxResults
	})
	return &types.QueryOutgoingLogicCallsResponse{Calls: calls}, nil
}

// BatchRequestByNonce queries the BatchRequestByNonce of the gravity module
func (k Querier) BatchRequestByNonce(
	c context.Context,
	req *types.QueryBatchRequestByNonceRequest) (*types.QueryBatchRequestByNonceResponse, error) {
	addr, err := types.NewEthAddress(req.ContractAddress)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, err.Error())
	}
	foundBatch := k.GetOutgoingTXBatch(sdk.UnwrapSDKContext(c), *addr, req.Nonce)
	if foundBatch == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Can not find tx batch")
	}
	return &types.QueryBatchRequestByNonceResponse{Batch: foundBatch.ToExternal()}, nil
}

// BatchConfirms returns the batch confirmations by nonce and token contract
func (k Querier) BatchConfirms(
	c context.Context,
	req *types.QueryBatchConfirmsRequest) (*types.QueryBatchConfirmsResponse, error) {
	var confirms []*types.MsgConfirmBatch
	contract, err := types.NewEthAddress(req.ContractAddress)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "invalid contract address in request")
	}
	k.IterateBatchConfirmByNonceAndTokenContract(sdk.UnwrapSDKContext(c),
		req.Nonce, *contract, func(_ []byte, c types.MsgConfirmBatch) bool {
			confirms = append(confirms, &c)
			return false
		})
	return &types.QueryBatchConfirmsResponse{Confirms: confirms}, nil
}

// LogicConfirms returns the Logic confirmations by nonce and token contract
func (k Querier) LogicConfirms(
	c context.Context,
	req *types.QueryLogicConfirmsRequest) (*types.QueryLogicConfirmsResponse, error) {
	var confirms []*types.MsgConfirmLogicCall
	k.IterateLogicConfirmByInvalidationIDAndNonce(sdk.UnwrapSDKContext(c), req.InvalidationId,
		req.InvalidationNonce, func(_ []byte, c *types.MsgConfirmLogicCall) bool {
			confirms = append(confirms, c)
			return false
		})

	return &types.QueryLogicConfirmsResponse{Confirms: confirms}, nil
}

// LastEventNonceByAddr returns the last event nonce for the given validator address,
// this allows eth oracles to figure out where they left off
func (k Querier) LastEventNonceByAddr(
	c context.Context,
	req *types.QueryLastEventNonceByAddrRequest) (*types.QueryLastEventNonceByAddrResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	var ret types.QueryLastEventNonceByAddrResponse
	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, req.Address)
	}
	validator, found := k.GetOrchestratorValidator(ctx, addr)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrUnknown, "address")
	}
	lastEventNonce := k.GetLastEventNonceByValidator(ctx, validator.GetOperator())
	ret.EventNonce = lastEventNonce
	return &ret, nil
}

// DenomToERC20 queries the Cosmos Denom that maps to an Ethereum ERC20
func (k Querier) DenomToERC20(
	c context.Context,
	req *types.QueryDenomToERC20Request) (*types.QueryDenomToERC20Response, error) {
	ctx := sdk.UnwrapSDKContext(c)
	cosmosOriginated, erc20, err := k.DenomToERC20Lookup(ctx, req.Denom)
	var ret types.QueryDenomToERC20Response
	ret.Erc20 = erc20.GetAddress()
	ret.CosmosOriginated = cosmosOriginated

	return &ret, err
}

// ERC20ToDenom queries the ERC20 contract that maps to an Ethereum ERC20 if any
func (k Querier) ERC20ToDenom(
	c context.Context,
	req *types.QueryERC20ToDenomRequest) (*types.QueryERC20ToDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	ethAddr, err := types.NewEthAddress(req.Erc20)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "invalid Erc20 in request: %s", req.Erc20)
	}
	cosmosOriginated, name := k.ERC20ToDenomLookup(ctx, *ethAddr)
	var ret types.QueryERC20ToDenomResponse
	ret.Denom = name
	ret.CosmosOriginated = cosmosOriginated

	return &ret, nil
}

// GetAttestations queries the attestation map
func (k Querier) GetAttestations(
	c context.Context,
	req *types.QueryAttestationsRequest) (*types.QueryAttestationsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	limit := req.Limit
	if limit > QUERY_ATTESTATIONS_LIMIT {
		limit = QUERY_ATTESTATIONS_LIMIT
	}
	attestations := k.GetMostRecentAttestations(ctx, limit)

	return &types.QueryAttestationsResponse{Attestations: attestations}, nil
}

func (k Querier) GetDelegateKeyByValidator(
	c context.Context,
	req *types.QueryDelegateKeysByValidatorAddress) (*types.QueryDelegateKeysByValidatorAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	keys := k.GetDelegateKeys(ctx)
	reqValidator, err := sdk.ValAddressFromBech32(req.ValidatorAddress)
	if err != nil {
		return nil, err
	}
	for _, key := range keys {
		keyValidator, err := sdk.ValAddressFromBech32(key.Validator)
		// this should be impossible due to the validate basic on the set orchestrator message
		if err != nil {
			panic("Invalid validator addr in store!")
		}
		if reqValidator.Equals(keyValidator) {
			return &types.QueryDelegateKeysByValidatorAddressResponse{EthAddress: key.EthAddress, OrchestratorAddress: key.Orchestrator}, nil
		}
	}

	return nil, sdkerrors.Wrap(types.ErrInvalid, "No validator")
}

func (k Querier) GetDelegateKeyByOrchestrator(
	c context.Context,
	req *types.QueryDelegateKeysByOrchestratorAddress) (*types.QueryDelegateKeysByOrchestratorAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	keys := k.GetDelegateKeys(ctx)
	reqOrchestrator, err := sdk.AccAddressFromBech32(req.OrchestratorAddress)
	if err != nil {
		return nil, err
	}
	for _, key := range keys {
		keyOrchestrator, err := sdk.AccAddressFromBech32(key.Orchestrator)
		// this should be impossible due to the validate basic on the set orchestrator message
		if err != nil {
			panic("Invalid orchestrator addr in store!")
		}
		if reqOrchestrator.Equals(keyOrchestrator) {
			return &types.QueryDelegateKeysByOrchestratorAddressResponse{ValidatorAddress: key.Validator, EthAddress: key.EthAddress}, nil
		}

	}
	return nil, sdkerrors.Wrap(types.ErrInvalid, "No validator")
}

func (k Querier) GetDelegateKeyByEth(
	c context.Context,
	req *types.QueryDelegateKeysByEthAddress) (*types.QueryDelegateKeysByEthAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	keys := k.GetDelegateKeys(ctx)
	if err := types.ValidateEthAddress(req.EthAddress); err != nil {
		return nil, sdkerrors.Wrap(err, "invalid eth address")
	}
	for _, key := range keys {
		if strings.EqualFold(req.EthAddress, key.EthAddress) {
			return &types.QueryDelegateKeysByEthAddressResponse{
				ValidatorAddress:    key.Validator,
				OrchestratorAddress: key.Orchestrator,
			}, nil
		}
	}

	return nil, sdkerrors.Wrap(types.ErrInvalid, "No validator")
}

func (k Querier) GetPendingSendToEth(
	c context.Context,
	req *types.QueryPendingSendToEth) (*types.QueryPendingSendToEthResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	batches := k.GetOutgoingTxBatches(ctx)
	unbatched_tx := k.GetUnbatchedTransactions(ctx)
	sender_address := req.GetSenderAddress()
	res := types.QueryPendingSendToEthResponse{
		TransfersInBatches: []*types.OutgoingTransferTx{},
		UnbatchedTransfers: []*types.OutgoingTransferTx{},
	}
	for _, batch := range batches {
		for _, tx := range batch.Transactions {
			if sender_address == "" || tx.Sender.String() == sender_address {
				res.TransfersInBatches = append(res.TransfersInBatches, tx.ToExternal())
			}
		}
	}
	for _, tx := range unbatched_tx {
		if sender_address == "" || tx.Sender.String() == sender_address {
			res.UnbatchedTransfers = append(res.UnbatchedTransfers, tx.ToExternal())
		}
	}

	return &res, nil
}
