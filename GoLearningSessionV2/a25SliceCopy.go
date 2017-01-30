package main

import (
    "fmt"
)

func main() {
    // You can copy a slice to other slices with the "copy" function.
    // Your destination slice must NOT be empty so it can hold values.
    // The destination slice must, at the very least, be initialized via "make"
    
    // NOTE: Copy will copy as many elements as it can from one slice to another.
    //       If the destination slice is smaller than the source,
    //       only the elements that can fit will be copied.

    srcSlice := []int{1, 2, 3}
    destSlice := make([]int, len(srcSlice))

    copy(destSlice, srcSlice)

    fmt.Println(destSlice)
    fmt.Println(srcSlice)
    // OUTPUT:
    // [1 2 3]
    // [1 2 3]
    //
    // ---------------------------------------------

    // If you create a Slice from an Array,
    // it will affect that array (because Slices work by Reference)
    // as well as any other slices that point to that array.

    letters := [4]string{
        "A", "B", "C", "D",
    }

    fmt.Println(letters)
    // OUTPUT:
    // [A B C D]

    // The Syntax for getting slice values is:
    // StartIndex : EndIndex (EXCLUDING value at said index)

    a := letters[0:2] // A to C, excluding C
    b := letters[1:3] // B to D, excluding D

    fmt.Println(a, b)
    // OUTPUT:
    // [A B] [B C]

    // This line affects both "a" and "b" slices, as well as the "letters" array
    b[0] = "XXX"

    fmt.Println(a, b)
    fmt.Println(letters)
    // OUTPUT:
    // [A XXX] [XXX C]
    // [A XXX C D]

    // ---------------------------------------------

    // When returning portions of a slice,
    // make sure to use the Copy command to save Memory.

    data1 := getThreeBytes()
    fmt.Println(len(data1), cap(data1), &data1[0])
    // SAMPLE OUTPUT:
    // 3 10000 0xc042076000
    //
    // Note the wasted memory here. We only wanted 3 elements,
    // but we also got the entire array as well! (as shown with the 10000 Capacity)
    //
    // You'll also notice that we have the same Pointer address as the array from getThreeBytes.

    // To avoid such memory wastage, use "copy" to only get the slice portions you need.
    //
    data2 := getThreeBytesOptimized()
    fmt.Println(len(data2), cap(data2), &data2[0])
    // SAMPLE OUTPUT:
    // 3 3 0xc04203c2d8
}

func getThreeBytes() []byte {
    raw := make([]byte, 10000)

    // Note how we use a Pointer to see the address of the Array
    fmt.Println(len(raw), cap(raw), &raw[0])
    // SAMPLE OUTPUT:
    // 10000 10000 0xc042076000

    return raw[:3]
}

func getThreeBytesOptimized() []byte {
    raw := make([]byte, 10000)

    fmt.Println(len(raw), cap(raw), &raw[0])
    // SAMPLE OUTPUT:
    // 10000 10000 0xc042078900

    // Our destination slice will only have the Capacity we need
    res := make([]byte, 3)

    copy(res, raw[:3])

    return res
}