package build

import (
	"fmt"
	"strings"

	"github.com/tmc/goteal/teal"
	"golang.org/x/tools/go/ssa"
)

func (b *Builder) convertSSAFunctionToTEAL(result *teal.Program, m *ssa.Function) error {
	result.AppendLine("")

	name := strings.ToLower(m.Name())
	// TODO: extract and preserve function comment.
	if b.DebugLevel > 0 {
		result.AppendLine(fmt.Sprintf("%s: // %v", name, m.Signature))
	} else {
		result.AppendLine(fmt.Sprintf("%s:", name))
	}

	var processingContract bool
	if m.Name() == "Contract" && !b.hasStartedProcessingContract {
		processingContract = true
		b.hasStartedProcessingContract = true
	}
	// fmt.Fprintln(os.Stderr, "fn to teal:", m)
	// fmt.Fprintln(os.Stderr, " params :", m.Params)
	for blockIndex, block := range m.Blocks {
		ctx := ConvertContext{BlockIndex: blockIndex}
		if b.DebugLevel > 0 {
			// fmt.Fprintln(os.Stderr, " block :", block.String())
			// result.AppendLine(fmt.Sprintf("// block: %v", block))
		}

		if block.Comment != "entry" {
			result.AppendLine(fmt.Sprintf("// block %v: %v", blockIndex, block.Comment))
			result.AppendLine(fmt.Sprintf("%v.block.%v:", name, blockIndex))
		}

		if processingContract && blockIndex == 0 {
			// If we're processing the primary Contract function, skip over the first few instructions that reference function arguments.
			// result.AppendLine(fmt.Sprintf("// started proceesing %v", blockIndex))
			// block.Instrs = stripInitialExpectedInstructions(block.Instrs)
		}

		for _, instr := range block.Instrs {
			if err := b.convertSSAInstructionToTEAL(ctx, result, instr); err != nil {
				return err
			}
		}

		// look up if a Phi node refers to this block and store a value if so.

		phis := phiReferrers(block, b.phis)
		for _, phi := range phis {
			if b.DebugLevel > 0 {
				result.AppendLine(fmt.Sprintf("// store phi node %v %v", phi.Name(), phi.v))
			}
			phiReg := strings.TrimLeft(phi.Name(), "t")
			result.AppendLine(b.resolve(phi.v))
			result.AppendLine(fmt.Sprintf("store %v", phiReg))
		}
		if b.DebugLevel > 1 {
			result.AppendLine(fmt.Sprintf("// endblock: %v", block))
		}
	}
	return nil
}

func stripInitialExpectedInstructions(insts []ssa.Instruction) []ssa.Instruction {
	var result []ssa.Instruction
	/*
		local github.com/tmc/goteal/types.Globals (globals)
		*t0 = globals
		local github.com/tmc/goteal/types.Transaction (txn)
		*t1 = txn
	*/
	for i, inst := range insts {
		if i < 4 && inst.Parent().Type().String() == ExpectedContractType {
			// fmt.Fprintln(os.Stderr, "skippping inst:", i, inst)
			// TODO: add validation here tha thte expected instructions are the ones being skipped
			// TODO: does this depend on in-method use? gtxn for example.
			continue
		}
		result = append(result, inst)
	}
	return result
}
