package pkg

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Version (branch string, tag string) string {
	var identifier string
	var new_version string
	var version_placement int
	var version string

	if tag == "" {
		version = Latest()
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
	case "bugfix":
		identifier = "patch"
		version_placement = 2
	case "hotfix":
		identifier = "patch"
		version_placement = 2
	default:
		fmt.Println("\nYou didn't include a proper branch name!")
		fmt.Println("")
		flag.Usage()
		os.Exit(2)
	}

	ver_array := strings.Split(strings.Split(version, "v")[1], ".")
	new_ver, err := strconv.Atoi(ver_array[version_placement])
	
	if err != nil {
		panic(err)
	}

	ver_array[version_placement] = strconv.Itoa(new_ver + 1)
	new_version = "v" + strings.Join(ver_array, ".")

	fmt.Printf("\nIncrementing the %s %s version.\n\n%s ➜ %s\n\n", version, identifier, version, new_version)

	Increment(new_version)	

	return "success"
}