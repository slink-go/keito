package util

import (
	"fmt"
	"slices"
)

func PrintMapSorted(input map[string]interface{}, prefix string) {

	var keys []string = make([]string, 0)
	for k, _ := range input {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	for _, k := range keys {
		fmt.Printf("%s%s = %v\n", prefix, k, input[k])
	}

}
