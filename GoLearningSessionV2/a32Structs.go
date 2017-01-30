package main

import (
    "fmt"
)

// Go does not have Classes -- instead, it uses Structs.
// Because there are no Classes, there is NO inheritance.

// Privacy in Go only works on the Package level.
// Anything Private cannot be used by Importers of a Package.
// However, anything Private can *STILL* be used within the Package -- including Private fields.

// Structs are a collection of Fields.
type Dog struct {
    Breed  string // Public field
    weight int    // private field
}

type Vertex struct {
	X, Y int
}

func main() {
    d := Dog{"Poodle", 34}
    
    fmt.Println(d)
    // OUTPUT:
    // {Poodle 34}
    
    fmt.Printf("%+v \n", d)
    // OUTPUT:
    // {Breed:Poodle Weight:34}

    fmt.Printf("Breed: %v \nWeight: %v \n", d.Breed, d.weight)
	// OUTPUT:
	// Breed: Poodle
	// Weight: 34

	// ---------------------------------------------
    //
	// If you don't initialize all Struct Fields,
	// they will get the zero-value for their data type.
    //
	var (
		v1 = Vertex{1, 2}  // initialize all fields

		v2 = Vertex{X: 1}  // Y is 0 by default
		v3 = Vertex{}      // X and Y are both 0 by default

		p  = &Vertex{3, 4} // initialize a Pointer to a newly-created Vertex
	)

	fmt.Println(v1)
    // OUTPUT:
    // {1 2}

	fmt.Println(v2)
    // OUTPUT:
    // {1 0}

    fmt.Println(v3)
    // OUTPUT:
    // {0 0}

    fmt.Println(*p)
    // OUTPUT:
    // {3 4}

	// ---------------------------------------------
    //
	// You can use Pointers to modify Structs

	v := Vertex{5, 6}

	p = &v

	p.X = 10

	fmt.Println(v)
	// OUTPUT:
	// {10 6}

    // ---------------------------------------------
    //
    // NOTE 1: You can't use implicit declarations with Struct fields
    //         (e.g. using the Return value of a function to initialize them):
    //
    //     var v1 Vertex
    //
    //     v1.X, v1.Y := ReturnTwoNumbers()
    //
    // OUTPUT (errors):
    // non-name v1.X on left side of :=
    // non-name v1.Y on left side of :=
    //
    // NOTE 2: Just use the "=" sign instead:
    //
    //     var v1 Vertex
    //
    //     v1.X, v1.Y = ReturnTwoNumbers()
}
