package maps

import "fmt"

// IterateMap : builds a nested map and iterates and prints the entries of the map
func IterateMap() {
	// initialize an empty map syntax: elements := make(map[string]map[string]string)
	elements := map[string]map[string]string{
		"H": {
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": {
			"name":  "Helium",
			"state": "gas",
		},
		"Li": {
			"name":  "Lithium",
			"state": "solid",
		},
		"Fe": {
			"name":  "Iron",
			"state": "solid",
		},
	}

	// attempts to lookup the key value in the map, if successful runs code in block
	if element, ok := elements["Fe"]; ok {
		fmt.Println(element["name"], element["state"])
	}
}
