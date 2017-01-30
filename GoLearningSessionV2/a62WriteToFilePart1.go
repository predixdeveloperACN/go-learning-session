package main

import (
    "fmt"
    "io"
    "os"
)

func main() {
    // This sample program writes a string into a text file
    // and prints how many characters it wrote into said file.

    content := "Hello, world!"

    // By default, the directory used by Go is the location of your Main .go file
    // If you want to use a custom filepath, make sure to also include the filename.
    file, err := os.Create("myfile.txt")
    checkError(err)
    
    // don't forget to close the connection!
    defer file.Close()

    // WriteString will always overwrite the contents of a file with new content.
    // It also returns the number of characters it wrote into the file by default
    length, err := io.WriteString(file, content)
    checkError(err)

    fmt.Printf("Done! %v characters written to file!", length)

    // OUTPUT:
    // Done! 13 characters written to file!
    //
    // ------------------------------------
    //
    // Inside myfile.txt (right next to this .go file):
    // Hello, world!
}

// This handy function stops the application and prints the error.
// We usually use "panic" when we have errors that shouldn't occur during
// normal operation, or ones that we aren't prepared to handle gracefully.
func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
