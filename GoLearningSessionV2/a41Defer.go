package main

import (
    "fmt"
)

// "defer" statements are recommended to be used for:
// a) Closing connections when opening files/connecting to databases
// b) Unlocking Mutexes (which we will learn about later on in the course)
// c) Other cleanup code
//
// If you use them for any other purpose, they will make your code HIGHLY unreadable.

// Order of statement execution in Go:
// All undeferred statements
// All deferred statements, in Last In, First Out order -- ONLY when the Function that owns them is exiting.
//
// In the example below, we have 2 functions (f1 and f2), each with their own defer statements, and f1 calls f2.
// The defer statements in f2 will execute (in Last In, First Out order) before
// The defer statements in f1 will execute (also in Last In, First Out order)

 func f1(){
     fmt.Println("f1 start")

     defer fmt.Println("f1 exit")
     defer fmt.Println("f1 cleaning up...")

     f2()
 }

 func f2() {
     fmt.Println("f2 start")

     defer fmt.Println("f2 exit")
     defer fmt.Println("f2 cleaning up...")
 }

func main() {
    fmt.Println("Open the file.")
    defer fmt.Println("Close the file.")
    fmt.Println("Do something with the file...")
    // OUTPUT:
    // Open the file.
    // Do something with the file...
    // Close the file.

    f1()
    // OUTPUT:
    // f1 start
    // f2 start
    // f2 cleaning up...
    // f2 exit
    // f1 cleaning up...
    // f1 exit
}
