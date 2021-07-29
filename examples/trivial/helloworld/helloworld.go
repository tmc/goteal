package helloworld

import (
	"fmt"

	"github.com/tmc/goteal/types"
)

// Contract defines a trivial contract that checks that the Transaction is a single payment.
func Contract(globals types.Globals, txn types.Transaction, gtxn types.TxGroup) (int, error) {
	isSingleTx := globals.GroupSize == 3
	// isPayment := txn.TypeEnum == types.PaymentTx

	// if isPayment && isSingleTx {
	if isSingleTx {
		return 2, nil
	}
	return 0, fmt.Errorf("failed condition")
}
