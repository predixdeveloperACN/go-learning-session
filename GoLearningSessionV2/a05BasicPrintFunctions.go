package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello, world!")
    // OUTPUT:
    // Hello, world!

    str1 := "Hi,"
    str2 := "my name is"
    str3 := "Vince!"
    
    // String Concatenation 1
    // NOTE: Spaces need to be added manually.
    //
    fmt.Println(str1 + " " + str2 + " " + str3)
    // OUTPUT:
    // Hi, my name is Vince!

    // String Concatenation 2
    // NOTE: Spaces are added for you.
    //
    fmt.Println(str1, str2, str3)
    // OUTPUT:
    // Hi, my name is Vince!
}
