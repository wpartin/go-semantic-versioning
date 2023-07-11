package pkg

import (
	"fmt"
	"os/exec"
)

func PrintCurrentVersion () {
	output := exec.Command("brew", "list", "--versions", "go-semantic-versioning").Run()

	fmt.Println(output)
}