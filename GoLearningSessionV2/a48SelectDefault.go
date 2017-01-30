package main

import (
    "fmt"
    "time"
)

func main() {
    // The "default" case in a "select" is run if no other cases are ready.
    // This is useful when you have a Select inside an infinite loop and you
    // want to perform an action while no channels are ready.

    // Populate each channel at different intervals:
    //
    // NOTE: time.Tick and time.After both return Time Channels after a specified interval
    //
    tick := time.Tick(100 * time.Millisecond)
    boom := time.After(500 * time.Millisecond)
    
    // This infinite loop has a Select that will execute whenever
    // a channel is populated and prints "." while idle.
    // It ends when the "boom" channel is populated.
    for {
        select {
        case <-tick:
            fmt.Println("tick.")
        case <-boom:
            fmt.Println("BOOM!")
            return
        default:
            fmt.Println("    .")
            time.Sleep(50 * time.Millisecond) // if we don't Sleep here, we'll get LOTS of periods :P
        }
    }

    // OUTPUT:
    //     .
    //     .
    // tick.
    //     .
    //     .
    // tick.
    //     .
    //     .
    // tick.
    //     .
    //     .
    // tick.
    //     .
    //     .
    // BOOM!
}