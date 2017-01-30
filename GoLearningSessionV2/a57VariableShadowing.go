package main

import (
    "fmt"
)

func main() {
    // Variable Shadowing occurs when an already-existing identifier
    // is reused on a new variable inside a different code block (e.g. If's, Switches, Goroutines)
    // or on Multiple Implicit Declarations.

    num := 1

    fmt.Println(num)

    // Shadow Variables *inside* code blocks do not affect the outside variable.
    // The only problem is when you accidentally use them when you meant to use the outside variable.
    //
    if num := 2; num % 2 == 0 {
        fmt.Println("even")
    } else {
        fmt.Println("odd")
    }

    fmt.Println(num)

    // OUTPUT:
    // 1
    // even
    // 1

    // ---------------------------------------------
    //
    // Shadow Variables inside multiple implicit declarations are more dangerous,
    // because we LOSE the original value of the variable.

    x := 10

    fmt.Println(x)
    // OUTPUT:
    // 10

    x, y := 20, 30

    fmt.Println(x, y)
    // OUTPUT:
    // 20 30
}