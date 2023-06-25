package pkg

import (
	"fmt"
	"strconv"
	"strings"
)

func Version (branch string) string {
	var identifier string
	var message string
	var version_placement int
	var version string = "1.0.0"
	var version_new string
	
	switch branch {
	case "release":
		identifier = "major"
		version_placement = 0
		
		vStr := strings.Split(version, ".")[version_placement]
		v, err := strconv.ParseInt(vStr, 10, 0)

		if err != nil {
			fmt.Println("Error during type conversion.")
			return "failure"
		}
		fmt.Println(v)
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
		message = "You didn't include a proper branch name!"
	}

	message = fmt.Sprintf("Incrementing %s version. v%s becomes v%s", identifier, version, version_new)

	if message != "You didn't include a proper branch name!" {
		fmt.Println(message, identifier, version_placement)
		return "success"
	} else {
		fmt.Println(message)
		return "error"
	}
}