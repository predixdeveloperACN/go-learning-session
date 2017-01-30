package main

import (
    "fmt"
    "time"
)

func main() {

    c1 := make(chan string, 1)

    // this channel will be populated in 2 seconds
    go func() {
        time.Sleep(time.Second * 2)
        c1 <- "result 1"
    }()

    // This Select has a Timeout Case that will execute after 1 second,
    // which is faster than the population of the channel (which takes 2 seconds),
    // so the expected output is "timeout 1".
    //
    // If we increased the timeout, or populated the channel faster,
    // then the channel case will execute instead and we will have an output of "result 1".
    //
    select {
    case res := <-c1:
        fmt.Println(res)
    case <-time.After(time.Second * 1):
        fmt.Println("timeout 1")
    }
    // OUTPUT:
    // timeout 1
    
    // ---------------------------------------------
    //
    c2 := make(chan string)

    // this channel will also be populated in 2 seconds
    go func() {
        time.Sleep(time.Second * 2)
        c2 <- "result 2"
    }()
    
    // This select has a longer timeout than the channel (3 seconds),
    // so the channel will finish first. If the channel and the timeout
    // finished at the same time, select will choose one at RANDOM.
    //
    // Note that for the select to choose one at random, the c2 channel MUST be Unbuffered.
    // It has something to do with Go's Channel Buffer locking implementation,
    // which will not be covered in this course.
    //
    select {
    case res := <-c2:
        fmt.Println(res)
    case <-time.After(time.Second * 3):
        fmt.Println("timeout 2")
    }
    // OUTPUT:
    // result 2
}