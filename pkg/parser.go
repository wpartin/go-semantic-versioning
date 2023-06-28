package pkg

import (
	"fmt"
	"strconv"
	"strings"
)

func Parser (version string, version_placement int) string {
	version = strings.Split(version, ".")[version_placement]
	v, err := strconv.ParseInt(version, 10, 0)

	if err != nil {
		fmt.Println("\nError during type conversion.")
	}

	return strconv.Itoa((int(v)))
}