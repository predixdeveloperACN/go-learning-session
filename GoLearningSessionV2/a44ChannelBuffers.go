package main

import (
    "fmt"
)

// By Default, if you Send to a Channel, you must first Receive from it before being able to Send to it again.
// We Buffer channels when we want to Send and Receive from them multiple times in succession.
//
// NOTE: Sends to a buffered channel will block only when the buffer is full.
//       Receives will block when the buffer is empty.
//
// Deadlocks will occur when you attempt to Send or Receive more than the number of buffers you specified.

func main() {

    // Declare that the Send/Receive Buffer is 3.
    // This means our channel can only hold a max of 3 values at any given time.
    //
    // Sending to a channel will add 1 to that channel's held values.
    // Receiving from a channel will remove 1 from that channel's held values.
    //
    ch := make(chan int, 3)

    ch <- 123
    ch <- 456
    fmt.Println(<-ch)
    fmt.Println(<-ch)

    ch <- 789
    fmt.Println(<-ch)

    ch <- 111
    ch <- 222
    fmt.Println(<-ch)
    fmt.Println(<-ch)

    // OUTPUT:
    // 123
    // 456
    // 789
    // 111
    // 222
    //
    // Note that the output will ALWAYS be in the order of Sends we performed.

    // ---------------------------------------------
    //
    // We can also Receive from a channel before we perform a Send

    ch2 := make(chan int)

    var a int

    // However, because Receives not preceded by Sends will block
    // (because the Channel has no value to give), we need to run
    // our Receives on a separate thread with the "go" command
    // inside an anonymous function, where it will wait for the
    // channel to be populated.
    //
    go func() {
        a = <- ch2
    }()

    ch2 <- 2

    // This Print is safe because we will not reach it without
    // Sending the value 2 to ch2 (see above line), which will
    // trigger our thread above where "a" will Receive from ch2.
    //
    fmt.Println(a)
}
