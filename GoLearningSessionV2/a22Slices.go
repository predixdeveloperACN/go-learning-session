package main

import (
    "fmt"
)

func main() {
    // Slices are just expandable, resizable Arrays.
    // They're best used for ordered collections of values.
    // All items also need to be of the same type.
    // They can be resized and sorted at runtime.

    var colors = []string {"Red", "Green", "Blue"}
    fmt.Println(colors)
    // OUTPUT:
    // [Red Green Blue]

    // The slice automatically grows as needed when you add items
    colors = append(colors, "Purple")
    fmt.Println(colors)
    // OUTPUT:
    // [Red Green Blue Purple]

    // you can add more than one item at a time
    colors = append(colors, "Yellow", "Magenta", "Fuchsia")
    fmt.Println(colors)
    // OUTPUT:
    // [Red Green Blue Purple Yellow Magenta Fuchsia]
    // ---------------------------------------------

    // The Syntax for getting slice values is:
    // StartIndex : EndIndex (EXCLUDING value at said index)

    // remove the first item by specifying
    // a start and end index of what we want to keep
    colors = append(colors[1:len(colors)])
    //
    // Notice how Intellij IDEA dims the text for "len(colors)".
    // This is because the length of the slice is the default value,
    // which you can omit like so:
    //
    //     colors = append(colors[1:])

    fmt.Println(colors)
    // OUTPUT:
    // [Green Blue Purple Yellow Magenta Fuchsia]
    //

    // remove the last item by specifying
    // a start and end index of what we want to keep
    colors = append(colors[0:len(colors)-1])
    //
    // You can omit the start index if you want to start from the
    // beginning of the slice since 0 is the default value:
    //
    //     colors = append(colors[:len(colors)-1])

    fmt.Println(colors)
    // OUTPUT:
    // [Green Blue Purple Yellow Magenta]
    //
    // ---------------------------------------------

    // You can initialize a slice with the "make" function (InitialSize, Capacity).
    // The Capacity is set to the Size you specify by default, so you can omit it if you like.
    //
    // Here, we make an int slice with an initial size of 5, and capacity is 5
    // The size of 5 means it will have 5 elements by default,
    // each with the zero-value of the slice's type (for ints, that value is 0)
    //
    numbers := make([]int, 5, 5)

    fmt.Println(numbers)
    // OUTPUT:
    // [0 0 0 0 0]

    // Check the capacity of a slice
    fmt.Println(cap(numbers))
    // OUTPUT:
    // 5

    // Also, if you add items to the slice and exceed the Capacity,
    // the slice will double in size to accommodate future values
    numbers = append(numbers, 1, 2, 3)

    fmt.Println(numbers)
    // OUTPUT:
    // [0 0 0 0 0 1 2 3]

    // Check the capacity of a slice
    fmt.Println(cap(numbers))
    // OUTPUT:
    // 10
}
