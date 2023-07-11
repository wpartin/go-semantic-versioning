package main

import (
	"flag"
	"fmt"
	"go-semantic-versioning/pkg"
	"os"
)

func main () {
	branch := flag.String("branch", "", "[REQUIRED] The branch to initate versioning from; release, feature, bugfix, hotfix.")
	release := flag.Bool("release", false, "[OPTIONAL] Create a GitHub Release with the new tag. Must have gh cli installed for this to work.")
	tag := flag.String("tag", "", "[OPTIONAL] The tag for a previous release that you want to increment. Must use the following pattern: \"v#.#.#\"")


	flag.Parse()

	if !pkg.IsFlagPresent("branch") {
		flag.Usage()
		os.Exit(2)
	} else {
		fmt.Println("\nStarting go-semantic-versioning.")

		version := pkg.Version(*branch, *tag)

		if *release {
			pkg.CreateRelease(*release, version)
		}
	}
}