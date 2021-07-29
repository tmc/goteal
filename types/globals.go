package types

// Globals encodes a set of global values relevant for an Algorand Smart Contract.
type Globals struct {
	// micro Algos
	MinTxnFee uint64
	// micro Algos
	MinBalance uint64
	// rounds
	MaxTxnLife uint64
	// 32 byte address of all zero bytes
	ZeroAddress Address
	// Number of transactions in this atomic transaction group. At least 1
	GroupSize uint64
	// Maximum supported TEAL version
	LogicSigVersion uint64
	// Current round number
	Round uint64
	// Last confirmed block UNIX timestamp. Fails if negative
	LatestTimestamp uint64
	// ID of current application executing. Fails if no such application is executing
	CurrentApplicationID uint64
	// Address of the creator of the current application. Fails if no such application is executing
	CreatorAddress []byte
}
