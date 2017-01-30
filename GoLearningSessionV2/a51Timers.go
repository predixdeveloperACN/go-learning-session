package main

import (
    "fmt"
    "time"
)

func main() {
    // Timers are useful when you want perform an action *once* in the future.
    // They can be cancelled before they execute.

    // NewTimer creates a Time Channel after a specified interval (2 seconds)
    timer1 := time.NewTimer(time.Second * 2)

    // Here, we have a Receive waiting on a separate thread:
    go func() {
        <-timer1.C
        fmt.Println("Timer 2 executed")
    }()
    
    // This line executes before 2 seconds pass, so we cancel timer1 before it executes.
    //
    // Note that Stop always returns a Boolean value of True -- handy for initializing flags.
    // Also note that you can just call Stop by itself, without the need to assign it to a variable.
    //
    stop2 := timer1.Stop()
    
    if stop2 {
        fmt.Println("Timer 2 stopped")
    }

    // OUTPUT:
    // Timer 2 stopped
}
