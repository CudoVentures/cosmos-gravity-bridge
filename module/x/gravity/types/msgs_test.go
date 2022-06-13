package types

import (
	"bytes"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
)

func TestValidateMsgSetOrchestratorAddress(t *testing.T) {
	var (
		ethAddress                   = "0xb462864E395d88d6bc7C5dd5F3F5eb4cc2599255"
		cosmosAddress sdk.AccAddress = bytes.Repeat([]byte{0x1}, sdk.AddrLen)
		valAddress    sdk.ValAddress = bytes.Repeat([]byte{0x1}, sdk.AddrLen)
	)
	specs := map[string]struct {
		srcCosmosAddr sdk.AccAddress
		srcValAddr    sdk.ValAddress
		srcETHAddr    string
		expErr        bool
	}{
		"all good": {
			srcCosmosAddr: cosmosAddress,
			srcValAddr:    valAddress,
			srcETHAddr:    ethAddress,
		},
		"empty validator address": {
			srcETHAddr:    ethAddress,
			srcCosmosAddr: cosmosAddress,
			expErr:        true,
		},
		"invalid validator address": {
			srcValAddr:    []byte{0x1},
			srcCosmosAddr: cosmosAddress,
			srcETHAddr:    ethAddress,
			expErr:        true,
		},
		"empty cosmos address": {
			srcValAddr: valAddress,
			srcETHAddr: ethAddress,
			expErr:     true,
		},
		"invalid cosmos address": {
			srcCosmosAddr: []byte{0x1},
			srcValAddr:    valAddress,
			srcETHAddr:    ethAddress,
			expErr:        true,
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			ethAddr, err := NewEthAddress(spec.srcETHAddr)
			msg := NewMsgSetOrchestratorAddress(spec.srcValAddr, spec.srcCosmosAddr, *ethAddr)
			// when
			err = msg.ValidateBasic()
			if spec.expErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}

func TestMsgCancelSendToEth(t *testing.T) {
	var (
		senderAddress sdk.AccAddress = sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address().Bytes())
	)

	specs := map[string]struct {
		sender        sdk.AccAddress
		transactionId uint64
		expErr        bool
	}{
		"all good": {
			sender:        senderAddress,
			transactionId: 1,
			expErr:        false,
		},
		"invalid address": {
			sender:        []byte{0x1},
			transactionId: 1,
			expErr:        true,
		},
	}

	for msg, spec := range specs {
		fmt.Println(msg)
		t.Run(msg, func(t *testing.T) {
			msg := NewMsgCancelSendToEth(spec.sender, spec.transactionId)
			// when
			err := msg.ValidateBasic()
			if spec.expErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
		})
	}
}
