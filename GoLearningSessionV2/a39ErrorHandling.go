package main

import (
    "fmt"
    "os"
    "errors"
    "strconv"
)

func main() {
    // we attempt to open a non-existent file in our Application's current directory
    f1, err := os.Open("missingFile.txt")

    if err == nil {
        fmt.Println(f1)
    } else {
        fmt.Println(err)
    }
    // BAD OUTPUT:
    // open missingFile.txt: The system cannot find the file specified.

    // --------------------------------------------
    //
    num, err := doubleValue("12")

    if err == nil {
        fmt.Println(num)
    } else {
        fmt.Println(err)
    }
    // OUTPUT:
    // 24

    num, err = (doubleValue("apples"))

    if err == nil {
        fmt.Println(num)
    } else {
        fmt.Println(err)
    }
    // BAD OUTPUT:
    // panic: Value is not a valid number.
}

// If you want to make your own error messages,
// you'll need to use the "errors" package.
//
func doubleValue(str string) (int, error) {
    num, err := strconv.Atoi(str)

    if err == nil {
        return num * 2, nil
    } else {
        return -1, errors.New("Value is not a valid number.")
    }
}