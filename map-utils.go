package simpleutils

import (
	"reflect"
	"sort"
)

//SortedKeys returns the key array for the map, sorted A-Z
func SortedKeys(m interface{}) []string {
	keys := reflect.ValueOf(m).MapKeys()

	sorted := make([]string, 0, len(keys))
	for _, k := range keys {
		sorted = append(sorted, k.String())
	}

	sort.Strings(sorted)

	return sorted
}
