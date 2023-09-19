package main

import (
	"flag"
	"fmt"
	"go-semantic-versioning/pkg"
	"os"
)

func main () {
	branch  := flag.String("branch", "", "[REQUIRED] The branch to initate versioning from; release, feature, bugfix, hotfix.")
	release := flag.Bool("release", false, "[OPTIONAL] Create a GitHub Release with goreleaser. Must have goreleaser installed for this to work.")
	tag     := flag.String("tag", "", "[OPTIONAL] The tag for a previous release that you want to increment. Must use the following pattern: \"v#.#.#\"")
	version := flag.Bool("version", false, "[OPTIONAL] Check the current version.")

	flag.Parse()

	if *version {
		pkg.PrintCurrentVersion()
	} else if !pkg.IsFlagPresent("branch") {
		flag.Usage()
		os.Exit(2)
	} else {
		fmt.Println("\nStarting go-semantic-versioning.")

		version := pkg.Version(*branch, *tag)

		if *release || *branch == "release" { pkg.CreateRelease(version) }
	}
}