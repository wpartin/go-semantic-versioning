package pkg

import "strings"

func RemoveLastItem(item []byte) []string {
	array_list := strings.Split(string(item), "\n")
	
	array_list = array_list[:len(array_list) - 1]

	return array_list
}