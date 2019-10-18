package tools

import (
	"github.com/fatih/color"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Default is the default spell which will be invocated by mage
var Default = Build

// Build is an invocation of multiple spell
func Build() {
	color.Red("# Installing tools ---------------------------------------------------------")
	mg.SerialDeps(Go.Vendor, Go.Tools)
}

// Go is a spell category
type Go mg.Namespace

// Vendor spell will download dependencies and vendor them
func (Go) Vendor() error {
	color.Blue("## Vendoring dependencies")
	return sh.RunV("go", "mod", "vendor")
}

// Tools spell will build a binarie for all tools
func (Go) Tools() error {
	color.Blue("## Intalling tools")
	return sh.RunV("go", "run", "github.com/izumin5210/gex/cmd/gex", "--build")
}
