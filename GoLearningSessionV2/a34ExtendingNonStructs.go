package main

import (
    "fmt"
)

// You can also "extend" non-struct types with your own custom functions.
//
// To do this, you first have to create an "alias" for the type you want to extend.
//
type MyFloat float64

// Here, we "extend" the float64 data type with
// our custom function by "tying" it to our alias:
//
func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }

    return float64(f)
}

func main() {
    f := MyFloat(-123.456)

    fmt.Println(f.Abs())
    // OUTPUT:
    // 123.456
}
