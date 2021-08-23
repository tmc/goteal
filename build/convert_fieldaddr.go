package build

import (
	"fmt"
	"go/types"

	"github.com/tmc/goteal/teal"
	"golang.org/x/tools/go/ssa"
)

func (b *Builder) convertSSAFieldAddrToTEAL(result *teal.Program, i *ssa.FieldAddr) error {

	// if handler, ok := specialCases[i.Call.Value.String()]; ok {
	// 	return handler(result, i)
	// }
	var fld *types.Var
	typ := i.X.Type()
	pType, ok := typ.Underlying().(*types.Pointer)
	if ok {
		structType, ok := pType.Elem().Underlying().(*types.Struct)
		if ok {
			fld = structType.Field(i.Field)
		}
	}

	// check for globals, txn, or gtxn access
	// fmt.Println("fieldaddr val:", i.Name(), i.String())
	// fmt.Println("fieldaddr x:", i.X)
	// fmt.Println("fieldaddr f:", i.Field)

	// handle global references
	// fmt.Println("global ref?", i.Name())
	// fmt.Println(i.X.String() == "local github.com/tmc/goteal/types.Globals (globals)")
	if i.X.String() == "local github.com/tmc/goteal/types.Globals (globals)" {
		b.resolved[fmt.Sprintf("*%v", i.Name())] = fmt.Sprintf("global %v", fld.Name())
		return nil
	}
	// TODO: there's probably a cleaner way to approach this
	if i.X.String() == "local github.com/tmc/goteal/types.Transaction (txn)" {
		b.resolved[fmt.Sprintf("*%v", i.Name())] = fmt.Sprintf("txn %v", fld.Name())
		return nil
	}
	if i.X.String() == "local github.com/tmc/goteal/types.TxGroup (gtxn)" {
		b.resolved[fmt.Sprintf("*%v", i.Name())] = fmt.Sprintf("gtxn %v", fld.Name())
		return nil
	}
	return fmt.Errorf("could not resolve field addr %v %v", i, i.X)
}
