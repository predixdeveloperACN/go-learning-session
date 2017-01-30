package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().Unix())

    // Intn returns a random int from 0 to the number provided.
    dayOfWeek := rand.Intn(6) + 1

    fmt.Println("Day of the Week:", dayOfWeek)
    // OUTPUT:
    // Day of the Week: 6

    // You can also use rand for random Floating values:
    fmt.Println("Today's lucky number:", rand.Float64(), "\n")
    // OUTPUT:
    // Today's lucky number: 0.3814709527493504

    // NOTE: If you seed a source with the same number, it will produce the same
    //       sequence of random numbers EVERY SINGLE TIME, including other Runs of the Application.
    //       This is why it's STRONGLY recommended to use the current time to seed your Randomizers.

    rand.Seed(4)
    fmt.Println(rand.Intn(456))
    fmt.Println(rand.Intn(456))
    fmt.Println(rand.Intn(456), "\n")
    // OUTPUT:
    // 181
    // 52
    // 253

    rand.Seed(4)
    fmt.Println(rand.Intn(456))
    fmt.Println(rand.Intn(456))
    fmt.Println(rand.Intn(456), "\n")
    // OUTPUT:
    // 181
    // 52
    // 253

    rand.Seed(5)
    fmt.Println(rand.Intn(456))
    fmt.Println(rand.Intn(456))
    fmt.Println(rand.Intn(456))
    // OUTPUT:
    // 450
    // 340
    // 217
}
