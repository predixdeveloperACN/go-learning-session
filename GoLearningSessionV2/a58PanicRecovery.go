package main

import (
    "fmt"
)

// Panics are errors that we throw when we want to stop our program because we
// encountered an error we don't want to handle, or shouldn't be happening at all.

func main() {
    BreakDown()
    // OUTPUT:
    // Recovered from: panic in func BreakDown

    DivideByZero(10)
    // OUTPUT:
    // Recovered from Panic:
    // runtime error: integer divide by zero
}

func BreakDown() {
    // we can Recover from Panics
    // (the program won't end and will continue execution)
    // by calling the "recover" function.
    //
    // However, it only works on a per-function basis,
    // and ONLY when it is *deferred* inside said function!
    //
    // When used on its own, it only serves to cancel the Panic:
    //
    //     defer func() {
    //         recover()
    //     }()
    //
    // However, when used inside a Print function, it shows the Panic it Recovered from:
    //
    defer func() {
        fmt.Println("Recovered from:", recover())
    }()

    panic("panic in func BreakDown")
}

// This function demonstrates a different approach to recovery --
// by first assigning its value to an error variable, and checking if it's not nil.
//
func DivideByZero(num int) {
    defer func() {
        if err := recover(); err != nil {
            fmt.Println("Recovered from Panic:")
            fmt.Println(err)
        }
    }()

    // Note how we make a separate variable for 0
    // This is because the Go Compiler is smart enough to throw an error
    // whenever it encounters divide by zero (/ 0) in the code.
    zero := 0
    res := num / zero

    fmt.Println(res)
}