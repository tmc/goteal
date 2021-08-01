package build

import (
	"fmt"
	"os"
	"sort"

	"github.com/tmc/goteal/teal"
	"golang.org/x/tools/go/ssa"
)

// convertSSAToTEAL converts a Go program to TEAL.
func (b *Builder) convertSSAToTEAL(pkg *ssa.Package) (*teal.Program, error) {
	result := &teal.Program{}

	result.AppendLine("#pragma version 4")
	result.AppendLine("// jump to main")
	result.AppendLine("b contract")

	var members []ssa.Member
	for _, m := range pkg.Members {
		members = append(members, m)
	}

	sort.Slice(members, func(i, j int) bool {
		return members[i].Pos() < members[j].Pos()
	})
	for _, m := range members {
		// TODO: support init
		if m.Name() == "init" || m.Name() == "init$guard" {
			continue
		}
		if b.Debug {
			fmt.Println(" > member:", m.Name(), m.Type())
		}

		var err error
		switch m := m.(type) {
		case *ssa.Function:
			err = b.convertSSAFunctionToTEAL(result, m)
		case *ssa.Global:
			if b.Debug {
				result.AppendLine(fmt.Sprintf("// global var: %v %v", m, m.Object()))
			}
			// TODO: how to get the actual value to add to b.resolved ?
			//b.resolved[m.Name()] = m..
		case *ssa.NamedConst:
			if b.Debug {
				result.AppendLine(fmt.Sprintf("// named const: %v = %v", m.Name(), m.Value))
			}
			b.resolved[m.Name()] = m.Value
		default:
			if b.Debug {
				fmt.Fprintln(os.Stderr, fmt.Sprintf(" > unhandled type %T", m))
			}
			//err = fmt.Errorf("unhandled type %T", m)
		}
		if err != nil {
			return nil, fmt.Errorf("issue converting %v: %w", m, err)
		}
	}
	// TODO: switch on *NamedConst, *Global, *Function, or *Type
	return result, nil
}
