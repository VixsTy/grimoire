package grimoire

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Ci is a mage namespace which will manage continuous integration actions
type Ci mg.Namespace

// Validate validate circleci configuration file (circleci/config.yml).
func (Ci) Validate() error {
	return sh.RunV("circleci-cli", "config", "validate")
}

// Build execute circleci job build on local.
func (ci Ci) Build() error {
	return ci.localExecute("build")
}

func (ci Ci) localExecute(job string) error {
	return sh.RunV("circleci-cli", "local", "execute", "--job", job)
}
