
// Each Go file is composed of 3 parts:
//
//     1. The Package where the code belongs
//     2. Import Statements (if any)
//     3. The actual code
//
// The entry point for Go programs is the "main" function inside the "main" package.
// The main function will ONLY run when it is inside the main package.

package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello, world!")
    // OUTPUT:
    // Hello, world!
}
