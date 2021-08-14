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
		// fmt.Println("sorting:", i, j, "pos:", members[i].Pos(), members[j].Pos())
		// return strings.Compare(members[i].Name(), members[j].Name()) < 0
		return members[i].Pos() <= members[j].Pos()
	})
	for _, m := range members {
		// TODO: support init guarding
		if m.Name() == "init$guard" {
			continue
		}
		if b.Debug {
			fmt.Printf(" > member: %s: %v: %T\n", m.Name(), m, m.Type())
		}
		var err error
		switch m := m.(type) {
		case *ssa.Function:
			// init is processed in a special fashion as it is where global vars are set.

			if m.Name() == "init" {
				err = b.convertSSAInitFunctionToTEAL(result, m)
			} else {
				err = b.convertSSAFunctionToTEAL(result, m)
			}
		case *ssa.Global:
			//fmt.Println("got a global:", m.Name(), m.Type())
			if _, ok := b.resolved[m.Name()]; !ok {
				b.resolved[m.Name()] = nil
			}
			// fmt.Println("val ?:", m.Type().Underlying().(*types.Pointer).Elem().Underlying())

			/*
				fmt.Println("deets:")
				fmt.Println(m.Name())
				fmt.Println(m.Object())
				// fmt.Println(m.Operands(r))
				fmt.Println(m.Package())
				fmt.Println(m.Parent())
				fmt.Println(m.Pos())
				fmt.Println(m.Referrers())
				// fmt.Println(m.RelString(f))
				fmt.Println(m.String())
				fmt.Println(m.Token())
				fmt.Println(m.Type())
			*/

			// fset := pkg.Prog.Fset
			// f, err := parser.ParseFile(fset, "src.go", src, 0)
			// if err != nil {
			// panic(err)
			// }
			// ast.Inspect(pkg.Prog.Fset, func(n ast.Node) bool {
			// 	var s string
			// 	switch x := n.(type) {
			// 	case *ast.BasicLit:
			// 		s = x.Value
			// 	case *ast.Ident:
			// 		s = x.Name
			// 	}
			// 	if s != "" {
			// 		fmt.Printf("%s:\t%s\n", fset.Position(n.Pos()), s)
			// 	}
			// 	return true
			// })

			// gPos := m.Pos()
			// P2 := pkg.Prog.Fset.Position(gPos)
			// spew.Dump(gPos, P2)
			if b.Debug {
				result.AppendLine(fmt.Sprintf("// global var: %v %v", m, m.Object()))
			}
			// v, ok := m.Object().(*types.Var)
			// if ok {
			// 	// result.AppendLine(fmt.Sprintf("// le var: %v", v))
			// 	// result.AppendLine(fmt.Sprintf("// le st: %v", v.String()))
			// 	// x, err := pkg.Prog.VarValue(v, pkg, nil)

			// 	// fm := pkg.Members[v.Name()]
			// 	// result.AppendLine(fmt.Sprintf("// from members: %v", fm))
			// 	// fv := pkg.Var(v.Name())
			// 	// result.AppendLine(fmt.Sprintf("// from var(): %v", fv))
			// 	// result.AppendLine(fmt.Sprintf("// le var: %v %T", x, x))
			// 	// result.AppendLine(fmt.Sprintf("// le err?: %v", err))
			// }
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
			err = fmt.Errorf("unhandled type %T", m)
		}
		if err != nil {
			return nil, fmt.Errorf("issue converting %v: %w", m, err)
		}
	}
	return result, nil
}
