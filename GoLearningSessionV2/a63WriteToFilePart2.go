package main

import (
    "io/ioutil"
)

func main() {
    // Below is an alternate method for writing to a text file

    content := "Hello, world!"

    bytes := []byte(content)

    // 0644 is to control permissions on the file being written
    ioutil.WriteFile("myfile2.txt", bytes, 0644)

    // Inside myfile2.txt (right next to this .go file):
    // Hello, world!
    //
    // -------------------------------------------------
    //
    // NOTE: Common Permission Usages
    //
    // 0755 Commonly used on web servers.
    //      The owner can read, write and execute. Everyone else can read and execute but not modify the file.
    //
    // 0777 Everyone can read, write and execute.
    //      On a web server, it is not advisable to use this permission for your files and folders,
    //      as it allows anyone to add malicious code to your server.
    //
    // 0644 Only the owner can read and write.
    //      Everyone else can only read. No one can execute the file.
    //
    // 0655 Only the owner can read and write, but not execute the file.
    //      Everyone else can read and execute, but cannot modify the file.
}
