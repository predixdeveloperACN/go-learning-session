package main

import (
    "fmt"
)

// Channels are a nice way to synchronize threads in Go.
//
// NOTE: You MUST ONLY use Channels for Goroutines (multi-threading) because
//       they will block your main thread when used on their own.
//
// Channels are a typed conduit through which you can send and receive values with the channel operator, "<-"
// The data flows in the direction of the arrow:
//
//     ch <-var    // Send value of "var" to channel "ch"
//
//     var = <-ch  // Receive from "ch", and assign value to "var"
//
// Like Maps and Slices, Channels must be created before use:
//
//     ch := make(chan int)
//
// By default, a Channel can only Send or Receive values ONE AT A TIME.
// Sending or Receiving more than once will block that Channel's thread.
// This behavior allows goroutines to synchronize without explicit locks or condition variables.

// The code below sums the numbers in a slice, distributing the work between two goroutines.
// Once both goroutines have completed their computation, it calculates the final result.
//
func main() {
    s := []int{7, 2, 8, -9, 4, 0}

    // make an Int Channel
    c := make(chan int)

    // Run the "sum" function on 2 separate threads
    // Note that we divide the work equally for both functions
    //
    go sum(s[:len(s)/2], c) //  7 + 2 + 8 = 17
    go sum(s[len(s)/2:], c) // -9 + 4 + 0 = -5

    // Receive from c
    // Note that this line will block until BOTH threads have finished.
    //
    // When you receive from a channel, you get results in the order of the Sends you performed in the code.
    // You can reuse the same channel to Receive multiple times, but each Receive MUST have a matching Send.
    //
    x, y := <-c, <-c

    fmt.Println(x, y, x+y)
    // OUTPUT:
    // -5 17 12
    //
    // Note that if we placed a Sleep Statement between our 2 go statements, the Output will be:
    // 17 -5 12
    //
    // The discrepancy is because of how unpredictable Goroutines can be --
    // sometimes, one thread will always be faster than another for no obvious reason.
    // We will learn more about how Channels handle Send and Receive orders later on in the course.
}

func sum(s []int, c chan int) {
    sum := 0

    for _, v := range s {
        sum += v
    }

    // send sum to c
    c <- sum
}