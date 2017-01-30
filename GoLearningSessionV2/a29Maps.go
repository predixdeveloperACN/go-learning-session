package main

import (
    "fmt"
)

func main() {
    // Maps are Dictionaries/Hashtables
    //
    // Everything is *unordered* because we work with Keys instead of Indexes
    //
	// When you print maps, the order is random.

    // Make a map with String keys and Int values
    m := make(map[string]int)

    m["key1"] = 42
    m["key2"] = 86
    m["key3"] = 123

    fmt.Println(m)
    // OUTPUT:
    // map[key1:42 key2:86 key3:123]

    // Get value from map via key
    fmt.Println(m["key1"])
    // OUTPUT:
    // 42

    // Keys are case-sensitive.
    // If the given key isn't found,
    // we get the zero-value of the data type
    fmt.Println(m["Key1"])
    // OUTPUT:
    // 0

	// If you give a Key that's already present in the map,
	// the value will get OVERWRITTEN.
	m["key1"] = 404
	fmt.Println(m)
	// OUTPUT:
	// map[key1:404 key2:86 key3:123]

    // Delete item via key
    delete(m, "key1")

    // When iterating over maps with range,
    // it returns the Key and the Value.
    for k, v := range m {
        fmt.Printf("%v: %v \n", k, v)
    }
    // OUTPUT:
    // key2: 86
    // key3: 123

    // Just like with slices, you can use an "_"
    // in place of the key if you only want the value
    for _, val := range m {
        fmt.Printf("%v \n", val)
    }
    // OUTPUT:
    // 86
    // 123

    // ---------------------------------------------
    //
    // alternate way of initializing maps
    attendance := map[string] bool{
        "Bob": false,
        "Alice": true,
    }

    // When you get the value from a map via Key,
    // Go returns an optional, second value --
    // whether the key was present or not
    attended, ok := attendance["Bob"]

    if ok {
        fmt.Println("Bob attended?", attended)
    } else {
        fmt.Println("No info for Bob.")
    }
    // OUTPUT:
    // Bob attended? false

    attended, ok = attendance["Ted"]

    if ok {
        fmt.Println("Ted attended?", attended)
    } else {
        fmt.Println("No info for Ted.")
    }
    // OUTPUT:
    // No info for Ted.
}
