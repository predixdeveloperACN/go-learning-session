package main

import (
    "fmt"
)

// A Sender can close a channel to indicate that no more values will be sent to it.
// Receivers can test whether a channel has been closed by assigning a second parameter to the Receive expression:
//
//     v, ok := <-ch
//
// Here, "ok" will be False if there are no more values to Receive and the Channel has been closed.

// When using "range" to receive values from a channel,
// the loop will block until the channel is explicitly closed.
// This is because channels can (in theory) hold infinite values by alternating Sends & Receives.
//
// Note 1: Only the sender should close a channel, NEVER the receiver.
//         Sending on a closed channel will cause a panic.
//
// Note 2: Channels aren't like files; you don't usually need to close them.
//         Closing is only necessary when the receiver must be explicitly told
//         that there are no more values to Receive, such as to terminate a range loop.

func main() {
    c := make(chan int, 10)
    
    // Here, "cap" gives us the capacity of channel c, which is 10 (the buffer).
    // It's always good practice to check "cap" to avoid exceeding a channel's buffer
    go fibonacci(cap(c), c)
    
    // "range" will iterate for every Send we do on the channel (in chronological order).
    // If we din't close the channel, the range loop would block when there are no more values available
    // (after the 10th value in this case), so ALWAYS close your channels when you plan to iterate over them with range.
    for i := range c {
        fmt.Println(i)
    }

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
}

func fibonacci(limit int, ch chan int) {
    x, y := 0, 1
    for i := 0; i < limit; i++ {
        // here, channel ch will hold the members of the fibonacci sequence
        ch <- x
        x, y = y, x+y
    }
    close(ch)
}