package main

import (
    "fmt"
    "strings"
)

func main() {
    str1 := "the quick brown fox"
    str2 := "ThE QuIcK BrOwN FoX"

    fmt.Println(strings.ToUpper(str1))
    // OUTPUT:
    // THE QUICK BROWN FOX

    fmt.Println(strings.Title(str1))
    // OUTPUT:
    // The Quick Brown Fox

    // Use "==" for case-sensitive comparison.
    // Note the escape sequence (\") used to add literal double-quotes in the output string.
    //
    fmt.Printf("\"%v\" identical to \"%v\"? %v \n", str1, str2, str1 == str2)
    // OUTPUT:
    // "the quick brown fox" identical to "ThE QuIcK BrOwN FoX"? false

    // Use "strings.EqualFold" for non-case-sensitive comparison
    fmt.Printf("\"%v\" has same letters as \"%v\"? %v \n", str1, str2, strings.EqualFold(str1, str2))
    // OUTPUT:
    // "the quick brown fox" has same letters as "ThE QuIcK BrOwN FoX"? true

    strToSearch := "peaches and cream"
    strValue := "peach"

    // Use "strings.Contains" to check if a string exists within a string
    fmt.Printf("\"%v\" contains \"%v\"? %v",strToSearch, strValue, strings.Contains(strToSearch, strValue))
    // OUTPUT:
    // "peaches and cream" contains "peach"? true
}
