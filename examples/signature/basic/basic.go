package basic

import (
	"fmt"

	"github.com/tmc/goteal/types"
)

// ExpectedReceiver is the expected recipient address.
const ExpectedReceiver = "ZZAF5ARA4MEC5PVDOP64JM5O5MQST63Q2KOY2FLYFLXXD3PFSNJJBYAFZM"

func Contract(globals types.Globals, txn types.Transaction, gtxn types.TxGroup) (int, error) {
	isPayment := txn.TypeEnum == types.PaymentTx
	isSingleTx := globals.GroupSize == 1
	isCorrectReceiver := txn.Receiver == ExpectedReceiver
	noCloseOutAddr := txn.CloseRemainderTo == globals.ZeroAddress
	noRekeyAddr := txn.RekeyTo == globals.ZeroAddress
	acceptableFee := txn.Fee <= 1000

	if isPayment && isSingleTx && isCorrectReceiver && noCloseOutAddr && noRekeyAddr && acceptableFee {
		return 1, nil
	}
	return 0, fmt.Errorf("failed condition")
}
