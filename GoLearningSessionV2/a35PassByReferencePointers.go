package main

import (
    "fmt"
)

type Vertex struct {
    X, Y float64
}

// This function uses Pointers to Pass by Reference,
// meaning it will modify/affect the parameter passed to it.
//
func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

// Here's what it would look like if we
// turned it into a general-purpose function:
//
//     func Scale(v *Vertex, f float64) {
//         v.X = v.X * f
//         v.Y = v.Y * f
//     }

func main() {
    v := Vertex{3, 4}

    v.Scale(10)

    fmt.Println(v)
    // OUTPUT:
    // {30 40}

    // Here's how to use our general-purpose function:
    // (Note how we give a Pointer Address (via "&") instead of a Struct.)
    //
    //     Scale(&v, 10)
    //
    // You can also use a Struct Pointer instead:
    //
    //    ps := &v
    //    Scale(ps, 10)
}
