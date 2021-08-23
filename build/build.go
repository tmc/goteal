package build

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/tmc/goteal/teal"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

// Builder handles the loading and compilation of a Go program to TEAL.
type Builder struct {
	DebugLevel int
	DumpSSA    bool

	pkgCfg *packages.Config

	program *ssa.Program
	pkgs    []*ssa.Package
	phis    []*Phi

	hasStartedProcessingContract bool

	resolved map[string]interface{}
}

type Phi struct {
	*ssa.Phi
	i ssa.Instruction

	// only populated in some cases
	v ssa.Value
}

// New returns a new Builder.
func New() *Builder {
	return &Builder{
		pkgCfg: &packages.Config{
			Mode: packages.LoadSyntax,
		},
		resolved: map[string]interface{}{},
	}
}

// LoadSources loads go source packages and files to prepare for compilation.
func (b *Builder) LoadSources(sources ...string) error {
	initial, err := packages.Load(b.pkgCfg, sources...)
	if err != nil {
		return err
	}
	if len(initial) == 0 {
		return fmt.Errorf("no packages")
	}
	if packages.PrintErrors(initial) > 0 {
		return fmt.Errorf("packages contain errors")
	}
	mode := ssa.BuilderMode(0)
	// mode := ssa.PrintPackages
	b.program, b.pkgs = ssautil.Packages(initial, mode)
	return nil
}

// Build assembles a TEAL program from a Go program.
func (b *Builder) Build() (*teal.Program, error) {
	for _, pkg := range b.pkgs {
		pkg.Build()
		if b.DumpSSA {
			buf := new(bytes.Buffer)
			ssa.WritePackage(buf, pkg)
			io.Copy(os.Stdout, buf)
		}

		// check for Contract()
		hasContract, err := packageDefinesContract(pkg)
		if err != nil {
			return nil, fmt.Errorf("supplied package does not export the correct Contract() signature. See https://github.com/tmc/goteal#constraints for details.\n%w", err)
		}
		if hasContract {
			return b.convertSSAToTEAL(pkg)
		}
	}
	return nil, fmt.Errorf("supplied package does not export the correct Contract() signature. See https://github.com/tmc/goteal#constraints for details")
}
