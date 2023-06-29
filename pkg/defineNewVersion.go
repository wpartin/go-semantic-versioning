package pkg

import (
	"strconv"
	"strings"
)

func DefineNewVersion (version string, version_placement int) string {
	ver_array := strings.Split(strings.Split(version, "v")[1], ".")
	new_ver, err := strconv.Atoi(ver_array[version_placement])
	
	if err != nil {
		panic(err)
	}

	ver_array[version_placement] = strconv.Itoa(new_ver + 1)
	new_version := "v" + strings.Join(ver_array, ".")

	return new_version
}