package main

// for the usual collecting of user input
import (
    "bufio"
    "os"
    "strconv" // for Parsing strings
    "strings" // for removing invisible characters like \r and \n before parsing
    "fmt"
)

func main() {
    // The fmt package has several Scan functions, but we won't be focusing on them here.

    // This NewReader object will be used to accept user input
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter your name: ")
    str1, _ := reader.ReadString('\n')

    fmt.Print("Enter your favorite number: ")
    str2, _ := reader.ReadString('\n')

    // You MUST use TrimSpace or else the input will also
    // include "\r\n" which will cause the parse to fail.
    //
    float1, err := strconv.ParseFloat(strings.TrimSpace(str2), 64) // 64-bit precision

    if err != nil {
        panic(err)
    } else {
        // remove newline characters from str1 and cast float64 back to string
        // (for example purposes only, and because Go requires Explicit casting)
        fmt.Println("Hello, " + strings.TrimSpace(str1) +
                    "! I hear you like the number " +
                    strconv.FormatFloat(float1, 'f', 2, 64) + "!")
    }

    // SAMPLE OUTPUT:
    // Hello, Vince! I hear you like the number 24.00!
}