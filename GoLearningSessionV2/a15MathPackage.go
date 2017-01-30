package main

// In Go, you need to import Main Packages and Subpackages explicitly
import (
    "fmt"
    "math"
    "math/big"
)

func main() {
    // Complex Math operations require the "math" package.

    // These "big" data types (like Float, shown below) are
    // necessary to have PRECISION in decimal computations.
    // Otherwise, you'll LOTS of trailing decimals, even for 2 decimal values.
    //
    var b1, b2, b3, bigSum big.Float

    // You need to use the "SetFloat64" function to assign values to "big" data types.
    // You will need to cast your values if you want to use "="
    b1.SetFloat64(23.5)
    b2.SetFloat64(65.1)
    b3.SetFloat64(76.3)

    // You can chain calls to ".Add" like so:
    //
    // NOTE: The Add function actually expects Pointers, which we will learn later on in the course.
    //       In the meantime, we will pass a Pointer Address instead (via the "&" character).
    //       Once you learn about Pointers, you can substitute them here instead of using the "&" character.
    //
    bigSum.Add(&b1, &b2).Add(&bigSum, &b3)

    // NOTE: "%.10g" means we want 10-digit precision (including whole numbers).
    //
    fmt.Printf("BigSum = %.10g\n", &bigSum)
    // OUTPUT:
    // BigSum = 164.9

    // NOTE: If you don't use "&" or a Pointer for Printing "big" data types,
    // you will get a BREAKDOWN of the variable, NOT its value.
    //
    // BAD OUTPUT:
    // BigSum = {%!g(uint32=0000000053) %!g(big.RoundingMode=0000000000)
    // %!g(big.Accuracy=-0000000001) %!g(big.form=0000000001) %!g(bool=false)
    // [%!g(big.Word=11882297256854315008)] %!g(int32=0000000008)}

    circleRadius := 15.5
    circumference := circleRadius * math.Pi

    fmt.Printf("Circumference: %.2f\n", circumference)
    // OUTPUT:
    // Circumference: 48.69
}
