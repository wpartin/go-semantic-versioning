package pkg

import (
	"strconv"
	"strings"
)

func DefineNewVersion (branch string, version string, version_placement int) string {
	ver_array := strings.Split(strings.Split(version, "v")[1], ".")
	new_ver, err := strconv.Atoi(ver_array[version_placement])
	var new_version string
	
	if err != nil {
		panic(err)
	}

	switch branch {
	case "release":
		ver_array[version_placement] = strconv.Itoa(new_ver + 1)
		ver_array[version_placement + 1] = strconv.Itoa(0)
		ver_array[version_placement + 2] = strconv.Itoa(0)
		new_version = "v" + strings.Join(ver_array, ".")
	case "feature":
		ver_array[version_placement] = strconv.Itoa(new_ver + 1)
		ver_array[version_placement + 1] = strconv.Itoa(0)
		new_version = "v" + strings.Join(ver_array, ".")
	case "bugfix", "hotfix":
		ver_array[version_placement] = strconv.Itoa(new_ver + 1)
		new_version = "v" + strings.Join(ver_array, ".")
	}

	return new_version
}