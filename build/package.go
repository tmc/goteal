package build

import (
	"fmt"

	"golang.org/x/tools/go/ssa"
)

// ExpectedContractType is the expected function signature that a package implements to indicate that it defines an Algorand Smart Contract.
const ExpectedContractType = "func(globals github.com/tmc/goteal/types.Globals, txn github.com/tmc/goteal/types.Transaction, gtxn github.com/tmc/goteal/types.TxGroup) (int, error)"

func packageDefinesContract(pkg *ssa.Package) (bool, error) {
	contractFn, ok := pkg.Members["Contract"]
	if !ok {
		return false, nil
	}
	typ := contractFn.Type().String()
	if typ == ExpectedContractType {
		return true, nil
	}
	return false, fmt.Errorf("expected\n	%v\ngot\n	%v", ExpectedContractType, typ)
}
