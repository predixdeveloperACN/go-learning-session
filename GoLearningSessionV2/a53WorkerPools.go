package main

import (
    "fmt"
    "time"
)

// We'll be running several concurrent instances of the worker function.
// These "workers" will receive work on the 'jobs' channel and send the corresponding
// results on the 'results' channel. We'll sleep a second per job to simulate an expensive task.
//
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "processing job", j)
        time.Sleep(time.Second)
        results <- j * 2 // we just multiply any value we're given to 2
    }
}

func main() {
    // In order to use our pool of workers, we need to send them work and collect their results.
    // We make 2 channels for this.
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    // This starts up 3 workers, each on their own thread, initially blocked because there are no jobs yet
    // (their range loops will wait for "jobs" to be populated).
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    // Here we send 9 'jobs' and then 'close' that channel
    // to indicate that that's all the work we have.
    for j := 1; j <= 9; j++ {
        jobs <- j
    }
    close(jobs)

    // Finally, we collect all the results of the work.
    //
    // Note that we're just tossing our results into a void since all
    // we're just using the For Loop below to keep our program from closing early. :P
    //
    // If we don't Receive from the results channel, our program will
    // terminate (even while workers are working) because our
    // workers are on their own thread.
    //
    for a := 1; a <= 9; a++ {
        <-results
    }
    
    // Our running program shows the 9 jobs being executed by various workers.
    // The program only takes about 3 seconds despite doing about 9 seconds of
    // total work because there are 3 workers operating concurrently.
    //
    // If you run this program, the 9 jobs will be printed in sets of 3
    //
    // SAMPLE OUTPUT:
    // worker 3 processing job 2
    // worker 2 processing job 3
    // worker 1 processing job 1
    // worker 2 processing job 4
    // worker 3 processing job 5
    // worker 1 processing job 6
    // worker 3 processing job 7
    // worker 1 processing job 9
    // worker 2 processing job 8

    // NOTE: Sometimes it's better to have multiple worker pools that specialize in one specific task.
    //
    //       Say we had a fast-food restaurant that serves food to customers (as all restaurants do).
    //       One implementation would be to have one large group of workers that:
    //
    //       Take Orders -> Cook Orders -> Serve Orders
    //
    //       Now, this is all well and good, but what if each of those 3 Tasks took a long time to complete?
    //       (e.g. Taking Orders = 1 minute, Cooking Orders = 5 minutes, Serving Orders = 3 minutes)
    //
    //       If we had 15 workers in our one large Worker Pool, each one would need to spend 8 minutes before
    //       being able to take the next order -- and nobody likes waiting in line at a fast-food restaurant.
    //
    //       Here, our workers are being "soft-blocked" while working on an expensive task, which is a sign that
    //       we are not using our worker pools efficiently.
    //
    //       In scenarios like these, it would be better to have separate worker pools for each Task.
    //       This approach lets us allocate workers to the most expensive tasks so the most amount of work
    //       gets done in the least amount of time.
    //
    //       Below, we allocate more workers to the Cooking and Serving Worker Pools because
    //       they take the longest to complete:
    //
    //       1. Worker Pool A (2 members) Takes Orders -> add to Orders Channel
    //       2. Orders Channel -> taken by Worker Pool B (9 members), which Cooks Orders -> add to Cooked Channel
    //       c) Cooked Channel -> taken by Worker Pool C (4 members), which Serves Orders to customers
    //
    //       This way, Orders are taken, cooked, and then served at the best possible time our resources will allow.
}
