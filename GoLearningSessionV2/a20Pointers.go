package main

import (
    "fmt"
)

func main() {
    // This is an Integer Pointer that can "point" to any Int variable
    // All Pointers are "nil" when not set
    var p *int
    
    var int1 int = 42
    
    // Point our pointer to a variable with the "&" sign
    //
    // NOTE: We're actually pointing our Pointer to the variable's Pointer Address.
    //
    p = &int1
    
    if p != nil {
        // to use Pointers, use an asterisk (*)
        fmt.Println("Value of p:", *p)
    } else {
        fmt.Println("p is not set!")
    }
    // OUTPUT:
    // Value of p: 42

    // ---------------------------------------------

    var float1 float64 = 42.13

    // implicit pointer
    p2 := &float1

    // make changes to the variable via Pointer
    *p2 = *p2 / 31

    fmt.Printf("value of p2: %.2f \n", *p2)
    fmt.Printf("value of float1: %.2f", float1)
    // OUTPUT:
    // value of p2: 1.36
    // value of float1: 1.36
}
