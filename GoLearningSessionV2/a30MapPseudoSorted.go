package main

import (
    "fmt"
	"sort"
)

func main() {
    // If we need to get values from a map in order,
    // we need to sort either by keys, or by values.
    // Here, we use a sorted slice of Keys to iterate over a map

    m := make(map[string]int)

    m["a"] = 1
    m["c"] = 3
    m["d"] = 4
    m["b"] = 2

    keys := make([]string, len(m))

    // first, make a slice containing all the keys
	i := 0
    for k := range m {
        keys[i] = k
        i++
    }

	// sort the keys in Ascending order
    sort.Strings(keys)

    // Iterate over our sorted "keys" slice, which we then use to
    // access our our map in Ascending order
    for i := range keys{
        fmt.Println(m[keys[i]])
    }
    // OUTPUT:
    // 1
    // 2
    // 3
	// 4
}

