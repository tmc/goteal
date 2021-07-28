// Package types defines the set of relevant types for Algorand Smart Contracts.
package types

import (
	algorand_types "github.com/algorand/go-algorand-sdk/types"
)

// Transaction represents the current transaction that a contract is executing against.
type Transaction algorand_types.Transaction
