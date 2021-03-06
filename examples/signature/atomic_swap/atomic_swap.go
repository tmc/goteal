package atomic_swap

import (
	"bytes"
	"fmt"

	"github.com/tmc/goteal/avm"
	"github.com/tmc/goteal/types"
)

const (
	Alice      = "6ZHGHH5Z5CTPCF5WCESXMGRSVK7QJETR63M3NY5FJCUYDHO57VTCMJOBGY"
	Bob        = "7Z5PWO2C6LFNQFGHWKSK5H47IQP5OJW2M3HA2QPXTY3WTNP5NU2MHBW27M"
	TmpFeeCond = 1000
	Timeout    = 3000
)

var Secret = []byte("232323232323232323")

func Contract(globals types.Globals, txn types.Transaction, gtxn types.TxGroup) (int, error) {
	// set fee condition
	feeCond := txn.Fee < TmpFeeCond

	// set safety conditions
	isPayment := txn.TypeEnum == types.PaymentTx
	isCloseSet := txn.CloseRemainderTo == globals.ZeroAddress
	isRekeySet := txn.RekeyTo == globals.ZeroAddress

	// maybe
	// isSafetyCond := (txn.TypeEnum == types.PaymentTx) && (txn.CloseRemainderTo == globals.ZeroAddress) && (txn.RekeyTo == globals.ZeroAddress)
	safetyCond := isPayment && isCloseSet && isRekeySet

	// set receive conditions
	isReceiverSeller := txn.Receiver == Alice
	//needs help
	isSecretMatch := bytes.Compare(avm.Sha256([]byte(Secret)), Secret) == 0
	receiveCond := isReceiverSeller && isSecretMatch

	//set escape conditions
	isReceiverBuyer := txn.Receiver == Bob
	isFirstValid := txn.FirstValid > Timeout

	escCond := isReceiverBuyer && isFirstValid

	if feeCond && safetyCond && (receiveCond || escCond) {
		return 1, nil
	}
	return 0, fmt.Errorf("failed condition")

}
