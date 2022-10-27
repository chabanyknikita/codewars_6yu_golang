package main

import (
	"sort"
	"strings"
)

func InArray(array1 []string, array2 []string) []string {
	var res_arr []string
	for i := 0; i < len(array2); i++ {
		for j := 0; j < len(array1); j++ {
			if strings.Contains(array2[i], array1[j]) {
				res_arr = append(res_arr, array1[j])
			}
		}
	}
	sort.Strings(res_arr)
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range res_arr {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
