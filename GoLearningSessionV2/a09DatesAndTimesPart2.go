package main

import (
    "fmt"
    "time"
)

func main() {
    // you can assign functions to variables and use them
    // as handy shortcuts!
    p := fmt.Println
    
    today := time.Now()

    // Printing basic Date/Time fields:
    p(today.Year())
    p(today.Month())
    p(today.Day())
    p(today.Hour())
    p(today.Minute())
    p(today.Second())
    p(today.Nanosecond())
    p(today.Location())
    p(today.Weekday())
    // SAMPLE OUTPUT:
    // 2017
    // January
    // 4
    // 11
    // 22
    // 28
    // 777622100
    // Local
    // Wednesday

    tomorrow := today.AddDate(0, 0, 1)

    // These methods compare two times, testing if the first one occurs
    // before, after, or at the same time as the second one.
    p(tomorrow.Before(today))
    p(tomorrow.After(today))
    p(tomorrow.Equal(today))
    // OUTPUT:
    // false
    // true
    // false

    // The Sub method takes in a Time and returns a Duration
    // representing the interval between the two Times.
    diff := today.Sub(tomorrow)

    // We can compute the length of the duration in various units.
    // Note that we have Negative values in our example because today is less than tomorrow
    p(diff.Hours())
    p(diff.Minutes())
    p(diff.Seconds())
    p(diff.Nanoseconds())
    // OUTPUT:
    // -24
    // -1440
    // -86400
    // -86400000000000
    //
    // Note that all of these values are equivalent to 24 Hours because
    // they are just a representation of the Duration in said time portion.
    //
    // ----------------------------------------------

    // You can use Add to advance a time with a positive Duration,
    // or move it backwards with a negative Duration.
    //
    // Note that using "-" on an already negative duration will make it positive.
    //
    // The following code will actually subtract time, since we are adding a Negative Duration.
    //
    p(tomorrow.Add(diff))
    // OUTPUT:
    // 2017-01-04 11:22:28.7776221 +0800 CST

    // this one will add time, since we are adding a Negative Duration and *Negating* the Negative
    p(tomorrow.Add(-diff))
    // OUTPUT:
    // 2017-01-06 11:22:28.7776221 +0800 CST
}
