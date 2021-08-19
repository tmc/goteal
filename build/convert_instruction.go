package build

import (
	"fmt"
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

// ConvertContext carries around a bit of state regarding what conversion is being carried out.
type ConvertContext struct {
	// BlockIndex encodes the position of the current block in the surrounding function.
	BlockIndex int

	// IsInit encodes whether or not we are in an init block which is handled with some special casing.
	IsInit bool
}

func (b *Builder) convertSSAInstructionToTEAL(ctx ConvertContext, result *teal.Program, i ssa.Instruction) error {
	if b.DebugLevel > 0 {
		result.AppendLine(fmt.Sprintf("// inst: %v %T", i, i))
	}
	if handler, ok := specialCaseInstructions[i.String()]; ok {
		return handler(result, i)
	}
	switch i := i.(type) {
	case *ssa.BinOp:
		if b.DebugLevel > 0 {
			result.AppendLine(fmt.Sprintf("// binop: %v = %v", i.Name(), i))
		}
		// spew.Dump("x:", i.Parent().ValueForExpr(i.X.Pos()))
		// spew.Dump("y:", i.Parent().ValueForExpr(i.X.Pos()))
		result.AppendLine(b.resolve(i.X))
		result.AppendLine(b.resolve(i.Y))
		result.AppendLine(fmt.Sprintf("%v", i.Op))
		// b.resolved[i.Name()] =
	case *ssa.FieldAddr:
		if b.DebugLevel > 0 {
			result.AppendLine(fmt.Sprintf("// fieldaddr: %v = %v", i.Name(), i))
		}
		if err := b.convertSSAFieldAddrToTEAL(result, i); err != nil {
			return fmt.Errorf("issue converting field addr: %w", err)
		}
	case *ssa.Call:
		if err := b.convertSSACallToTEAL(ctx, result, i); err != nil {
			return fmt.Errorf("issue converting call: %w", err)
		}
	case *ssa.Store:
		if b.DebugLevel > 0 {
			result.AppendLine(fmt.Sprintf("// store: %v", i))
		}
		b.resolved[i.Addr.Name()] = i.Val
		if b.DebugLevel > 1 {
			result.AppendLine(fmt.Sprintf("// stored: %q", b.resolved))
		}
	case *ssa.UnOp:
		if b.DebugLevel > 0 {
			result.AppendLine(fmt.Sprintf("// unop: %v = %v", i.Name(), i))
			result.AppendLine(fmt.Sprintf("// unop: x: %v", i.X))
			result.AppendLine(fmt.Sprintf("// unop: x type: %v", i.X.Type()))
		}
		b.resolved[i.Name()] = b.resolve(i.X)
	case *ssa.Alloc:
		if b.DebugLevel > 0 {
			result.AppendLine(fmt.Sprintf("// alloc: %v = %v", i.Name(), i))
		}
		b.resolved[i.Name()] = i
	case *ssa.If:
		if b.DebugLevel > 0 {
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
		result.AppendLine(fmt.Sprintf("bnz %v", strings.ToLower(i.Block().Parent().Name())+fmt.Sprintf(".block.%v", ctx.BlockIndex+1)))
		result.AppendLine(fmt.Sprintf("b %v", strings.ToLower(i.Block().Parent().Name())+fmt.Sprintf(".block.%v", ctx.BlockIndex+2)))

	case *ssa.Return:
		rands := i.Operands(nil)

		if ctx.IsInit && len(rands) == 0 {
			return nil
		}

		if len(rands) == 1 || len(rands) == 2 {
			result.AppendLine(b.resolve(*rands[0]))
		} else {
			result.AppendLine(fmt.Sprintf("// unexpected number of return operands: %v", len(rands)))
		}
		// TODO: this is a bit naive
		if i.Parent().Name() == "Contract" {
			result.AppendLine("return")
		} else {
			result.AppendLine("retsub")
		}
	case *ssa.Convert:
		if b.DebugLevel > 0 {
			result.AppendLine(fmt.Sprintf("// convert: %v to %v", i.X, i.Type()))
		}
		result.AppendLine(fmt.Sprintf("// convert: %v to %v", i.X, i.Type()))
	case *ssa.Jump:
		if ctx.IsInit {
			return nil
		}
		if b.DebugLevel > 0 {
			result.AppendLine(fmt.Sprintf("// jump: %v", i))
		}
		fnName := strings.ToLower(i.Parent().Name())
		result.AppendLine(fmt.Sprintf("b %v.block.%v", fnName, i.Block().Succs[0].Index))
	case *ssa.Phi:
		if b.DebugLevel > 0 {
			result.AppendLine(fmt.Sprintf("// ɸ %v %v", i.Name(), i))
		}
		// phiReg := strings.TrimLeft(i.Name(), "t")
		// result.AppendLine(fmt.Sprintf("load %v", phiReg))
	default:
		result.AppendLine(fmt.Sprintf("// convertSSAInstructionToTEAL: unexpected type: %T - %v", i, i))
	}
	return nil
}
