package decoder

import (
	"encoding/hex"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func DecodeMint(rawEvent types.Log) MintEvent {
	data := hex.EncodeToString(rawEvent.Data)
	value, _ := new(big.Int).SetString(data[0:64], 16)
	balanceIncrease, _ := new(big.Int).SetString(data[64:128], 16)
	index, _ := new(big.Int).SetString(data[128:192], 16)

	return MintEvent{
		Address:         rawEvent.Address.Hex(),
		BlockNumber:     big.NewInt(int64(rawEvent.BlockNumber)),
		TxHash:          rawEvent.TxHash.Hex(),
		Caller:          common.HexToAddress(rawEvent.Topics[1].String()).String(),
		OnBehalfOf:      common.HexToAddress(rawEvent.Topics[2].String()).String(),
		Value:           value,
		BalanceIncrease: balanceIncrease,
		Index:           index,
	}
}

func DecodeDeficitCovered(rawEvent types.Log) DeficitCoveredEvent {
	data := hex.EncodeToString(rawEvent.Data)
	caller := common.HexToAddress(data[0:64]).String()
	amountCovered, _ := new(big.Int).SetString(data[64:128], 16)
	return DeficitCoveredEvent{
		Address:       rawEvent.Address.Hex(),
		BlockNumber:   big.NewInt(int64(rawEvent.BlockNumber)),
		TxHash:        rawEvent.TxHash.Hex(),
		Reserve:       common.HexToAddress(rawEvent.Topics[1].String()).String(),
		Caller:        caller,
		AmountCovered: amountCovered,
	}
}

func DecodeBorrowAllowanceDelegated(rawEvent types.Log) BorrowAllowanceDelegatedEvent {
	data := hex.EncodeToString(rawEvent.Data)
	amount, _ := new(big.Int).SetString(data[0:64], 16)
	return BorrowAllowanceDelegatedEvent{
		Address:     rawEvent.Address.Hex(),
		BlockNumber: big.NewInt(int64(rawEvent.BlockNumber)),
		TxHash:      rawEvent.TxHash.Hex(),
		FromUser:    common.HexToAddress(rawEvent.Topics[1].String()).String(),
		ToUser:      common.HexToAddress(rawEvent.Topics[2].String()).String(),
		Asset:       common.HexToAddress(rawEvent.Topics[3].String()).String(),
		Amount:      amount,
	}
}

func DecodeRepay(rawEvent types.Log) RepayEvent {
	data := hex.EncodeToString(rawEvent.Data)
	amount, _ := new(big.Int).SetString(data[0:64], 16)
	useATokens, _ := new(big.Int).SetString(data[64:128], 16)
	return RepayEvent{
		Address:     rawEvent.Address.Hex(),
		BlockNumber: big.NewInt(int64(rawEvent.BlockNumber)),
		TxHash:      rawEvent.TxHash.Hex(),
		Reserve:     common.HexToAddress(rawEvent.Topics[1].String()).String(),
		User:        common.HexToAddress(rawEvent.Topics[2].String()).String(),
		Repayer:     common.HexToAddress(rawEvent.Topics[3].String()).String(),
		Amount:      amount,
		UseATokens:  useATokens.Cmp(big.NewInt(1)) == 0,
	}
}

func DecodeIsolationModeTotalDebtUpdated(rawEvent types.Log) IsolationModeTotalDebtUpdatedEvent {
	data := hex.EncodeToString(rawEvent.Data)
	totalDebt, _ := new(big.Int).SetString(data[0:64], 16)
	return IsolationModeTotalDebtUpdatedEvent{
		Address:     rawEvent.Address.Hex(),
		BlockNumber: big.NewInt(int64(rawEvent.BlockNumber)),
		TxHash:      rawEvent.TxHash.Hex(),
		Asset:       common.HexToAddress(rawEvent.Topics[1].String()).String(),
		TotalDebt:   totalDebt,
	}
}

func DecodeDeficitCreated(rawEvent types.Log) DeficitCreatedEvent {
	data := hex.EncodeToString(rawEvent.Data)
	amountCreated, _ := new(big.Int).SetString(data[0:64], 16)
	return DeficitCreatedEvent{
		Address:       rawEvent.Address.Hex(),
		BlockNumber:   big.NewInt(int64(rawEvent.BlockNumber)),
		TxHash:        rawEvent.TxHash.Hex(),
		User:          common.HexToAddress(rawEvent.Topics[1].String()).String(),
		DebtAsset:     common.HexToAddress(rawEvent.Topics[2].String()).String(),
		AmountCreated: amountCreated,
	}
}

func DecodeBackUnbacked(rawEvent types.Log) BackUnbackedEvent {
	data := hex.EncodeToString(rawEvent.Data)
	amount, _ := new(big.Int).SetString(data[0:64], 16)
	fee, _ := new(big.Int).SetString(data[64:128], 16)
	return BackUnbackedEvent{
		Address:     rawEvent.Address.Hex(),
		BlockNumber: big.NewInt(int64(rawEvent.BlockNumber)),
		TxHash:      rawEvent.TxHash.Hex(),
		Reserve:     common.HexToAddress(rawEvent.Topics[1].String()).String(),
		Backer:      common.HexToAddress(rawEvent.Topics[2].String()).String(),
		Amount:      amount,
		Fee:         fee,
	}
}

func DecodeBurn(rawEvent types.Log) BurnEvent {
	data := hex.EncodeToString(rawEvent.Data)
	value, _ := new(big.Int).SetString(data[0:64], 16)
	balanceIncrease, _ := new(big.Int).SetString(data[64:128], 16)
	index, _ := new(big.Int).SetString(data[128:192], 16)
	return BurnEvent{
		Address:         rawEvent.Address.Hex(),
		BlockNumber:     big.NewInt(int64(rawEvent.BlockNumber)),
		TxHash:          rawEvent.TxHash.Hex(),
		From:            common.HexToAddress(rawEvent.Topics[1].String()).String(),
		Target:          common.HexToAddress(rawEvent.Topics[2].String()).String(),
		Value:           value,
		BalanceIncrease: balanceIncrease,
		Index:           index,
	}
}

func DecodeReserveDataUpdated(rawEvent types.Log) ReserveDataUpdatedEvent {
	data := hex.EncodeToString(rawEvent.Data)
	liqRate, _ := new(big.Int).SetString(data[0:64], 16)
	stbRate, _ := new(big.Int).SetString(data[64:128], 16)
	vabRate, _ := new(big.Int).SetString(data[128:192], 16)
	liqIdx, _ := new(big.Int).SetString(data[192:256], 16)
	vabIdx, _ := new(big.Int).SetString(data[256:320], 16)

	return ReserveDataUpdatedEvent{
		Address:             rawEvent.Address.Hex(),
		BlockNumber:         big.NewInt(int64(rawEvent.BlockNumber)),
		TxHash:              rawEvent.TxHash.Hex(),
		Reserve:             common.HexToAddress(rawEvent.Topics[1].String()).String(),
		LiquidityRate:       liqRate,
		StableBorrowRate:    stbRate,
		VariableBorrowRate:  vabRate,
		LiquidityIndex:      liqIdx,
		VariableBorrowIndex: vabIdx,
	}
}

func DecodeOwnershipTransferred(rawEvent types.Log) OwnershipTransferredEvent {
	return OwnershipTransferredEvent{
		Address:       rawEvent.Address.Hex(),
		BlockNumber:   big.NewInt(int64(rawEvent.BlockNumber)),
		TxHash:        rawEvent.TxHash.Hex(),
		PreviousOwner: common.HexToAddress(rawEvent.Topics[1].String()).String(),
		NewOwner:      common.HexToAddress(rawEvent.Topics[2].String()).String(),
	}
}

func DecodeLiquidationCall(rawEvent types.Log) LiquidationCallEvent {
	data := hex.EncodeToString(rawEvent.Data)
	debtToCover, _ := new(big.Int).SetString(data[0:64], 16)
	liquidatedCollateralAmount, _ := new(big.Int).SetString(data[64:128], 16)
	liquidator := common.HexToAddress(data[128:192]).String()
	receiveAToken, _ := new(big.Int).SetString(data[192:256], 16)
	return LiquidationCallEvent{
		Address:                    rawEvent.Address.Hex(),
		BlockNumber:                big.NewInt(int64(rawEvent.BlockNumber)),
		TxHash:                     rawEvent.TxHash.Hex(),
		CollateralAsset:            common.HexToAddress(rawEvent.Topics[1].String()).String(),
		DebtAsset:                  common.HexToAddress(rawEvent.Topics[2].String()).String(),
		User:                       common.HexToAddress(rawEvent.Topics[3].String()).String(),
		DebtToCover:                debtToCover,
		LiquidatedCollateralAmount: liquidatedCollateralAmount,
		Liquidator:                 liquidator,
		ReceiveAToken:              receiveAToken.Cmp(big.NewInt(1)) == 0,
	}
}

func DecodeBorrow(rawEvent types.Log) BorrowEvent {
	data := hex.EncodeToString(rawEvent.Data)
	user := common.HexToAddress(data[0:64]).String()
	amount, _ := new(big.Int).SetString(data[64:128], 16)
	interestRateMode, _ := new(big.Int).SetString(data[128:192], 16)
	borrowRate, _ := new(big.Int).SetString(data[192:256], 16)
	referralCode, _ := new(big.Int).SetString(rawEvent.Topics[3].String()[2:], 16)
	return BorrowEvent{
		Address:          rawEvent.Address.Hex(),
		BlockNumber:      big.NewInt(int64(rawEvent.BlockNumber)),
		TxHash:           rawEvent.TxHash.Hex(),
		Reserve:          common.HexToAddress(rawEvent.Topics[1].String()).String(),
		User:             user,
		OnBehalfOf:       common.HexToAddress(rawEvent.Topics[2].String()).String(),
		Amount:           amount,
		InterestRateMode: interestRateMode,
		BorrowRate:       borrowRate,
		ReferralCode:     referralCode,
	}
}

func DecodeUserEModeSet(rawEvent types.Log) UserEModeSetEvent {
	data := hex.EncodeToString(rawEvent.Data)
	categoryId, _ := new(big.Int).SetString(data[0:64], 16)
	return UserEModeSetEvent{
		Address:     rawEvent.Address.Hex(),
		BlockNumber: big.NewInt(int64(rawEvent.BlockNumber)),
		TxHash:      rawEvent.TxHash.Hex(),
		User:        common.HexToAddress(rawEvent.Topics[1].String()).String(),
		CategoryId:  categoryId,
	}
}

func DecodeBalanceTransfer(rawEvent types.Log) BalanceTransferEvent {
	data := hex.EncodeToString(rawEvent.Data)
	value, _ := new(big.Int).SetString(data[0:64], 16)
	index, _ := new(big.Int).SetString(data[64:128], 16)
	return BalanceTransferEvent{
		Address:     rawEvent.Address.Hex(),
		BlockNumber: big.NewInt(int64(rawEvent.BlockNumber)),
		TxHash:      rawEvent.TxHash.Hex(),
		From:        common.HexToAddress(rawEvent.Topics[1].String()).String(),
		To:          common.HexToAddress(rawEvent.Topics[2].String()).String(),
		Value:       value,
		Index:       index,
	}
}

func DecodeSupply(rawEvent types.Log) SupplyEvent {
	data := hex.EncodeToString(rawEvent.Data)
	user := common.HexToAddress(data[0:64]).String()
	amount, _ := new(big.Int).SetString(data[64:128], 16)
	referralCode, _ := new(big.Int).SetString(rawEvent.Topics[3].String()[2:], 16)
	return SupplyEvent{
		Address:      rawEvent.Address.Hex(),
		BlockNumber:  big.NewInt(int64(rawEvent.BlockNumber)),
		TxHash:       rawEvent.TxHash.Hex(),
		Reserve:      common.HexToAddress(rawEvent.Topics[1].String()).String(),
		User:         user,
		OnBehalfOf:   common.HexToAddress(rawEvent.Topics[2].String()).String(),
		Amount:       amount,
		ReferralCode: referralCode,
	}
}

func DecodeTransfer(rawEvent types.Log) TransferEvent {
	data := hex.EncodeToString(rawEvent.Data)
	value, _ := new(big.Int).SetString(data[0:64], 16)
	return TransferEvent{
		Address:     rawEvent.Address.Hex(),
		BlockNumber: big.NewInt(int64(rawEvent.BlockNumber)),
		TxHash:      rawEvent.TxHash.Hex(),
		From:        common.HexToAddress(rawEvent.Topics[1].String()).String(),
		To:          common.HexToAddress(rawEvent.Topics[2].String()).String(),
		Value:       value,
	}
}

func DecodeMintedToTreasury(rawEvent types.Log) MintedToTreasuryEvent {
	data := hex.EncodeToString(rawEvent.Data)
	amountMinted, _ := new(big.Int).SetString(data[0:64], 16)
	return MintedToTreasuryEvent{
		Address:      rawEvent.Address.Hex(),
		BlockNumber:  big.NewInt(int64(rawEvent.BlockNumber)),
		TxHash:       rawEvent.TxHash.Hex(),
		Reserve:      common.HexToAddress(rawEvent.Topics[1].String()).String(),
		AmountMinted: amountMinted,
	}
}

func DecodeReserveUsedAsCollateralDisabled(rawEvent types.Log) ReserveUsedAsCollateralDisabledEvent {
	return ReserveUsedAsCollateralDisabledEvent{
		Address:     rawEvent.Address.Hex(),
		BlockNumber: big.NewInt(int64(rawEvent.BlockNumber)),
		TxHash:      rawEvent.TxHash.Hex(),
		Reserve:     common.HexToAddress(rawEvent.Topics[1].String()).String(),
		User:        common.HexToAddress(rawEvent.Topics[2].String()).String(),
	}
}

func DecodeReserveUsedAsCollateralEnabled(rawEvent types.Log) ReserveUsedAsCollateralEnabledEvent {
	return ReserveUsedAsCollateralEnabledEvent{
		Address:     rawEvent.Address.Hex(),
		BlockNumber: big.NewInt(int64(rawEvent.BlockNumber)),
		TxHash:      rawEvent.TxHash.Hex(),
		Reserve:     common.HexToAddress(rawEvent.Topics[1].String()).String(),
		User:        common.HexToAddress(rawEvent.Topics[2].String()).String(),
	}
}

func DecodeApproval(rawEvent types.Log) ApprovalEvent {
	data := hex.EncodeToString(rawEvent.Data)
	value, _ := new(big.Int).SetString(data[0:64], 16)
	return ApprovalEvent{
		Address:     rawEvent.Address.Hex(),
		BlockNumber: big.NewInt(int64(rawEvent.BlockNumber)),
		TxHash:      rawEvent.TxHash.Hex(),
		Owner:       common.HexToAddress(rawEvent.Topics[1].String()).String(),
		Spender:     common.HexToAddress(rawEvent.Topics[2].String()).String(),
		Value:       value,
	}
}

func DecodeWithdraw(rawEvent types.Log) WithdrawEvent {
	data := hex.EncodeToString(rawEvent.Data)
	amount, _ := new(big.Int).SetString(data[0:64], 16)
	return WithdrawEvent{
		Address:     rawEvent.Address.Hex(),
		BlockNumber: big.NewInt(int64(rawEvent.BlockNumber)),
		TxHash:      rawEvent.TxHash.Hex(),
		Reserve:     common.HexToAddress(rawEvent.Topics[1].String()).String(),
		User:        common.HexToAddress(rawEvent.Topics[2].String()).String(),
		To:          common.HexToAddress(rawEvent.Topics[3].String()).String(),
		Amount:      amount,
	}
}

func DecodeMintUnbacked(rawEvent types.Log) MintUnbackedEvent {
	data := hex.EncodeToString(rawEvent.Data)
	user := common.HexToAddress(data[0:64]).String()
	amount, _ := new(big.Int).SetString(data[64:128], 16)
	referralCode, _ := new(big.Int).SetString(rawEvent.Topics[3].String()[2:], 16)
	return MintUnbackedEvent{
		Address:      rawEvent.Address.Hex(),
		BlockNumber:  big.NewInt(int64(rawEvent.BlockNumber)),
		TxHash:       rawEvent.TxHash.Hex(),
		Reserve:      common.HexToAddress(rawEvent.Topics[1].String()).String(),
		User:         user,
		OnBehalfOf:   common.HexToAddress(rawEvent.Topics[2].String()).String(),
		Amount:       amount,
		ReferralCode: referralCode,
	}
}

func DecodeFlashLoan(rawEvent types.Log) FlashLoanEvent {
	data := hex.EncodeToString(rawEvent.Data)
	initiator := common.HexToAddress(data[0:64]).String()
	amount, _ := new(big.Int).SetString(data[64:128], 16)
	interestRateMode, _ := new(big.Int).SetString(data[128:192], 16)
	premium, _ := new(big.Int).SetString(data[192:256], 16)
	referralCode, _ := new(big.Int).SetString(rawEvent.Topics[3].String()[2:], 16)
	return FlashLoanEvent{
		Address:          rawEvent.Address.Hex(),
		BlockNumber:      big.NewInt(int64(rawEvent.BlockNumber)),
		TxHash:           rawEvent.TxHash.Hex(),
		Target:           common.HexToAddress(rawEvent.Topics[1].String()).String(),
		Initiator:        initiator,
		Asset:            common.HexToAddress(rawEvent.Topics[2].String()).String(),
		Amount:           amount,
		InterestRateMode: interestRateMode,
		Premium:          premium,
		ReferralCode:     referralCode,
	}
}
