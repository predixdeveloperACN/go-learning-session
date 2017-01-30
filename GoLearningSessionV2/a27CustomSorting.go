package main

import (
    "fmt"
    "sort"
)

// Sometimes we need to sort slices in an order other than Ascending or Descending.
// For example, suppose we wanted to sort strings by their length instead of alphabetically.
// Here’s an example of custom sorts in Go:

// In order to sort by a custom function in Go,
// we need a define an alias for the type of slice we want to affect.
///
// Here, we've created a 'ByLength' type that will extend the String Slice type.
type ByLength []string

// We need to implement "sort.Interface", and we do this by
// implementing "Len", "Less", and "Swap" for our alias so
// we can use the sort package's generic 'Sort' function.
//
// 'Len' and 'Swap' will usually be the same for most Data types and
// 'Less' will hold the actual custom sorting logic.
//
// In our case, we want to sort in order of increasing string length,
// so we use 'len(s[i])' and 'len(s[j])' here.

func (s ByLength) Len() int {
    return len(s)
}

func (s ByLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}

func main() {
    fruits := []string{"peach", "banana", "kiwi"}

    // We use our custom sort by first casting our slice to
    // 'ByLength', and then use 'sort.Sort' on that typed slice.
    sort.Sort(ByLength(fruits))

    fmt.Println(fruits)
    // OUTPUT:
    // [kiwi peach banana]
}
