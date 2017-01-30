package main

import (
    "fmt"
    "sync"
    "time"
)

// We use the "sync" package whenever we want to ensure that only one Goroutine can
// access something at any given time to avoid conflicts. The concept is known as Mutual Exclusion.

type SafeCounter struct {
    // this Mutex makes this struct safe to use concurrently
    mux sync.Mutex
    v   map[string]int
}

// Inc just increments the Value for the given Key.
//
// The Lock and Unlock functions make it so only one
// goroutine can access our Struct at a time
//
func (c *SafeCounter) Inc(key string) {
    c.mux.Lock()
    
    c.v[key]++

    c.mux.Unlock()
}

// NOTE: If we changed:
//
//           mux sync.Mutex
//
//       Inside our SafeCounter struct to:
//
//           sync.Mutex
//
//       We can just write our Lock and Unlock codes like so:
//
//           c.Lock
//           c.Unlock
//
//       These are called Anonymous fields, and are only applicable to certain types,
//       such as other Structs and Mutexes, for example.
//       We will learn more about Anonymous Fields later on in the course.

// Value returns the current value for the given key.
func (c *SafeCounter) Value(key string) int {
    c.mux.Lock()

    // here's another use for the defer statement --
    // Unlocking after Locking a variable
    defer c.mux.Unlock()
    
    return c.v[key]
}

func main() {
    // Note how Mutex does not need to be initialized
    c := SafeCounter{v: make(map[string]int)}

    // Note that if we don't use Mutex to lock upon reading/writing,
    // we will get a fatal error for concurrent reads and writes to a map.
    //
    // In our case, no two threads can Write on the same map at the same time.

    for i := 0; i < 1000; i++ {
        go c.Inc("somekey")
    }

    // this just ensures our program has enough time to
    // reach the needed value (1000) before printing. :P
    time.Sleep(time.Second)

    fmt.Println(c.Value("somekey"))
}
