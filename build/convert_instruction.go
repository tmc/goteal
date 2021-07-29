package build

import (
	"fmt"
	"go/types"
	"strings"

	"github.com/tmc/goteal/teal"
	"golang.org/x/tools/go/ssa"
)

var specialCaseInstructions = map[string]func(result *teal.Program, i ssa.Instruction) error{
	"local github.com/tmc/goteal/types.Globals (globals)": func(result *teal.Program, i ssa.Instruction) error {
		return nil
		// rands := i.Operands(nil)
		// result.AppendLine(fmt.Sprintf("global %v", rands))
		// return nil
	},
}

func (b *Builder) convertSSAInstructionToTEAL(result *teal.Program, blockIndex int, i ssa.Instruction) error {
	// result.AppendLine(fmt.Sprintf("\n// inst: %v", i))
	if handler, ok := specialCaseInstructions[i.String()]; ok {
		return handler(result, i)
	}
	switch i := i.(type) {
	case *ssa.BinOp:
		if b.Debug {
			result.AppendLine(fmt.Sprintf("// binop: %v = %v", i.Name(), i))
		}
		result.AppendLine(b.resolve(i.X))
		result.AppendLine(b.resolve(i.Y))
		result.AppendLine(fmt.Sprintf("%v", i.Op))
		// b.resolved[i.Name()] =
	case *ssa.FieldAddr:
		if b.Debug {
			result.AppendLine(fmt.Sprintf("// fieldaddr: %v = %v", i.Name(), i))
		}
		if err := b.convertSSAFieldAddrToTEAL(result, i); err != nil {
			return fmt.Errorf("issue converting field addr: %w", err)
		}
	case *ssa.Call:
		if b.Debug {
			result.AppendLine(fmt.Sprintf("// call: %v = %v", i.Name(), i))
		}
		if err := b.convertSSACallToTEAL(result, i); err != nil {
			return fmt.Errorf("issue converting call: %w", err)
		}
	case *ssa.Store:
		if b.Debug {
			result.AppendLine(fmt.Sprintf("// store: %v", i))
		}
		b.resolved[i.Addr.Name()] = i.Val
	case *ssa.UnOp:
		if b.Debug {
			result.AppendLine(fmt.Sprintf("// unop: %v = %v", i.Name(), i))
			result.AppendLine(fmt.Sprintf("// unop: x: %v", i.X))
			result.AppendLine(fmt.Sprintf("// unop: x type: %v", i.X.Type()))
		}
		b.resolved[i.Name()] = b.resolve(i.X)
	case *ssa.Alloc:
		if b.Debug {
			result.AppendLine(fmt.Sprintf("// alloc: %v = %v", i.Name(), i))
		}
		b.resolved[i.Name()] = i
	case *ssa.If:
		if b.Debug {
			result.AppendLine(fmt.Sprintf("// if: %v", i))
		}
		rands := i.Operands(nil)
		for opi, rand := range rands {
			//result.AppendLine(fmt.Sprintf("// if: %v, operand %v: %v", i, opi, *rand))
			_, _ = opi, rand
		}
		// for _, ref := range *i.Referrers() {
		// 	result.AppendLine(fmt.Sprintf("// ref: %v", ref))
		// }
		result.AppendLine(fmt.Sprintf("bnz %v", strings.ToLower(i.Block().Parent().Name())+fmt.Sprintf(".block.%v", blockIndex+1)))
		result.AppendLine(fmt.Sprintf("b %v", strings.ToLower(i.Block().Parent().Name())+fmt.Sprintf(".block.%v", blockIndex+2)))

	case *ssa.Return:
		rands := i.Operands(nil)
		result.AppendLine(b.resolve(*rands[0]))
		// TODO: this is a bit naive
		if i.Parent().Name() == "Contract" {
			result.AppendLine("return")
		} else {
			result.AppendLine("retsub")
		}
	default:
		result.AppendLine(fmt.Sprintf("// convertSSAInstructionToTEAL: unexpected type: %T - %v", i, i))
	}
	return nil
}

func (b *Builder) resolve(x ssa.Value) string {
	// fmt.Println("checking for", x.Name())
	// fmt.Println(b.resolved)
	r, ok := b.resolved[x.Name()]
	if ok {
		if s, ok := r.(string); ok {
			return s
		}
		if v, ok := r.(ssa.Value); ok {
			return v.String()
		}
		return "err // issue with resolved value"
	}

	// fmt.Println("resolve:", x.Name(), x.Type(), x.String())
	// fmt.Printf("typ: %T\n", x.Type())

	prebaked := map[string]string{
		"&t0.GroupSize [#4]": "global GroupSize",
	}
	if p, ok := prebaked[x.String()]; ok {
		return p
	}

	switch t := x.Type().(type) {
	case *types.Pointer:
		return fmt.Sprintf("pointer %v", x)
	case *types.Basic:
		val := strings.Split(x.String(), ":")[0]
		return fmt.Sprintf("int %v", val)
	default:
		return fmt.Sprintf("err // failed to resolve %v", t)
	}
}
