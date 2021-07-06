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
	result.AppendLine("b main")

	var members []ssa.Member
	for _, m := range pkg.Members {
		members = append(members, m)
	}

	sort.Slice(members, func(i, j int) bool {
		return members[i].Pos() < members[j].Pos()
	})
	for _, m := range members {
		if b.Debug {
			fmt.Println(" member:", m.Name())
		}
		// TODO: support init
		if m.Name() == "init" {
			continue
		}

		var err error
		switch m := m.(type) {
		case *ssa.Function:
			err = b.convertSSAFunctionToTEAL(result, m)
		default:
			if b.Debug {
				fmt.Fprintln(os.Stderr, fmt.Sprintf("unhandled type %T", m))
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

func (b *Builder) convertSSAFunctionToTEAL(result *teal.Program, m *ssa.Function) error {
	result.AppendLine("")
	// TODO: extract and preserve function comment.
	result.AppendLine(fmt.Sprintf("%s: // %v", m.Name(), m.Signature))
	// fmt.Fprintln(os.Stderr, "fn to teal:", m)
	// fmt.Fprintln(os.Stderr, " params :", m.Params)
	for _, block := range m.Blocks {
		// fmt.Fprintln(os.Stderr, " block :", block.String())

		if block.Comment != "entry" {
			result.AppendLine(fmt.Sprintf("// %v", block.Comment))
		}

		for _, instr := range block.Instrs {
			result.AppendLine(fmt.Sprintf("// %v", instr))
			if err := b.convertSSAInstructionToTEAL(result, instr); err != nil {

			}
		}

	}
	return nil
}
func (b *Builder) convertSSAInstructionToTEAL(result *teal.Program, i ssa.Instruction) error {
	switch i := i.(type) {
	case *ssa.BinOp:
		result.AppendLine(fmt.Sprintf("%v", i.Op))
	case *ssa.Call:
		result.AppendLine(fmt.Sprintf("callsub %s", i.Common().Value.Name()))
	case *ssa.Return:
		if i.Parent().Name() == "main" {
			result.AppendLine("return")
		} else {
			result.AppendLine("retsub")
		}
	default:
		result.AppendLine(fmt.Sprintf("// convertSSAInstructionToTEAL: unexpected type: %T - %v", i, i))
	}
	return nil
}
