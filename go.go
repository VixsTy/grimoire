package grimoire

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Go is a mage namespace which will manage golang actions
type Go mg.Namespace

// Tidy add/remove depenedencies.
func (Go) Tidy() error {
	fmt.Println("## Cleaning go modules")
	return sh.RunV("go", "mod", "tidy", "-v")
}

// Deps install dependency tools.
func (Go) Deps() error {
	color.Cyan("## Vendoring dependencies")
	return sh.RunV("go", "mod", "vendor")
}

// License check that project dont used blacklisted licenses
func (Go) License() error {
	color.Cyan("## Check license")
	return sh.RunV("wwhrd", "check")
}

// Format runs gofmt on everything
func (Go) Format() error {
	color.Cyan("## Format everything")
	args := []string{"-s", "-w"}
	args = append(args, goFiles...)
	return sh.RunV("gofumpt", args...)
}

// Import runs goimports on everything
func (Go) Import() error {
	color.Cyan("## Process imports")
	args := []string{"-w"}
	args = append(args, goSrcFiles...)
	return sh.RunV("goreturns", args...)
}

// Lint run linter.
func (Go) Lint() error {
	mg.Deps(Go.Format)
	color.Cyan("## Lint go code")
	return sh.RunV("golangci-lint", "run")
}
