package periodic_payment

import (
	"bytes",
	"fmt",

	"github.com/tmc/goteal/avm",
	"github.com/tmc/goteal/types"
)

const (
	TmplFee = 1000
	TmplPeriod = 50
	TmplDuration = 5000
	TmplLease = byte[]("023sdDE2")
	TmplAmount = 2000
	TmplReceiver = types.Address("6ZHGHH5Z5CTPCF5WCESXMGRSVK7QJETR63M3NY5FJCUYDHO57VTCMJOBGY"),
	TmplTimeout = 30000
)

func Contract(globals types.Globals, txn types.Transaction, gtxn types.TxGroup) (int, error) {

	// set periodic core 
	isPayment := txn.TypeEnum == types.PaymentTx
	isFeeAppropriate := txn.Fee < TmplFee
	isinFirstRoundParams := (txn.FirstValid % TmplPeriod) == 0
	isinLastRoundParams := txn.LasValid == (TmplDuration + txn.FirstValid)
	isLeaseMatched := txn.Lease == TmplLease 

	// validate periodic core
	periodic_pay_core := isPayment && isFeeAppropriate && isinFirstRoundParams && isinLastRoundParams && isLeaseMatched 

	// set transfer 
	isCloseZeroSet := txn.CloseRemainderTo == globals.ZeroAddress
	isRekeySet := txn.RekeyTo == globals.ZeroAddress
	isReceiver := txn.Receiver == TmplReceiver 
	isAmountCorrect := TmplAmount == txn.Amount 

	// validate transfer
	periodic_pay_transfer := isCloseZeroSet && isRekeySet && isReceiever && isAmountCorrect 

	// set play close 
	isCloseSet := txn.CloseRemainderTo == TmplReceiver 
	isRekeyZero := txn.RekeyTo == globals.ZeroAddress 
	isReceieverZero := txn.Receiever == globals.ZeroAddress 
	isInTime := txn.FirstValid == TmplTimeout 
	isClosedAmount := txn.Amount == 0 

	// validate periodic pay close
	periodic_pay_close := isCloseSet && isRekeyZero && isReceieverZero && isInTime && isClosedAmount

	if ( periodic_pay_core && (periodic_pay_transfer || periodic_pay_close)) {
		return 1, nil
	}
	return 0, fmt.Error("failed condition")

}