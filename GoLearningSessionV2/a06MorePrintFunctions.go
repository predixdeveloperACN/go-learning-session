package main

import (
    "fmt"
)

func main() {
    tmpStr := "Quick brown fox"

    // Println automaticaly prints a Newline Character at the end.
    // You can use the built-in function "len" to find the length of a string.
    //
    fmt.Println("String length:", len(tmpStr))
    // OUTPUT:
    // String length: 15

    // NOTE 1: Go has MANY string formats for values. However, they
    //         are only usable with fmt.Printf and fmt.Sprintf
    //
    // NOTE 2: Printf formats according to a format specifier and writes to standard output.
    //         It returns the number of bytes written and any write error encountered.
    //
    //         Sprintf formats according to a format specifier and returns the resulting string.
    //
    //  Here are some:
    //
    //  %v -- the direct value
    // %+v -- when used with Structs, this will also include the field names
    //  %T -- the type of the value
    //  %d -- base-10 decimal format
    //  %b -- binary representation
    //  %c -- prints character corresponding to the given integer
    //  %x -- base-16 hex format
    //  %f -- basic decimal formatting
    //  %s -- basic string
    //  %q -- adds double quotes

    num1 := 123

    // Printf doesn't add a newline at the end, so you need to add one manually.
    // NOTE: Like in most languages, "\n" is used to add new lines
    //
    fmt.Printf("Value of a number: %v\n", num1)
    // OUTPUT:
    // Value of a number: 123

    // Format 2 decimal places
    fmt.Printf("%.2f %.2f \n", 1.2, 3.45)
    // OUTPUT:
    // 1.20 3.45

    // Sample Casting
    fmt.Printf("Value of a number as float: %.2f \n", float64(num1))
    // OUTPUT:
    // Value of a number as float: 123.00

    str1 := "Apple"
    var str2 string
    str3 := str1
    isCorrect := true
    
    // Print the Data Types
    fmt.Printf("Data types: %T, %T, %T, %T and %T \n",
               str1, str2, str3, num1, isCorrect)
    // OUTPUT:
    // Data types: string, string, string, int and bool

    // The "Print" function also prints without a new line.
    // This is useful if you don't plan on using string formats very much.
    //
    // Note how Intellij IDEA flags a possible mistake the user made with highlights --
    // Perhaps we meant to use a Printf instead? ;)
    //
    fmt.Print("Data types: %T, %T, %T, %T and %T",
        str1, str2, str3, num1, isCorrect)
    // OUTPUT:
    // Data types: %T, %T, %T, %T and %TAppleApple123 true
}
