package build

import (
	"fmt"

	"github.com/tmc/goteal/teal"
	"golang.org/x/tools/go/ssa"
)

var specialCases = map[string]func(result *teal.Program, i *ssa.Call) error{
	"fmt.Errorf": func(result *teal.Program, i *ssa.Call) error {
		// result.AppendLine(fmt.Sprintf("// %v", i.Call.Args[0]))
		return nil
	},
}

func (b *Builder) convertSSACallToTEAL(result *teal.Program, i *ssa.Call) error {

	if handler, ok := specialCases[i.Call.Value.String()]; ok {
		return handler(result, i)
	}

	// fmt.Println("call val:", i.Name(), i.Call.Value)
	for _, arg := range i.Call.Args {
		if c, ok := arg.(*ssa.Const); ok {
			fmt.Println(c)
			result.AppendLine(fmt.Sprintf("%s %s", arg.Type(), c.Value.ExactString()))
		} else {
			result.AppendLine(fmt.Sprintf(" // err: unknown arg type %T", arg))
		}
	}
	result.AppendLine(fmt.Sprintf("callsub %s", i.Common().Value.Name()))
	return nil
}
