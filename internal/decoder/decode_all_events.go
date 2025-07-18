package decoder

import (
	"fmt"
	"slices"

	"github.com/ethereum/go-ethereum/core/types"
)

type DecodedEventsCollection struct {
	// mu                              sync.Mutex
	NumberEventsNotDetected int
	ActiveUsers             []string

	Mint                            []MintEvent
	ReserveUsedAsCollateralEnabled  []ReserveUsedAsCollateralEnabledEvent
	UserEModeSet                    []UserEModeSetEvent
	DeficitCovered                  []DeficitCoveredEvent
	BalanceTransfer                 []BalanceTransferEvent
	MintUnbacked                    []MintUnbackedEvent
	Repay                           []RepayEvent
	IsolationModeTotalDebtUpdated   []IsolationModeTotalDebtUpdatedEvent
	Withdraw                        []WithdrawEvent
	Transfer                        []TransferEvent
	ReserveUsedAsCollateralDisabled []ReserveUsedAsCollateralDisabledEvent
	DeficitCreated                  []DeficitCreatedEvent
	FlashLoan                       []FlashLoanEvent
	Approval                        []ApprovalEvent
	LiquidationCall                 []LiquidationCallEvent
	ReserveDataUpdated              []ReserveDataUpdatedEvent
	Burn                            []BurnEvent
	Supply                          []SupplyEvent
	Borrow                          []BorrowEvent
	BorrowAllowanceDelegated        []BorrowAllowanceDelegatedEvent
	BackUnbacked                    []BackUnbackedEvent
	OwnershipTransferred            []OwnershipTransferredEvent
	MintedToTreasury                []MintedToTreasuryEvent
}

func (c *DecodedEventsCollection) AddUser(user string) {
	if !slices.Contains(c.ActiveUsers, user) && user != "0x0000000000000000000000000000000000000000" {
		c.ActiveUsers = append(c.ActiveUsers, user)
	}
}

func DecodeRawEvents(allRawEvents []types.Log, allDecodedEvents *DecodedEventsCollection, eventsCodes map[string]string) {

	for idx, rawEvent := range allRawEvents {
		signature := rawEvent.Topics[0].String()
		switch signature {
		case eventsCodes["LiquidationCall(address,address,address,uint256,uint256,address,bool)"]:
			decodedEvent := DecodeLiquidationCall(rawEvent)
			allDecodedEvents.LiquidationCall = append(allDecodedEvents.LiquidationCall, decodedEvent)
			allDecodedEvents.AddUser(decodedEvent.User)
			allDecodedEvents.AddUser(decodedEvent.Liquidator)

		case eventsCodes["Transfer(address,address,uint256)"]:
			decodedEvent := DecodeTransfer(rawEvent)
			allDecodedEvents.Transfer = append(allDecodedEvents.Transfer, decodedEvent)
			allDecodedEvents.AddUser(decodedEvent.From)
			allDecodedEvents.AddUser(decodedEvent.To)

		case eventsCodes["Initialized(address,address,address,uint8,string,string,bytes)"]:
			fmt.Printf("RawEvent %v is Initialized, ignoring this event", idx)
			allDecodedEvents.NumberEventsNotDetected -= 1

		case eventsCodes["Withdraw(address,address,address,uint256)"]:
			decodedEvent := DecodeWithdraw(rawEvent)
			allDecodedEvents.Withdraw = append(allDecodedEvents.Withdraw, decodedEvent)
			allDecodedEvents.AddUser(decodedEvent.User)
			allDecodedEvents.AddUser(decodedEvent.To)

		case eventsCodes["BackUnbacked(address,address,uint256,uint256)"]:
			decodedEvent := DecodeBackUnbacked(rawEvent)
			allDecodedEvents.BackUnbacked = append(allDecodedEvents.BackUnbacked, decodedEvent)
			allDecodedEvents.AddUser(decodedEvent.Backer)

		case eventsCodes["ReserveUsedAsCollateralDisabled(address,address)"]:
			decodedEvent := DecodeReserveUsedAsCollateralDisabled(rawEvent)
			allDecodedEvents.ReserveUsedAsCollateralDisabled = append(allDecodedEvents.ReserveUsedAsCollateralDisabled, decodedEvent)
			allDecodedEvents.AddUser(decodedEvent.User)

		case eventsCodes["MintUnbacked(address,address,address,uint256,uint16)"]:
			decodedEvent := DecodeMintUnbacked(rawEvent)
			allDecodedEvents.MintUnbacked = append(allDecodedEvents.MintUnbacked, decodedEvent)
			allDecodedEvents.AddUser(decodedEvent.User)
			allDecodedEvents.AddUser(decodedEvent.OnBehalfOf)

		case eventsCodes["Borrow(address,address,address,uint256,uint8,uint256,uint16)"]:
			decodedEvent := DecodeBorrow(rawEvent)
			allDecodedEvents.Borrow = append(allDecodedEvents.Borrow, decodedEvent)
			allDecodedEvents.AddUser(decodedEvent.User)
			allDecodedEvents.AddUser(decodedEvent.OnBehalfOf)

		case eventsCodes["BorrowAllowanceDelegated(address,address,address,uint256)"]:
			decodedEvent := DecodeBorrowAllowanceDelegated(rawEvent)
			allDecodedEvents.BorrowAllowanceDelegated = append(allDecodedEvents.BorrowAllowanceDelegated, decodedEvent)
			allDecodedEvents.AddUser(decodedEvent.FromUser)
			allDecodedEvents.AddUser(decodedEvent.ToUser)

		case eventsCodes["BalanceTransfer(address,address,uint256,uint256)"]:
			decodedEvent := DecodeBalanceTransfer(rawEvent)
			allDecodedEvents.BalanceTransfer = append(allDecodedEvents.BalanceTransfer, decodedEvent)
			allDecodedEvents.AddUser(decodedEvent.From)
			allDecodedEvents.AddUser(decodedEvent.To)

		case eventsCodes["MintedToTreasury(address,uint256)"]:
			decodedEvent := DecodeMintedToTreasury(rawEvent)
			allDecodedEvents.MintedToTreasury = append(allDecodedEvents.MintedToTreasury, decodedEvent)
			// No user implied

		case eventsCodes["DeficitCovered(address,address,uint256)"]:
			decodedEvent := DecodeDeficitCovered(rawEvent)
			allDecodedEvents.DeficitCovered = append(allDecodedEvents.DeficitCovered, decodedEvent)
			allDecodedEvents.AddUser(decodedEvent.Caller)

		case eventsCodes["Burn(address,address,uint256,uint256,uint256)"]:
			decodedEvent := DecodeBurn(rawEvent)
			allDecodedEvents.Burn = append(allDecodedEvents.Burn, decodedEvent)
			allDecodedEvents.AddUser(decodedEvent.From)
			allDecodedEvents.AddUser(decodedEvent.Target)

		case eventsCodes["OwnershipTransferred(address,address)"]:
			decodedEvent := DecodeOwnershipTransferred(rawEvent)
			allDecodedEvents.OwnershipTransferred = append(allDecodedEvents.OwnershipTransferred, decodedEvent)
			allDecodedEvents.AddUser(decodedEvent.PreviousOwner)
			allDecodedEvents.AddUser(decodedEvent.NewOwner)

		case eventsCodes["Approval(address,address,uint256)"]:
			decodedEvent := DecodeApproval(rawEvent)
			allDecodedEvents.Approval = append(allDecodedEvents.Approval, decodedEvent)
			allDecodedEvents.AddUser(decodedEvent.Owner)
			allDecodedEvents.AddUser(decodedEvent.Spender)

		case eventsCodes["Initialized(address,address,address,address,uint8,string,string,bytes)"]:
			fmt.Printf("RawEvent %v is Initialized, ignoring this event", idx)
			allDecodedEvents.NumberEventsNotDetected -= 1

		case eventsCodes["UserEModeSet(address,uint8)"]:
			decodedEvent := DecodeUserEModeSet(rawEvent)
			allDecodedEvents.UserEModeSet = append(allDecodedEvents.UserEModeSet, decodedEvent)
			allDecodedEvents.AddUser(decodedEvent.User)

		case eventsCodes["DeficitCreated(address,address,uint256)"]:
			decodedEvent := DecodeDeficitCreated(rawEvent)
			allDecodedEvents.DeficitCreated = append(allDecodedEvents.DeficitCreated, decodedEvent)
			allDecodedEvents.AddUser(decodedEvent.User)

		case eventsCodes["Mint(address,address,uint256,uint256,uint256)"]:
			decodedEvent := DecodeMint(rawEvent)
			allDecodedEvents.Mint = append(allDecodedEvents.Mint, decodedEvent)
			allDecodedEvents.AddUser(decodedEvent.Caller)
			allDecodedEvents.AddUser(decodedEvent.OnBehalfOf)

		case eventsCodes["IsolationModeTotalDebtUpdated(address,uint256)"]:
			decodedEvent := DecodeIsolationModeTotalDebtUpdated(rawEvent)
			allDecodedEvents.IsolationModeTotalDebtUpdated = append(allDecodedEvents.IsolationModeTotalDebtUpdated, decodedEvent)
			// No user implied

		case eventsCodes["FlashLoan(address,address,address,uint256,uint8,uint256,uint16)"]:
			decodedEvent := DecodeFlashLoan(rawEvent)
			allDecodedEvents.FlashLoan = append(allDecodedEvents.FlashLoan, decodedEvent)
			allDecodedEvents.AddUser(decodedEvent.Target)
			allDecodedEvents.AddUser(decodedEvent.Initiator)

		case eventsCodes["ReserveDataUpdated(address,uint256,uint256,uint256,uint256,uint256)"]:
			decodedEvent := DecodeReserveDataUpdated(rawEvent)
			allDecodedEvents.ReserveDataUpdated = append(allDecodedEvents.ReserveDataUpdated, decodedEvent)
			// No user implied

		case eventsCodes["Supply(address,address,address,uint256,uint16)"]:
			decodedEvent := DecodeSupply(rawEvent)
			allDecodedEvents.Supply = append(allDecodedEvents.Supply, decodedEvent)
			allDecodedEvents.AddUser(decodedEvent.User)
			allDecodedEvents.AddUser(decodedEvent.OnBehalfOf)

		case eventsCodes["ReserveUsedAsCollateralEnabled(address,address)"]:
			decodedEvent := DecodeReserveUsedAsCollateralEnabled(rawEvent)
			allDecodedEvents.ReserveUsedAsCollateralEnabled = append(allDecodedEvents.ReserveUsedAsCollateralEnabled, decodedEvent)
			allDecodedEvents.AddUser(decodedEvent.User)

		case eventsCodes["Repay(address,address,address,uint256,bool)"]:
			decodedEvent := DecodeRepay(rawEvent)
			allDecodedEvents.Repay = append(allDecodedEvents.Repay, decodedEvent)
			allDecodedEvents.AddUser(decodedEvent.User)
			allDecodedEvents.AddUser(decodedEvent.Repayer)
		}
	}
}
