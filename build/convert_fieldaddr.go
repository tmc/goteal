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

	// check for globals, txn, or gtxn access
	// fmt.Println("fieldaddr val:", i.Name(), i.String())
	// fmt.Println("fieldaddr x:", i.X)
	// fmt.Println("fieldaddr f:", i.Field)

	// handle global references
	// fmt.Println("global ref?", i.Name())
	// fmt.Println(i.X.String() == "local github.com/tmc/goteal/types.Globals (globals)")
	if i.X.String() == "local github.com/tmc/goteal/types.Globals (globals)" {
		typ := i.X.Type()
		pType, ok := typ.Underlying().(*types.Pointer)
		if ok {
			structType, ok := pType.Elem().Underlying().(*types.Struct)
			if ok {
				fld := structType.Field(i.Field)
				// fmt.Println("fieldaddr f:", structType.Field(i.Field))
				// result.AppendLine(fmt.Sprintf("global %v", fld.Name()))
				b.resolved[fmt.Sprintf("*%v", i.Name())] = fmt.Sprintf("global %v", fld.Name())
				// b.resolved[fmt.Sprintf("*%v", i.Name())] = i.X
				return nil
			}
		}
	}
	return fmt.Errorf("coult not resolve field addr %v", i)

	// for _, arg := range i.Call.Args {
	// 	if c, ok := arg.(*ssa.Const); ok {
	// 		fmt.Println(c)
	// 		result.AppendLine(fmt.Sprintf("%s %s", arg.Type(), c.Value.ExactString()))
	// 	} else {
	// 		result.AppendLine(fmt.Sprintf(" // err: unknown arg type %T", arg))
	// 	}
	// }
	//result.AppendLine(fmt.Sprintf("callsub %s", i.Common().Value.Name()))
}
