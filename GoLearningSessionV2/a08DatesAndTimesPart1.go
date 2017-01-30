package main

import (
    "fmt"
    "time"
)

func main() {
    // The complete format for declaring DateTimes in Go is:
    // year, month, date, hour, minute, seconds, nanoseconds, location
    //
    t1 := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)

    fmt.Printf("Go was launched at %s\n", t1)
    // OUTPUT:
    // Go was launched at 2009-11-10 23:00:00 +0000 UTC

    now := time.Now()

    fmt.Printf("The current time is %s\n", now)
    // OUTPUT:
    // The current time is 2017-01-04 11:21:29.6412876 +0800 CST

    fmt.Println("The month is", now.Month())
    fmt.Println("The day is", now.Day())
    fmt.Println("The weekday is", now.Weekday())
    // OUTPUT:
    // The month is January
    // The day is 4
    // The weekday is Wednesday

    // When adding to Dates, the order is:
    // (year, month, day)
    //
    tomorrow := now.AddDate(0, 0, 1)

    // You can customize your own formats, using the formatting constant of:
    //
    //     Monday, January 2, 3:04 PM, 2006 Mountain Standard Time (-07:00)
    //
    // Note how the pattern *almost* follows the sequence of 1, 2, 3 and so on.
    //
    longFormat := "Monday, January 2, 2006"
    shortFormat := "01/02/06"

    fmt.Println("Tomorrow is", tomorrow.Format(longFormat))
    fmt.Println("Tomorrow is", tomorrow.Format(shortFormat))
    // OUTPUT:
    // Tomorrow is Thursday, January 5, 2017
    // Tomorrow is 01/05/17
}
