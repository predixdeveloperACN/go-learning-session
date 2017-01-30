package main

import (
    "fmt"
    "strings"
)

func main() {

    // A slice of int slices.
    // The external slice has a size of 2 while
    // each internal slice has a size of size 3
    //
    // NOTE: When the closing brace is not on the same
    //       line as the last element of a Slice/Array/Map,
    //       you NEED to add a trailing comma to the last element.
    //
    board1 := [][]int{
        []int{0,0,0},
        []int{0,0,0},
    }

    // If the closing brace is on the same line as the
    // last element, no trailing comma is needed.
    //
    //     board1 := [][]int{
    //         []int{0,0,0},
    //         []int{0,0,0}}

    board1[0][0] = 1
    board1[0][1] = 2
    board1[0][2] = 3
    board1[1][0] = 4
    board1[1][1] = 5
    board1[1][2] = 6

    // Nested range iteration is the only way to properly
    // iterate over multi-dimensional slices.
    //
    // Remember that the second return value of range is
    // the value at that index.
    //
    // In this case, "i" will hold each internal slice.
    //
    for _, i := range board1 {
        for _, j := range i {
            fmt.Print(j)
        }
        fmt.Println()
    }
    // OUTPUT:
    // 123
    // 456

    // ---------------------------------------------
    //
    // A slice of string slices.
    // The external slice has a size of 3 while
    // its elements are slices of size 3
    //
    board2 := [][]string{
        []string{"_", "_", "_"},
        []string{"_", "_", "_"},
        []string{"_", "_", "_"},
    }
    // The coordinates are as follows:
    // 0,0 | 0,1 | 0,2
    // 1,0 | 1,1 | 1,2
    // 2,0 | 2,1 | 2,2

    board2[0][0] = "A"
    board2[0][2] = "B"
    board2[1][0] = "C"
    board2[1][1] = "D"
    board2[1][2] = "E"
    board2[2][2] = "F"

    // String Slice Iteration Version 1 (With Brackets)
    //
    // Here, we're just printing each string slice on its own line.
    // Remember that when a slice is Printed, spaces are automatically
    // added for each element. Also, by default, the brackets are included.
    //
    for i := 0; i < len(board2); i++ {
        fmt.Printf("%s\n", board2[i])
    }
    // OUTPUT:
    // [A _ B]
    // [C D E]
    // [_ _ F]

    // String Slice Iteration Version 2 (Without Brackets)
    //
    // This version uses strings.Join to append each element to
    // each other, so we need to add our spaces manually.
    // This version does not include brackets.
    //
    for i := 0; i < len(board2); i++ {
        fmt.Printf("%s\n", strings.Join(board2[i], " "))
    }
    // OUTPUT:
    // A _ B
    // C D E
    // _ _ F
}

    
