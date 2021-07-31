package helloworld

import (
	"fmt"

	"github.com/tmc/goteal/types"
)

// Contract defines a trivial contract that checks that the Transaction is a single payment.
func Contract(globals types.Globals, txn types.Transaction, gtxn types.TxGroup) (int, error) {
	isSingleTx := globals.GroupSize == 1

	if isSingleTx {
		return 1, nil
	}
	return 0, fmt.Errorf("failed condition")
}
