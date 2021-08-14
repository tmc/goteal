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

// ConvertContext carries around a bit of state regarding what conversion is being carried out.
type ConvertContext struct {
	// BlockIndex encodes the position of the current block in the surrounding function.
	BlockIndex int

	// IsInit encodes whether or not we are in an init block which is handled with some special casing.
	IsInit bool
}

func (b *Builder) convertSSAInstructionToTEAL(ctx ConvertContext, result *teal.Program, i ssa.Instruction) error {
	// result.AppendLine(fmt.Sprintf("\n// inst: %v", i))
	if handler, ok := specialCaseInstructions[i.String()]; ok {
		return handler(result, i)
	}
	switch i := i.(type) {
	case *ssa.BinOp:
		if b.Debug {
			result.AppendLine(fmt.Sprintf("// binop: %v = %v", i.Name(), i))
		}
		// spew.Dump("x:", i.Parent().ValueForExpr(i.X.Pos()))
		// spew.Dump("y:", i.Parent().ValueForExpr(i.X.Pos()))
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
		if err := b.convertSSACallToTEAL(ctx, result, i); err != nil {
			return fmt.Errorf("issue converting call: %w", err)
		}
	case *ssa.Store:
		if b.Debug {
			result.AppendLine(fmt.Sprintf("// store: %v", i))
		}
		b.resolved[i.Addr.Name()] = i.Val
		if b.Debug {
			result.AppendLine(fmt.Sprintf("// stored: %q", b.resolved))
		}
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
		if b.Debug {
			result.AppendLine(fmt.Sprintf("// convert: %v to %v", i.X, i.Type()))
		}
		result.AppendLine(fmt.Sprintf("// convert: %v to %v", i.X, i.Type()))
	case *ssa.Jump:
		if ctx.IsInit {
			return nil
		}
		if b.Debug {
			result.AppendLine(fmt.Sprintf("// jump: %v", i))
		}
	default:
		result.AppendLine(fmt.Sprintf("// convertSSAInstructionToTEAL: unexpected type: %T - %v", i, i))
	}
	return nil
}

// resolve produces a TEAL-compatible representation of an ssa.Value.
func (b *Builder) resolve(v ssa.Value) string {
	if b.Debug {
		fmt.Println("-> resolve()", v.Name())
	}
	//fmt.Println(b.resolved)
	value := v
	// first check if we already have a value resolved
	r, ok := b.resolved[value.Name()]
	if ok {
		if s, ok := r.(string); ok {
			return s
		}
		value, ok = r.(ssa.Value)
		if !ok {
			return fmt.Sprintf("err // issue with resolved value: %v %v %T", value.Name(), r, r)
		}
	}
	switch v := value.(type) {
	case *ssa.Const:
		return fmt.Sprintf("%v %v", v.Type().Underlying(), v.Value.ExactString())
		// default:
		//return "err // unkown type " + fmt.Sprintf("%T", v)
	}

	sparts := strings.Split(value.String(), " ")
	sp1 := sparts[0]
	dotParts := strings.Split(sp1, ".")

	// TODO: this is pretty gross special-casing.

	fieldAccessRoots := map[string]string{
		"&t0": "global",
		"&t1": "txn",
	}
	if fa, ok := fieldAccessRoots[dotParts[0]]; ok {
		return fmt.Sprintf("%v %v", fa, dotParts[1])
	}

	switch t := value.Type().(type) {
	case *types.Pointer:
		return fmt.Sprintf("pointer %v", value)
	case *types.Basic:
		val := strings.Split(value.String(), ":")[0]
		return fmt.Sprintf("int %v", val)
	default:
		return fmt.Sprintf("err // failed to resolve %v", t)
	}
}
