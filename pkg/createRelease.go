package pkg

import (
	"bufio"
	"fmt"
	"os/exec"
)

func CreateRelease (choice bool, version string) {
	fmt.Printf("Creating release for %s.", version)
	fmt.Println()

	cmd := exec.Command("goreleaser", "release", "--clean")

	stderr,	err := cmd.StderrPipe()
	stdout, outErr := cmd.StdoutPipe()

	if err != nil {
		panic(err)
	}

	if outErr != nil {
		panic(outErr)
	}

	cmd.Start()
	
	scannerErr := bufio.NewScanner(stderr)
	scannerOut := bufio.NewScanner(stdout)

	for scannerErr.Scan() {
		message := scannerErr.Text()
		fmt.Println(message)
	}

	for scannerOut.Scan() {
		message := scannerOut.Text()
		fmt.Println(message)
	}

	cmd.Wait()
}