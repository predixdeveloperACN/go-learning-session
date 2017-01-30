package main

import (
    "fmt"
)

func main() {
    doSomething(21)
    // OUTPUT:
    // Twice 21 is 42

    doSomething("hello")
    // OUTPUT:
    // "hello" is 5 bytes long

    doSomething(true)
    // OUTPUT:
    // I don't know about type bool!
}

// This function gives different outputs depending on the Type of parameter passed to it.
// Anything can implement the "blank" interface in Go (interface{}).
// We will learn more about Interfaces later on in the course.
// For now, just think of it as a way to accept parameters of ANY data type.
//
func doSomething(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("Twice %v is %v\n", v, v*2)
    case string:
        fmt.Printf("%q is %v bytes long\n", v, len(v))
    default:
        fmt.Printf("I don't know about type %T!\n", v)
    }
}