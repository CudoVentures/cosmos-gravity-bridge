package simulation

import (
	"bytes"

	"github.com/althea-net/cosmos-gravity-bridge/module/x/gravity/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	gethcommon "github.com/ethereum/go-ethereum/common"
)

// RandomizedGenState generates a random GenesisState for simulations
func RandomizedGenState(simState *module.SimulationState) {
	stakingGenStateBz := simState.GenState[stakingtypes.ModuleName]
	stakingGenState := stakingtypes.GenesisState{}
	simState.Cdc.MustUnmarshalJSON(stakingGenStateBz, &stakingGenState)

	gravityGenState := types.DefaultGenesisState()
	for i, v := range stakingGenState.Validators {
		orchAddr := sdk.AccAddress(v.GetOperator()).String()
		ethAddr := gethcommon.BytesToAddress(bytes.Repeat([]byte{byte(i)}, 20)).String()

		gravityGenState.StaticValCosmosAddrs = append(gravityGenState.StaticValCosmosAddrs, orchAddr)

		gravityGenState.DelegateKeys = append(gravityGenState.DelegateKeys, &types.MsgSetOrchestratorAddress{
			Validator:    v.OperatorAddress,
			Orchestrator: orchAddr,
			EthAddress:   ethAddr,
		})
	}

	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(gravityGenState)
}
