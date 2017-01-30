package main

import (
    "fmt"
)

func main(){
    fmt.Println("Program start!")

    goto endOfProgram
    
    fmt.Println("This will NOT be printed!")

    endOfProgram:
        fmt.Println("Program end.")

    // OUTPUT:
    // Program start!
    // Program end.
}
