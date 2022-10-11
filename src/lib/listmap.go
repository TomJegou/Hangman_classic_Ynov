package lib

import (
	"sort"
)

/*
Function that list a map in a key ascending order and also return all the keys of the map
*/

func Listmap(m map[string]string) []string {
	keys := []string{}
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for index_key, key := range keys {
		PrintColor("["+key+"]"+": "+m[keys[index_key]]+"\n", "White")
	}
	return keys
}
