package pkg

import (
	"fmt"
	"os/exec"
)

func PrintCurrentVersion () {
	output, _ := exec.Command("brew", "list", "--versions", "go-semantic-versioning").Output()

	fmt.Println()
	fmt.Println(string(output))
}