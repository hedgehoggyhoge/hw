package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(s string) []string {
	var result []string

	slice := strings.Fields(s)

	cache := make(map[string]int, len(slice))

	for _, word := range slice {
		c, ok := cache[word]
		if ok {
			cache[word] = c + 1
		} else {
			cache[word] = 1
		}
	}

	type keyValue struct {
		Key   string
		Value int
	}

	var sortedStruct []keyValue

	for key, value := range cache {
		sortedStruct = append(sortedStruct, keyValue{key, value})
	}

	sort.Slice(sortedStruct, func(i, j int) bool {
		return (sortedStruct[i].Value > sortedStruct[j].Value) || ((sortedStruct[i].Value == sortedStruct[j].Value) && (sortedStruct[i].Key < sortedStruct[j].Key))
	})

	if len(sortedStruct) > 10 {
		sortedStruct = sortedStruct[:10]
	}

	for _, keyValue := range sortedStruct {
		result = append(result, keyValue.Key)
	}
	return result
}
