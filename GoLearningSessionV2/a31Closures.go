package main

import (
    "fmt"
)

// Go supports anonymous functions, which can form Closures.
//
// Anonymous functions are useful when you want to
// define a function inline without having to name it.
//
// Closures are useful when you want to have
// Multiple Instances of the same Function that can each
// "remember" or keep track of their own values.

// Below, the function "intSeq" returns another function, which we define anonymously inside it.
// Said anonymous function "closes over" intSeq's variable "i" to form a Closure.
// Any variable we assign to intSeq will remember its own value for "i".
//
func intSeq() func() int {
    i := 0

    return func() int {
        i += 1
        return i
    }
}

func main() {
    nextInt := intSeq() // intSeq is assigned to a variable.
                        // nextInt will now remember its own instance of "i"
    
    fmt.Println("nextInt", nextInt()) // intSeq's default "i" of 0 becomes 1
    fmt.Println("nextInt", nextInt()) // the value of "i" is "remembered" and becomes 2
    fmt.Println("nextInt", nextInt()) // 2 becomes 3
    // OUTPUT:
    // nextInt 1
    // nextInt 2
    // nextInt 3
    
    newInts := intSeq() // intSeq is assigned to a different variable.
	                    // newInts is a DIFFERENT instance and will remember its own version of "i"

    fmt.Println("newInts", newInts()) // intSeq's default "i" of 0 becomes 1
    // OUTPUT:
    // newInts 1

    // meanwhile, nextInt still remembers its own instance of "i"
    fmt.Println("nextInt", nextInt())
    // OUTPUT:
    // nextInt 4
}
