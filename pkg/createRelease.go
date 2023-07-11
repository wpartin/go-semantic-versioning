package pkg

import (
	"fmt"
	"os/exec"
)

func CreateRelease (choice bool, version string) {
	fmt.Printf("\nCreating release for %s.", version)

	exec.Command("goreleaser", "release", "--clean").Run()
}