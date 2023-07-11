package pkg

import (
	"flag"
	"fmt"
	"os"
)

func Version (branch string, tag string) string {
	var identifier string
	var new_version string
	var version_placement int
	var version string

	if tag == "" {
		version = GetLatestTag()
	} else {
		version = tag
	}
	
	switch branch {
	case "release":
		identifier = "major"
		version_placement = 0
	case "feature":
		identifier = "minor"
		version_placement = 1
	case "bugfix", "hotfix":
		identifier = "patch"
		version_placement = 2
	default:
		fmt.Println("\nYou didn't include a proper branch name!")
		fmt.Println("")
		flag.Usage()
		os.Exit(2)
	}

	if version != "No existing tags" && tag == "" {
		new_version = DefineNewVersion(branch, version, version_placement)

		fmt.Printf("\nIncrementing the %s %s version. (latest) \n\n%s ➜ %s\n\n", version, identifier, version, new_version)

		IncrementVersion(new_version)	
	} else if tag != "" {
		new_version = DefineNewVersion(branch, tag, version_placement)

		fmt.Printf("\nIncrementing the %s %s version. (tag) \n\n%s ➜ %s\n\n", tag, identifier, tag, new_version)

		IncrementVersion(new_version)
	} else {
		version = "v0.0.0"

		new_version = DefineNewVersion(branch, version, version_placement)

		fmt.Printf("\nNo existing tags found. Creating an initial version ➜ %s.\n\n", new_version)

		IncrementVersion(new_version)
	}

	return new_version
}