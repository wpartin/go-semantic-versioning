package pkg

import (
	"fmt"
	"os/exec"
)

func CreateRelease (choice bool, version string) {
	fmt.Printf("Creating release for %s.", version)
	fmt.Println()

	exec.Command("goreleaser", "release", "--clean").Run()
}