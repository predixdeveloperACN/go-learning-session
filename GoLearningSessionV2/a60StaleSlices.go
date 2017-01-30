package main

import (
    "fmt"
)

// Stale Slices are when arrays that "point" to
// each other unexpectedly lose their references.

func main() {
    s1 := []int{1, 2, 3}

    fmt.Println(s1)
    // OUTPUT:
    // [1 2 3]

    s2 := s1[1:]

    fmt.Println(s2)
    // OUTPUT:
    // [2 3]

    // Adding 20 to all items in s2 also affects s1 (because slices are by reference)
    for i := range s2 { s2[i] += 20 }

    fmt.Println(s1)
    fmt.Println(s2)
    // OUTPUT:
    // [1 22 23]
    // [22 23]

    // When you append to a child slice, the reference to the
    // parent slice will ONLY be preserved if the parent slice's Capacity
    // can ALSO handle the append.
    //
    // NOTE: Appending to child slices will NEVER affect the parent slice.
    //       Only element updates (to the same elements) will affect the parent slice.
    //
    // Here, because s1's capacity is 3 (because we initialized it directly,
    // it copies the InitialSize by default), it can no longer handle the
    // Append to s2, which cuts our reference to it.
    //
    s2 = append(s2, 4)

    // Add 10 to all items in s2
    // This does not affect s1 because s2 is no longer pointing to it
    for i := range s2 { s2[i] += 10 }

    // s1 is now "stale"
    fmt.Println(s1)
    fmt.Println(s2)
    // OUTPUT:
    // [1 22 23]
    // [32 33 14]

    // ---------------------------------------------
    //
    // If we used "make" for our parent slice and set our capacity to 5,
    // it can handle 2 Appends from any child slice before going "stale"
    //
    //     s1 := make([]int, 3, 5)
    //     s1[0] = 1
    //     s1[1] = 2
    //     s1[2] = 3
}
