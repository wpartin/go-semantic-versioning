package pkg

import (
	"bufio"
	"fmt"
	"os/exec"
)

func CreateRelease (version string) {
	fmt.Printf("Creating release for %s.\n\n", version)

	cmd := exec.Command("goreleaser", "release", "--clean")

	stderr,	err := cmd.StderrPipe(); if err != nil { panic(err) }
	stdout, outErr := cmd.StdoutPipe(); if outErr != nil { panic(outErr) }

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