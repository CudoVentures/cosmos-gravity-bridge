package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"

	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	authzcodec "github.com/cosmos/cosmos-sdk/x/authz/codec"
	govcodec "github.com/cosmos/cosmos-sdk/x/gov/codec"
	groupcodec "github.com/cosmos/cosmos-sdk/x/group/codec"
)

// ModuleCdc is the codec for the module
var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

func init() {
	RegisterLegacyAminoCodec(Amino)
	cryptocodec.RegisterCrypto(Amino)
	sdk.RegisterLegacyAminoCodec(Amino)

	// Register all Amino interfaces and concrete types on the authz and gov Amino codec so that this can later be
	// used to properly serialize MsgGrant, MsgExec and MsgSubmitProposal instances
	RegisterLegacyAminoCodec(authzcodec.Amino)
	RegisterLegacyAminoCodec(govcodec.Amino)
	RegisterLegacyAminoCodec(groupcodec.Amino)
}

// nolint: exhaustivestruct
// RegisterInterfaces registers the interfaces for the proto stuff
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgValsetConfirm{},
		&MsgSendToEth{},
		&MsgRequestBatch{},
		&MsgSetMinFeeTransferToEth{},
		&MsgConfirmBatch{},
		&MsgConfirmLogicCall{},
		&MsgSendToCosmosClaim{},
		&MsgBatchSendToEthClaim{},
		&MsgERC20DeployedClaim{},
		&MsgSetOrchestratorAddress{},
		&MsgLogicCallExecutedClaim{},
		&MsgValsetUpdatedClaim{},
		&MsgCancelSendToEth{},
		&MsgSubmitBadSignatureEvidence{},
	)

	registry.RegisterInterface(
		"gravity.v1beta1.EthereumClaim",
		(*EthereumClaim)(nil),
		&MsgSendToCosmosClaim{},
		&MsgBatchSendToEthClaim{},
		&MsgERC20DeployedClaim{},
		&MsgLogicCallExecutedClaim{},
		&MsgValsetUpdatedClaim{},
	)

	registry.RegisterInterface("gravity.v1beta1.EthereumSigned", (*EthereumSigned)(nil), &Valset{}, &OutgoingTxBatch{}, &OutgoingLogicCall{})

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

// nolint: exhaustivestruct
// RegisterCodec registers concrete types on the Amino codec
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*EthereumClaim)(nil), nil)
	cdc.RegisterConcrete(&MsgSetOrchestratorAddress{}, "gravity/MsgSetOrchestratorAddress", nil)
	cdc.RegisterConcrete(&MsgValsetConfirm{}, "gravity/MsgValsetConfirm", nil)
	cdc.RegisterConcrete(&MsgSendToEth{}, "gravity/MsgSendToEth", nil)
	cdc.RegisterConcrete(&MsgRequestBatch{}, "gravity/MsgRequestBatch", nil)
	cdc.RegisterConcrete(&MsgSetMinFeeTransferToEth{}, "gravity/MsgSetMinFeeTransferToEth", nil)
	cdc.RegisterConcrete(&MsgConfirmBatch{}, "gravity/MsgConfirmBatch", nil)
	cdc.RegisterConcrete(&MsgConfirmLogicCall{}, "gravity/MsgConfirmLogicCall", nil)
	cdc.RegisterConcrete(&Valset{}, "gravity/Valset", nil)
	cdc.RegisterConcrete(&MsgSendToCosmosClaim{}, "gravity/MsgSendToCosmosClaim", nil)
	cdc.RegisterConcrete(&MsgBatchSendToEthClaim{}, "gravity/MsgBatchSendToEthClaim", nil)
	cdc.RegisterConcrete(&MsgERC20DeployedClaim{}, "gravity/MsgERC20DeployedClaim", nil)
	cdc.RegisterConcrete(&MsgLogicCallExecutedClaim{}, "gravity/MsgLogicCallExecutedClaim", nil)
	cdc.RegisterConcrete(&MsgValsetUpdatedClaim{}, "gravity/MsgValsetUpdatedClaim", nil)
	cdc.RegisterConcrete(&OutgoingTxBatch{}, "gravity/OutgoingTxBatch", nil)
	cdc.RegisterConcrete(&MsgCancelSendToEth{}, "gravity/MsgCancelSendToEth", nil)
	cdc.RegisterConcrete(&OutgoingTransferTx{}, "gravity/OutgoingTransferTx", nil)
	cdc.RegisterConcrete(&ERC20Token{}, "gravity/ERC20Token", nil)
	cdc.RegisterConcrete(&IDSet{}, "gravity/IDSet", nil)
	cdc.RegisterConcrete(&Attestation{}, "gravity/Attestation", nil)
	cdc.RegisterConcrete(&MsgSubmitBadSignatureEvidence{}, "gravity/MsgSubmitBadSignatureEvidence", nil)
}
