package main

import (
    "time"
    "fmt"
)

func main() {
    // Go’s "select" lets you wait on multiple channel operations.
    // Combining Goroutines and Channels with Select is a powerful feature of Go.

    c1 := make(chan string)
    c2 := make(chan string)

    // Each channel is on its own thread and will receive a value after
    // some amount of time to simulate an expensive task.

    go func() {
        time.Sleep(time.Second * 1)
        c1 <- "one"
    }()

    go func() {
        time.Sleep(time.Second * 2)
        c2 <- "two"
    }()

    // We'll use "select" to await both of these values simultaneously,
    // printing each one as it arrives.
    //
    // Remember that Select cases will lock until it can run something,
    // so this 2-iteration loop is sure to iterate 2 times successfully
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
    }

    // OUTPUT:
    // received one
    // received two
}
