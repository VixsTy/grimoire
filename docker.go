package grimoire

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Docker is a mage namespace which will manage docker actions
type Docker mg.Namespace

// Build build a docker image.
func (Docker) Build() error {
	color.Red("# Docker -------------------------------------------------------------------")
	fmt.Printf("BUILD_DATE : %s\n", time.Now().Format(time.RFC3339))
	fmt.Printf("VERSION : %s\n", Tag())
	fmt.Printf("VCS_REF : %s\n", Hash())

	fmt.Printf(" > Production image\n")
	return sh.RunV("docker", "build",
		"-f", "Dockerfile",
		"--build-arg", fmt.Sprintf("BUILD_DATE=%s", time.Now().Format(time.RFC3339)),
		"--build-arg", fmt.Sprintf("VERSION=%s", Tag()),
		"--build-arg", fmt.Sprintf("VCS_REF=%s", Hash()),
		"-t", MainDirectoryName()+":latest",
		".")
}
