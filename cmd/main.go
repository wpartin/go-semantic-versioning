package main

import (
	"flag"
	"fmt"
	"go-semantic-versioning/pkg"
	"os"
)

func main () {
	branch := flag.String("branch", "", "[REQUIRED] The branch to initate versioning from; release, feature, bugfix, hotfix.")
	tag := flag.String("tag", "", "[OPTIONAL] The tag for a previous release that you want to increment. Must use the following pattern: \"v#.#.#\"")

	flag.Parse()

	if !pkg.IsFlagPresent("branch") {
		flag.Usage()
		os.Exit(2)
	} else {
		fmt.Println("\nStarting go-semantic-versioning.")

		pkg.Version(*branch, *tag)
	}
}