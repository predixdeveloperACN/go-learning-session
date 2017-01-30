package main

import (
    "fmt"
    "time"
)

// Goroutines are lightweight threads managed by the Go runtime.
//
// NOTE: Goroutines are not thread-safe by default. You can use the sync package's data types,
//       although you won't need them as much in Go because there are other basic data types
//       (e.g. Channels) which we will learn more about later on in the course.

func main() {
    // We run code on its own thread with the "go" command.
    // Now, both functions below will run at the same time,
    // even though they Sleep for 100 ms.
    //
    go say("world")
    say("hello")
    // OUTPUT:
    // world
    // hello
    // world
    // hello
    // hello
    // world
    // world
    // hello
    // hello

    // If we didn't use the "go" command, the Output would be:
    // world
    // world
    // world
    // world
    // world
    // hello
    // hello
    // hello
    // hello
    // hello

    // ---------------------------------------------
    //
    // You have to be careful when using Closures inside loops,
    // because iteration variables are REUSED in each iteration.
    //
    // NOTE: The following is just a sample scenario. It's not recommended
    //       to use Goroutines with Slices if you plan to preserve the order of elements
    //       because it's impossible to tell which Goroutine will finish first.
    //       More examples of this Goroutine behavior later.
    //
    data := []string{"one","two","three"}

    for _,v := range data {
        // this Goroutine is wrong, because by the time the 3 Goroutines are created,
        // they are all pointing to the last item in the range loop, which is "three".
        //
        go func() {
            fmt.Println(v)
        }()
    }
    time.Sleep(3 * time.Second)
    // OUTPUT:
    // three
    // three
    // three

    // One solution is to copy the value of
    // the iteration variable into a local variable:
    //
    for _,v := range data {
        vcopy := v
        go func() {
            fmt.Println(vcopy)
        }()
    }
    time.Sleep(3 * time.Second)
    // OUTPUT:
    // one
    // two
    // three

    // Another solution is to pass the iteration variable
    // as a parameter to our Closure inside the Goroutine:
    //
    for _,v := range data {
        go func(str string) { //  <-- Closure (Anonymous function) signature
            fmt.Println(str)
        }(v) //  <-- we pass our iteration variable here
    }
    time.Sleep(3 * time.Second)
    // OUTPUT:
    // one
    // two
    // three

    // ---------------------------------------------
    //
    // Here's a different scenario using Structs.
    //
    // One reason this scenario may arise is when your object
    // has a very expensive function, which you want to run
    // on its own thread.

    animals := []animal{{"dog"},{"cat"},{"cow"}}

    for _,v := range animals {
        go v.sayName()
    }
    time.Sleep(3 * time.Second)
    // OUTPUT:
    // cow
    // cow
    // cow

    // Here, we use the value-copy solution:
    //
    for _, val := range animals {
        a := val
        go a.sayName()
    }
    time.Sleep(3 * time.Second)
    // OUTPUT:
    // dog
    // cat
    // cow

    // Here, we use the pass-as-argument solution:
    //
    for _, val := range animals {
        a := val
        go func(anim animal) {
            anim.sayName()
        }(a)
    }
    time.Sleep(3 * time.Second)
    // OUTPUT:
    // cat
    // dog
    // cow
}

type animal struct {
    name string
}

func (a *animal) sayName() {
    fmt.Println(a.name)
}

// this function sleeps for 100ms to simulate a long task
//
func say(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}