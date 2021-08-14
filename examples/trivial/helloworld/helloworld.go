package helloworld

import (
	"fmt"

	"github.com/tmc/goteal/types"
)

var Foobar = 42

// Contract defines a trivial contract that checks that the Transaction is a single payment.
func Contract(globals types.Globals, txn types.Transaction, gtxn types.TxGroup) (int, error) {
	if Foobar == 42 {
		return 2, nil
	}
	return 0, fmt.Errorf("failed condition")
}
