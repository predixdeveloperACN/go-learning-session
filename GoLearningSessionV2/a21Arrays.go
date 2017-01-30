package main

import (
    "fmt"
)

func main() {
    // NOTE: Arrays CANNOT be sorted or resized at runtime
    //       All items in an array MUST be of the same type

    var colors [3]string

    // Assign some values to your Array
    colors[0] = "Red"
    colors[1] = "Green"
    colors[2] = "Blue"

    // You can also assign values to your Arrays upon declaration
    var numbers = [5]int {5,3,1,2,4}

    fmt.Println(numbers)
    // OUTPUT:
    // [5 3 1 2 4]

    fmt.Println(colors)
    // OUTPUT:
    // [Red Green Blue]

    fmt.Println(colors[1])
    // OUTPUT:
    // Green

    fmt.Println("Size of colors:", len(colors))
    // OUTPUT:
    // Size of colors: 3
    //
    // NOTE: Even if we just put in 2 colors in the Array, the Len will still be 3
    //       This is because Len gives us the SIZE of the array, *NOT* the number of elements
}
