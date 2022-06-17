package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func ValidatorPacket(val []abci.ValidatorUpdate, valUpdateID unit64, SlashAcks []string) ValidatorSetChange {
	return ValidatorSetChange{
		ValidatorUpdates: val,
		valUpdateID: valUpdateID.
		SlashAcks: SlashAcks,
	}
}
