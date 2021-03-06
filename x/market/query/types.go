package query

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ovrclk/akash/x/market/types"
)

type (
	// Order type
	Order types.Order
	// Orders - Slice of Order Struct
	Orders []Order

	// Bid type
	Bid types.Bid
	// Bids - Slice of Bid Struct
	Bids []Bid

	// Lease type
	Lease types.Lease
	// Leases - Slice of Lease Struct
	Leases []Lease
)

const (
	todo = "TODO see deployment/query/types.go"
)

// OrderFilters defines flags for order list filter
type OrderFilters struct {
	Owner sdk.AccAddress
	// State flag value given
	StateFlagVal string
	// Actual state value decoded from OrderStateMap
	State types.OrderState
}

// BidFilters defines flags for bid list filter
type BidFilters struct {
	Owner sdk.AccAddress
	// State flag value given
	StateFlagVal string
	// Actual state value decoded from BidStateMap
	State types.BidState
}

// LeaseFilters defines flags for lease list filter
type LeaseFilters struct {
	Owner sdk.AccAddress
	// State flag value given
	StateFlagVal string
	// Actual state value decoded from LeaseStateMap
	State types.LeaseState
}

// Accept returns true if object matches filter requirements
func (f OrderFilters) Accept(obj types.Order, isValidState bool) bool {
	if (f.Owner.Empty() && !isValidState) ||
		(f.Owner.Empty() && (obj.State == f.State)) ||
		(!isValidState && obj.OrderID.Owner.Equals(f.Owner)) ||
		(obj.OrderID.Owner.Equals(f.Owner) && obj.State == f.State) {
		return true
	}

	return false
}

// Accept returns true if object matches filter requirements
func (f BidFilters) Accept(obj types.Bid, isValidState bool) bool {
	if (f.Owner.Empty() && !isValidState) ||
		(f.Owner.Empty() && (obj.State == f.State)) ||
		(!isValidState && obj.BidID.Owner.Equals(f.Owner)) ||
		(obj.BidID.Owner.Equals(f.Owner) && obj.State == f.State) {
		return true
	}

	return false
}

// Accept returns true if object matches filter requirements
func (f LeaseFilters) Accept(obj types.Lease, isValidState bool) bool {
	if (f.Owner.Empty() && !isValidState) ||
		(f.Owner.Empty() && (obj.State == f.State)) ||
		(!isValidState && (obj.LeaseID.Owner.Equals(f.Owner))) ||
		(obj.LeaseID.Owner.Equals(f.Owner) && obj.State == f.State) {
		return true
	}
	return false
}

func (obj Order) String() string {
	return todo
}

func (obj Orders) String() string {
	return todo
}

func (obj Bid) String() string {
	return todo
}

func (obj Bids) String() string {
	return todo
}

func (obj Lease) String() string {
	return todo
}

func (obj Leases) String() string {
	return todo
}
