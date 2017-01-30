package main

import (
    "fmt"
    "strconv"
)

func main() {
    // The "Break" statement only stops the Loop/Switch/Select it belongs to.
    for i := 0; i < 10; i++ {
        fmt.Printf(strconv.Itoa(i) + " ")

        for j := 11; j < 20; j++ {
            fmt.Printf(strconv.Itoa(j) + " ")

            if j == 15 {
                break // this break only stops the inner loop
            }
        }

        fmt.Println()
    }

    // Below, we see that the outer loop (0-9) finishes
    // while the inner loop (11-19) only iterates up to 15
    // (where our "break" was).
    //
    // OUTPUT:
    // 0 11 12 13 14 15
    // 1 11 12 13 14 15
    // 2 11 12 13 14 15
    // 3 11 12 13 14 15
    // 4 11 12 13 14 15
    // 5 11 12 13 14 15
    // 6 11 12 13 14 15
    // 7 11 12 13 14 15
    // 8 11 12 13 14 15
    // 9 11 12 13 14 15

    // If using "return" is not an option, the next best thing is to
    // define a label, and "break" said label.
    //
    // NOTE: Use "Ctrl + /" to toggle commenting blocks of code!
    //
    loopToTen:
        for i := 0; i < 10; i++ {
            switch i {
            case 5:
                break loopToTen
            default:
                fmt.Printf(strconv.Itoa(i) + " ")
            }
        }

    // OUTPUT:
    // 0 1 2 3 4
}