package pkg

import (
	"log"
	"os/exec"
)

func GetLatestTag() string {
	var result string
	versions, err := exec.Command("git", "tag", "-l").Output()
	
	if err != nil {
		log.Fatal(err)
	}

	if len(versions) != 0 {
		version_list := RemoveLastItem(versions)
		result = version_list[len(version_list) - 1]
	} else {
		result = "No existing tags"
	}
	
	return result
}