package main

import (
    "fmt"
    "sort"
)

func main() {
    // The sort package sorts in Ascending order by default.
    // It affects the slice passed to it.

    words := []string{"crab", "apple", "dog", "banana"}

    sort.Strings(words)

    fmt.Println(words)
    // OUTPUT:
    // [apple banana crab dog]

    numbers := []int{4, 2, 7, 9, 3}

    sort.Ints(numbers)

    fmt.Println(numbers)
    // OUTPUT:
    // [2 3 4 7 9]
    //
    // ---------------------------------------------

    nums := []int{5, 2, 6, 3, 1, 4}

    // Sorting in Descending order is a little more complicated.
    //
    // First, we sort the slice in Ascending order (sort.IntSlice)
    // Next, we Reverse the order (sort.Reverse)
    // Finally, we use "sort.Sort" to reflect our changes to the slice
    //
    sort.Sort(sort.Reverse(sort.IntSlice(nums)))

    fmt.Println(nums)
    // OUTPUT:
    // [6 5 4 3 2 1]

    // Returns whether a slice is sorted in Ascending order or not.
    // Sadly, there's no built-in function to check for Descending order.
    fmt.Println(sort.IntsAreSorted(numbers))
    // OUTPUT:
    // true
}
