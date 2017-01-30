package main

import (
    "fmt"
    "time"
)

func main() {
    // Tickers are useful when you want to do something repeatedly at regular intervals.
    // They fire periodically until we stop them.

    // Tickers use a similar mechanism to timers -- a Time channel that is sent values.
    // Here, we'll use "range" to iterate over the channel as it gets populated every 500ms.
    ticker1 := time.NewTicker(time.Millisecond * 500)
    
    // This block runs on its own thread,
    // blocking as it waits for data to be Received from the channel
    go func() {
        for t := range ticker1.C {
            fmt.Println("Tick at", t)
        }
    }()

    // Tickers can be stopped like timers.
    // Once a ticker is stopped, it won't receive any more values on its channel.
    // Below, we stop ours after 1600ms.
    time.Sleep(time.Millisecond * 1600)

    ticker1.Stop()

    fmt.Println("Ticker1 stopped")

    // SAMPLE OUTPUT:
    // Tick at 2017-01-05 14:31:35.4911282 +0800 CST
    // Tick at 2017-01-05 14:31:35.9923645 +0800 CST
    // Tick at 2017-01-05 14:31:36.491347 +0800 CST
    // Ticker1 stopped

    // ---------------------------------------------
    //
    // Here's another example where we call a function every 500ms.

    ticker2 := time.NewTicker(time.Millisecond * 500)

    go func() {
        // NOTE: You can completely omit the assignment portion of a range loop
        //
        for range ticker2.C {
            SayHello()
        }
    }()

    time.Sleep(time.Millisecond * 1600)

    ticker2.Stop()

    fmt.Println("Ticker2 stopped")

    // OUTPUT:
    // Hello!
    // Hello!
    // Hello!
    // Ticker2 stopped
}

func SayHello() {
    fmt.Println("Hello!")
}
