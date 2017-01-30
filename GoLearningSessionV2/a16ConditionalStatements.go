package main

import (
    "fmt"
)

func main() {
    // NOTE: Opening Braces need to be on the FIRST line of the block, NOT on the next.
    
    a := -1
    var result string
    
    if a < 0 {
        result = "Less than zero"
    } else if a == 0 {
        result = "Equal to zero"
    } else {
        result = "Greater than zero"
    }

    fmt.Println(result)
    // OUTPUT:
    // Less than zero

    // You can add *ONE* statement just before your If Condition, separated by a semicolon (;)
    // This statement is usually used to initialize a variable you plan to evaluate WITHIN the If statement.
    //
    // Note how we are able to reuse "a", even though it's already been declared above.
    // This is called Variable Shadowing, and while it is a useful feature of Go,
    // it can also pose potential problems, which we will learn later on in the course.
    //
    if a := 42; a < 0 {
        result = "Less than zero"
    } else if a == 0 {
        result = "Equal to zero"
    } else {
        result = "Greater than zero"
    }

    fmt.Println(result)
    // OUTPUT:
    // Greater than zero

    var err error
    
    if err != nil {
        panic(err)
    } else {
        fmt.Println("No errors!")
    }
    // OUTPUT:
    // No errors!
}
