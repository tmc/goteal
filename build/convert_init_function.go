package build

import (
	"fmt"
	"strings"

	"github.com/tmc/goteal/teal"
	"golang.org/x/tools/go/ssa"
)

func (b *Builder) convertSSAInitFunctionToTEAL(result *teal.Program, m *ssa.Function) error {
	result.AppendLine("")

	name := strings.ToLower(m.Name())
	// TODO: extract and preserve function comment.
	if b.Debug {
		result.AppendLine(fmt.Sprintf("// processing %v", name))
	}

	// fmt.Fprintln(os.Stderr, "fn to teal:", m)
	// fmt.Fprintln(os.Stderr, " params :", m.Params)
	for blockIndex, block := range m.Blocks {
		ctx := ConvertContext{
			BlockIndex: blockIndex,
			IsInit:     true,
		}
		if b.Debug {
			// fmt.Fprintln(os.Stderr, " block :", block.String())
			// result.AppendLine(fmt.Sprintf("// block: %v", block))
		}

		if blockIndex == 0 {
			continue
		}
		if blockIndex == 1 {
			block.Instrs = stripInitialExpectedInitInstructions(block.Instrs)
		}

		for _, instr := range block.Instrs {
			if err := b.convertSSAInstructionToTEAL(ctx, result, instr); err != nil {
				return err
			}
		}

		if b.Debug {
			// result.AppendLine(fmt.Sprintf("// endblock: %v", block))
		}
	}
	return nil
}

func stripInitialExpectedInitInstructions(insts []ssa.Instruction) []ssa.Instruction {
	var result []ssa.Instruction
	for i, inst := range insts {
		if i < 1 {
			//fmt.Fprintln(os.Stderr, "skippping inst:", i, inst)
			// TODO: add validation here tha thte expected instructions are the ones being skipped
			// TODO: does this depend on in-method use? gtxn for example.
			continue
		}
		result = append(result, inst)
	}
	return result
}
