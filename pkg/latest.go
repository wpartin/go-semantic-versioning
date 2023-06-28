package pkg

import (
	"log"
	"os/exec"
)

func Latest() string {
	versions, err := exec.Command("git", "tag", "-l").Output()
	
	if err != nil {
		log.Fatal(err)
	}
	
	version_list := RemoveLastItem(versions)

	latest := version_list[len(version_list) - 1]
	
	return latest
}