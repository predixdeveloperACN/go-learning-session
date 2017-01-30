package main

import (
    "fmt"
)

// The select statement lets a goroutine wait on multiple channel operations.
//
// A select blocks until one of its cases can run, then it executes that case.
// It chooses one at RANDOM if multiple are ready.

// The following code will run the fibonacci sequence 10 times before quitting:
//
func fibonacci(c, quit chan int) {
    x, y := 0, 1
    for {
        // select will run in the order that we populated our channels
        select {
        case c <- x:
            x, y = y, x+y
        case <-quit:
            fmt.Println("quit")
            return
        }
    }
}

func main() {
    // make 2 channels
    c := make(chan int)
    quit := make(chan int)
    
    // Receive from "c" 10 times, then populate "quit" 1 time.
    //
    // NOTE: Our program will lock until "c" Receives 10 times,
    //       so we need to run this on a separate thread with "go",
    //       once again inside an anonymous function.
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(<-c)
        }
        quit <- 0
    }()

    fibonacci(c, quit)

    // OUTPUT:
    // 0
    // 1
    // 1
    // 2
    // 3
    // 5
    // 8
    // 13
    // 21
    // 34
    // quit

    // Go remembers the order of how Channels were populated.
    //
    // Here, we populated the c channel 10 times, then the "quit" channel 1 time.
    // This resulted in the fibonacci sequence running 10 times before quitting.

    // If we had 3 channels -- c, b and quit -- plus an extra Select Case for b,
    // and we populated them in this order:
    //
    //     c - 5 times
    //     quit - 1 time
    //     b - 5 times
    //
    // The fibonacci would only run 5 times before quitting.

    // However, if we populated them in this order:
    //
    //    c - 5 times
    //    b - 5 times
    //    quit - 1 time
    //
    // Then the fibonacci would also run 10 times before quitting.
}
