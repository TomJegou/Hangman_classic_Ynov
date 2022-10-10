package lib

import (
	"sort"
)

func Listmap(m map[string]string) {
	keys := []string{}
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for index_key, key := range keys {
		PrintColor("["+key+"]"+": "+m[keys[index_key]]+"\n", "White")
	}
}
