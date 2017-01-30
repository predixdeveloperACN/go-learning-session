package main

import (
    "fmt"
)

func main() {
    colors := []string {"red", "blue", "yellow"}

    // Slice iteration via for loop
    // (NOT recommended)
    for i := 0; i < len(colors); i++ {
        fmt.Println(colors[i])
    }
    // OUTPUT:
    // red
    // blue
    // yellow

    // "range" returns 2 values:
    // The index, and the Value at that index
    //
    // Slice iteration with index only:
    //
    for i := range colors {
        fmt.Println(colors[i])
    }
    // OUTPUT:
    // red
    // blue
    // yellow

    // slice iteration with index and value:
    //
    for i, v := range colors {
        fmt.Printf("%v = %v\n", i, v)
    }
    // OUTPUT:
    //0 = red
    //1 = blue
    //2 = yellow
}
