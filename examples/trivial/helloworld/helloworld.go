package helloworld

import (
	"github.com/tmc/goteal/types"
)

// Contract defines a trivial contract that checks that the Transaction is a single payment.
func Contract(globals types.Globals, txn types.Transaction, gtxn types.TxGroup) (int, error) {
	var x int
	if true {
		x = 42
	}
	if true {
		x = 43
	}
	return x, nil
}
