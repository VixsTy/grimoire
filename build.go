package grimoire

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Build is a mage namespace which will manage building actions
type Build mg.Namespace

// Binary will trigger the compilation
func (Build) Binary(packageName, out string) error {
	fmt.Printf(" > Building %s [%s]\n", out, packageName)
	repoName := MainModule()

	varsSetByLinker := map[string]string{
		repoName + "/internal/version.Version":   Tag(),
		repoName + "/internal/version.Revision":  Hash(),
		repoName + "/internal/version.Branch":    Branch(),
		repoName + "/internal/version.BuildUser": os.Getenv("USER"),
		repoName + "/internal/version.BuildDate": time.Now().Format(time.RFC3339),
		repoName + "/internal/version.GoVersion": runtime.Version(),
	}
	var linkerArgs []string
	for name, value := range varsSetByLinker {
		linkerArgs = append(linkerArgs, "-X", fmt.Sprintf("%s=%s", name, value))
	}
	linkerArgs = append(linkerArgs, "-s", "-w")

	return sh.RunWith(map[string]string{
		"CGO_ENABLED": "0",
	}, "go", "build", "-ldflags", strings.Join(linkerArgs, " "), "-mod=vendor", "-o", fmt.Sprintf("bin/%s", out), packageName)
}
