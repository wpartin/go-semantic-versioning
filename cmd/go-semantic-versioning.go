package main

import (
	"flag"
	"fmt"
	"go-semantic-versioning/pkg"
	"os"
)

func main () {
	branch := flag.String("branch", "", "[REQUIRED] The branch to initate versioning from; release, feature, bugfix, hotfix.")
	release := flag.String("gh-release", "", "[OPTIONAL] Create a GitHub Release with the new tag. Must have gh cli installed for this to work. Set true to use.")
	tag := flag.String("tag", "", "[OPTIONAL] The tag for a previous release that you want to increment. Must use the following pattern: \"v#.#.#\"")
	version := flag.String("version", "", "[OPTIONAL] Check the current version.")

	flag.Parse()

	if !pkg.IsFlagPresent("branch") {
		flag.Usage()
		os.Exit(2)
	} else if pkg.IsFlagPresent("version") {
		fmt.Printf("go-semantic-versioning %s %s", *version, pkg.GetLatestTag())
	} else {
		fmt.Println("\nStarting go-semantic-versioning.")

		version := pkg.Version(*branch, *tag)

		pkg.CreateRelease(*release, version)
	}
}