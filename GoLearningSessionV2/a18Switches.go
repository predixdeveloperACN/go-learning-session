package main

import (
    "fmt"
    "math/rand"
)

func main() {    
    dayOfWeek := 1
    var result string

    switch dayOfWeek {
    case 1:
        result = "It's Sunday!"
    case 7:
        result = "It's Saturday!"
    default:
        result = "It's a weekday!"
    }
    fmt.Println(result)
    // OUTPUT:
    // It's Sunday!

    // Same with If Statements, you can add *ONE* statement before you actually begin your Switch
    //
    switch dayOfWeek := rand.Intn(6) + 1; dayOfWeek {
    case 1:
        result = "It's Sunday!"
    case 7:
        result = "It's Saturday!"
    default:
        result = "It's a weekday!"
    }
    fmt.Println(result)
    // OUTPUT:
    // It's a weekday!

    // you can also do multiple cases in one line, like so:
    switch dayOfWeek := 4; dayOfWeek {
    case 1:
        result = "It's Sunday!"
    case 7:
        result = "It's Saturday!"
    case 2, 3, 4, 5, 6:
        result = "It's a weekday!"
    }
    fmt.Println(result)
    // OUTPUT:
    // It's a weekday!

    // You can also use Condition Statements instead of literal Values.
    // This way, you can cleanly rewrite long If-Then-Else statements.
    x := 22
    switch {
    case x < 0:
        result = "Negative Number"
    case x >= 1 && x <= 10:
        result = "Within 1 to 10"
    case x >= 11 && x <= 20:
        result = "Within 11 to 20"
    case x >= 21 && x <= 30:
        result = "Within 21 to 30"
    }
    fmt.Println(result)
    // OUTPUT:
    // Within 21 to 30

    x = -42
    switch {
    case x < 0:
        fmt.Println("Less than zero")
        // fall-through up to the next case
        fallthrough
    case x == 0:
        fmt.Println("Equal to zero")
    case x > 0:
        result = "Greater than zero"
    }
    // OUTPUT:
    // Less than zero
    // Equal to zero
}
