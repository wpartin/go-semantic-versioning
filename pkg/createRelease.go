package pkg

import (
	"fmt"
	"os/exec"
)

func CreateRelease (choice string, version string) {
	switch choice {
	case "true":
		fmt.Printf("\nCreating release for %s.", version)

		exec.Command("gh", "release", "create", version).Run()
	}
}