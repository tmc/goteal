package helloworld

import (
	"fmt"

	"github.com/tmc/goteal/types"
)

var Foobar = 42
var Foobaz = 43

// Contract defines a trivial contract that checks that the Transaction is a single payment.
func Contract(globals types.Globals, txn types.Transaction, gtxn types.TxGroup) (int, error) {
	x := Foobar == 42 && Foobaz == 43
	if x {
		return 2, nil
	}
	return 0, fmt.Errorf("failed condition")
}
