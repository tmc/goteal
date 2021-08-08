package recurring_swap

import (
	"fmt"

	"github.com/tmc/goteal/avm"
	"github.com/tmc/goteal/types"
)

const (
	TmplBuyer    = types.Address("6ZHGHH5Z5CTPCF5WCESXMGRSVK7QJETR63M3NY5FJCUYDHO57VTCMJOBGY")
	TmplProvider = types.Address("7Z5PWO2C6LFNQFGHWKSK5H47IQP5OJW2M3HA2QPXTY3WTNP5NU2MHBW27M")
	TmplAmount   = 100000
	TmplFee      = 1000
	TmplTimeout  = 100000
)

func Contract(globals types.Globals, txn types.Transaction, gtxn types.TxGroup) (int, error) {
	feeCond := txn.Fee < TmplFee

	// set type cond
	typeCond := (txn.TypeEnum == 1) && (txn.CloseRemainderTo == globals.ZeroAddress)

	// set receive conditions
	isCloset := txn.CloseRemainderTo == globals.ZeroAddress
	isReceiverSet := txn.Receiver == TmplProvider
	isAmountCorrect := txn.Amount == TmplAmount

	// needs check here on how to implement various utilities in avm/types\
	// TMC - thoughts on pyteals Arg(0) for edwards->itob
	//https://github.com/algorand/pyteal/blob/master/examples/signature/recurring_swap.py
	isEdwards := avm.Ed25519verify(avm.Itob(txn.FirstValid, 0, TmplProvider))
	isLease := txn.Lease == avm.Sha256(avm.Itob(txn.FirstValid))

	// set receive condition
	recCond := isCloset && isReceiverSet && isAmountCorrect && isEdwards && isLease

	//set close condtion
	closeCond := (txn.CloseRemainderTo == TmplBuyer) && (txn.Amount == 0) && (txn.FirstValid >= TmplTimeout)

	if feeCond && typeCond && (recCond || closeCond) {
		return 1, nil
	}
	return 0, fmt.Errorf("failed condition")

}
