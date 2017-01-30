package main

import (
    "fmt"
    "strconv"
)

func main() {
    // Parsing means converting numbers to strings (and vice versa)

    // Here, we're parsing a string to base-10
    // (0 means base the bit size on the string's prefix),
    // with 64-bit precision.
    i, _ := strconv.ParseInt("123", 0, 64)

    fmt.Println(i)
    // OUTPUT:
    // 123

    f, _ := strconv.ParseFloat("1.234", 64)

    fmt.Println(f)
    // OUTPUT:
    // 1.234

    // Parse functions return errors when they fail.
    // Atoi means "convert string to int".
    //
    // NOTE: To convert an int to string, use ".Itoa"
    //
    n, e := strconv.Atoi("apples")
    fmt.Println(n, "\n", e)
    // OUTPUT:
    // 0
    // strconv.ParseInt: parsing "apples": invalid syntax
}
