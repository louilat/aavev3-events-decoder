package decoder

import "math/big"

type DeficitCreatedEvent struct {
	Address       string
	BlockNumber   *big.Int
	TxHash        string
	User          string
	DebtAsset     string
	AmountCreated *big.Int
}

type RepayEvent struct {
	Address     string
	BlockNumber *big.Int
	TxHash      string
	Reserve     string
	User        string
	Repayer     string
	Amount      *big.Int
	UseATokens  bool
}

type ApprovalEvent struct {
	Address     string
	BlockNumber *big.Int
	TxHash      string
	Owner       string
	Spender     string
	Value       *big.Int
}

type BackUnbackedEvent struct {
	Address     string
	BlockNumber *big.Int
	TxHash      string
	Reserve     string
	Backer      string
	Amount      *big.Int
	Fee         *big.Int
}

type BorrowEvent struct {
	Address          string
	BlockNumber      *big.Int
	TxHash           string
	Reserve          string
	User             string
	OnBehalfOf       string
	Amount           *big.Int
	InterestRateMode *big.Int
	BorrowRate       *big.Int
	ReferralCode     *big.Int
}

type MintedToTreasuryEvent struct {
	Address      string
	BlockNumber  *big.Int
	TxHash       string
	Reserve      string
	AmountMinted *big.Int
}

type FlashLoanEvent struct {
	Address          string
	BlockNumber      *big.Int
	TxHash           string
	Target           string
	Initiator        string
	Asset            string
	Amount           *big.Int
	InterestRateMode *big.Int
	Premium          *big.Int
	ReferralCode     *big.Int
}

type BalanceTransferEvent struct {
	Address     string
	BlockNumber *big.Int
	TxHash      string
	From        string
	To          string
	Value       *big.Int
	Index       *big.Int
}

type BurnEvent struct {
	Address         string
	BlockNumber     *big.Int
	TxHash          string
	From            string
	Target          string
	Value           *big.Int
	BalanceIncrease *big.Int
	Index           *big.Int
}

type MintUnbackedEvent struct {
	Address      string
	BlockNumber  *big.Int
	TxHash       string
	Reserve      string
	User         string
	OnBehalfOf   string
	Amount       *big.Int
	ReferralCode *big.Int
}

type UserEModeSetEvent struct {
	Address     string
	BlockNumber *big.Int
	TxHash      string
	User        string
	CategoryId  *big.Int
}

type ReserveUsedAsCollateralDisabledEvent struct {
	Address     string
	BlockNumber *big.Int
	TxHash      string
	Reserve     string
	User        string
}

type ReserveDataUpdatedEvent struct {
	Address             string
	BlockNumber         *big.Int
	TxHash              string
	Reserve             string
	LiquidityRate       *big.Int
	StableBorrowRate    *big.Int
	VariableBorrowRate  *big.Int
	LiquidityIndex      *big.Int
	VariableBorrowIndex *big.Int
}

type BorrowAllowanceDelegatedEvent struct {
	Address     string
	BlockNumber *big.Int
	TxHash      string
	FromUser    string
	ToUser      string
	Asset       string
	Amount      *big.Int
}

type LiquidationCallEvent struct {
	Address                    string
	BlockNumber                *big.Int
	TxHash                     string
	CollateralAsset            string
	DebtAsset                  string
	User                       string
	DebtToCover                *big.Int
	LiquidatedCollateralAmount *big.Int
	Liquidator                 string
	ReceiveAToken              bool
}

type IsolationModeTotalDebtUpdatedEvent struct {
	Address     string
	BlockNumber *big.Int
	TxHash      string
	Asset       string
	TotalDebt   *big.Int
}

type WithdrawEvent struct {
	Address     string
	BlockNumber *big.Int
	TxHash      string
	Reserve     string
	User        string
	To          string
	Amount      *big.Int
}

type ReserveUsedAsCollateralEnabledEvent struct {
	Address     string
	BlockNumber *big.Int
	TxHash      string
	Reserve     string
	User        string
}

type TransferEvent struct {
	Address     string
	BlockNumber *big.Int
	TxHash      string
	From        string
	To          string
	Value       *big.Int
}

type SupplyEvent struct {
	Address      string
	BlockNumber  *big.Int
	TxHash       string
	Reserve      string
	User         string
	OnBehalfOf   string
	Amount       *big.Int
	ReferralCode *big.Int
}

type MintEvent struct {
	Address         string
	BlockNumber     *big.Int
	TxHash          string
	Caller          string
	OnBehalfOf      string
	Value           *big.Int
	BalanceIncrease *big.Int
	Index           *big.Int
}

type OwnershipTransferredEvent struct {
	Address       string
	BlockNumber   *big.Int
	TxHash        string
	PreviousOwner string
	NewOwner      string
}

type DeficitCoveredEvent struct {
	Address       string
	BlockNumber   *big.Int
	TxHash        string
	Reserve       string
	Caller        string
	AmountCovered *big.Int
}
