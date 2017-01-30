package main

import (
    "fmt"
)

func main() {
    sum := 0

    // For loop
    for i := 0; i < 10; i++ {
        sum += i
    }

    // While loop
    sum = 1
    for sum < 1000 {
        sum += sum    
    }

    for i := 0; i < 10; i++ {
        if i == 3 {
            // skip an iteration
            continue
        } else if i == 5 {
            //stop the entire loop
            break
        }
        fmt.Println(i)
    }
    // OUTPUT:
    // 0
    // 1
    // 2
    // 4

    // Infinite loop
    for {
        fmt.Println("We have entered an endless recursion of time.")
    }
}
