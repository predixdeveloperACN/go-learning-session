package main

import (
	"fmt"
	"math"
)

type Vertex struct {
    X, Y float64
}

// You don't declare methods inside a Struct.
// Instead, you "tie" them to a Struct data type.
//
func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Here's what it would look like if Abs was NOT a Method of Vertex:
//
// It turns into a general-purpose function that takes a Vertex parameter
// instead of being "bound" to the Vertex struct.
//
//     func Abs(v Vertex) float64 {
//         return math.Sqrt(v.X*v.X + v.Y*v.Y)
//     }

func main() {
    v := Vertex{3, 4}

    fmt.Println(v.Abs())
	// OUTPUT:
	// 5

    // Here's how we use the general-purpose version of our function:
    // NOTE: The Output is the same.
    //
    //     fmt.Println(Abs(v))
}
