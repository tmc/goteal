package split

import (
	"fmt"

	"github.com/tmc/goteal/types"
)

const (
	Fee              = 1000
	RCV1             = types.Address("6ZHGHH5Z5CTPCF5WCESXMGRSVK7QJETR63M3NY5FJCUYDHO57VTCMJOBGY")
	RCV2             = types.Address("7Z5PWO2C6LFNQFGHWKSK5H47IQP5OJW2M3HA2QPXTY3WTNP5NU2MHBW27M")
	RCVOwner         = types.Address("5MK5NGBRT5RL6IGUSYDIX5P7TNNZKRVXKT6FGVI6UVK6IZAWTYQGE4RZIQ")
	RatioNumerator   = 1
	RatioDenominator = 3
	MinPay           = 1000
	Timeout          = 3000
)

func Contract(globals types.Globals, txn types.Transaction, gtxn types.TxGroup) (int, error) {

	//split core inner conditions
	isPayment := txn.TypeEnum == types.PaymentTx
	isFeeAppropriate := txn.Fee < Fee
	isRekeySet := txn.RekeyTo == globals.ZeroAddress
	// set split core bool
	splitCore := isPayment && isFeeAppropriate && isRekeySet

	//split transfer inner conditions
	isSameSender := gtxn[0].Sender == gtxn[1].Sender
	isCloseSet := txn.CloseRemainderTo == globals.ZeroAddress
	isReceiverOneTmplReceiverOne := gtxn[0].Receiver == RCV1
	isReceiverTwoTmplReceiverTwo := gtxn[1] == RCV2
	splitSumEqualsBalance := gtxn[0].Amount == ((gtxn[0].Amount + gtxn[1].Amount) * RatioNumerator / RatioDenominator)
	hasMinBalance := gtxn[0].Amount == MinPay
	// set split transfer
	splitTransfer := isSameSender && isCloseSet && isReceiverOneTmplReceiverOne && isReceiverTwoTmplReceiverTwo && splitSumEqualsBalance && hasMinBalance

	//split close inner conditions
	isRemainderSet := txn.CloseRemainderTo == RCVOwner
	isReceiverZeroSet := txn.Receiver == globals.ZeroAddress
	isBalanceZeroNow := txn.Amount == 0
	isNotTimedOut := txn.FirstValid > Timeout

	// set split close
	splitClose := isRemainderSet && isReceiverZeroSet && isBalanceZeroNow && isNotTimedOut

	if globals.GroupSize == 2 {
		if splitCore && splitTransfer && splitClose {
			return 1, nil
		}
	} else {
		if splitCore {
			return 1, nil
		}
	}
	return 0, fmt.Errorf("failed condition")

}
