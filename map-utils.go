package simpleutils

import (
	"reflect"
	"sort"
)

func SortedKeys(m interface{}) []string {
	keys := reflect.ValueOf(m).MapKeys()

	sorted := make([]string, 0, len(keys))
	for _, k := range keys {
		sorted = append(sorted, k.String())
	}

	sort.Strings(sorted)

	return sorted
}

