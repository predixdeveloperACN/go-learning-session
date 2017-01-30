package main

import (
    "fmt"
)

// Slices are always Passed by Reference, so there is no need for Pointers.
// Slices also don't have very good support for element removal.
// Elements after the removed one must be shifted forward, which is slow,
// especially for very large slices.

func main() {
    numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

    fmt.Println(numbers)
    // OUTPUT:
    // [0 1 2 3 4 5 6 7 8 9]

    // remove 5
    n := RemoveAtIndexOrdered(numbers, 5)

    fmt.Println(n)
    // OUTPUT:
    // [0 1 2 3 4 6 7 8 9]

    // remove 4
    n = RemoveAtIndexUnordered(numbers, 4)

    fmt.Println(n)
    // OUTPUT:
    // [0 1 2 3 9 6 7 8]
}

// This function removes an element from a Slice at a specified Index
// while PRESERVING the order of the slice.
//
// It does so by cutting the slice into 2 (excluding the element at the index)
// and then combining those 2 slices back into 1 slice.
//
// NOTE: Ellipses (...) are used to append one slice to another
//
func RemoveAtIndexOrdered(s []int, index int) []int {
    return append(s[:index], s[index+1:]...)
}

// This function also removes an element from a Slice at a specified Index,
// but the order is *NOT* preserved, which makes it faster.
//
// Here, we copy the value of the last item to the Index provided,
// and then truncate the Slice
//
func RemoveAtIndexUnordered(s []int, i int) []int {
    s[len(s)-1], s[i] = s[i], s[len(s)-1]
    return s[:len(s)-2]
}
