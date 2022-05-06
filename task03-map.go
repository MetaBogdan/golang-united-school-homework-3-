package homework

import (
	"fmt"
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {

	m := map[string]int{"value1": 1, "value2": 3, "value3": 2}
	var resultNew []string

	type key_value struct {
		Key   string
		Value int
	}

	var sorted_struct []key_value

	for key, value := range m {
		sorted_struct = append(sorted_struct, key_value{key, value})
	}

	sort.Slice(sorted_struct, func(i, j int) bool {
		return sorted_struct[i].Value < sorted_struct[j].Value
	})

	for _, key_value := range sorted_struct {
		resultNew = append(resultNew, key_value.Key)
		fmt.Printf("%s, %d\n", key_value.Key, key_value.Value)

	}

	return resultNew
}
