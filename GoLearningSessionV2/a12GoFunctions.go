package main

import (
    "fmt"
)

// There is NO method overloading in Go.
// All function signatures must be unique.

func main() {
    str1 := "Apples"
    str2 := "Oranges"

    // the built-in Println function returns both a string and an error object
    //
    tmpStr, err := fmt.Println(str1, str2)
    // OUTPUT:
    // Apples Oranges

    // NOTE: "nil" is Go's version of Null/Nothing
    if err != nil {
        panic(err)
    }

    fmt.Println(tmpStr)
    // OUTPUT:
    // Apples Oranges

    // You can use an underscore as a placeholder to ignore a returned value
    tmpStr, _ = fmt.Println(str1, str2)
    // OUTPUT:
    // Apples Oranges

    a, b := swap("hello", "world")

    fmt.Println(a, b)
    // OUTPUT:
    // world hello

    fmt.Println(doubleValues(3, 4))
    // OUTPUT:
    // 6 8
    
    sum := addAllValues(12, 54, 79)
    fmt.Println(sum)
    // OUTPUT:
    // 145
}

// Go functions can return more than 1 result
// This function returns 2 strings
//
func swap(x, y string) (string, string) {
    return y, x
}

// "naked" returns (also called "lazy" returns)
// are functions that specify which variables it will return beforehand.
// "return" is called on its own
//
// Note how x and y are not explicitly declared.
// This is because they were declared in the function signature
//
func doubleValues(num1, num2 int) (x, y int) {
    x = num1 * 2
    y = num2 * 2
    return
}

// This function accepts an unknown number of arguments of the same type.
// Note how it uses Slices (which we will learn more about later in the course)
//
func addAllValues(values ...int) int {
    sum := 0

    // We use "range" to iterate over Slices.
    // We will also learn more about "range" later on in the course.
    //
    for i := range values{
        sum += values[i]
    }

    return sum
}
