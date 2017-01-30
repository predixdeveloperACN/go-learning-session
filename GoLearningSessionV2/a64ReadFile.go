package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    filename := "myfile.txt"

    content, err := ioutil.ReadFile(filename)
    checkError(err)

    // convert byte arrays to string
    result := string(content)

    fmt.Printf("Read from file:\n" + result)

    // OUTPUT:
    // Read from file:
    // Hello, world!

    // NOTE: If the file does not exist, we'll get an error:
    //
    // panic: open myfile.txt: The system cannot find the file specified.
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
