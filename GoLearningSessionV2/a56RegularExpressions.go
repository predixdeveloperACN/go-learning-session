package main

import (
    "fmt"
    "regexp"
    "bytes"
)

func main() {
    // This tests whether a pattern matches a string.
    match, _ := regexp.MatchString("p([a-z]+)ch", "peach")

    fmt.Println(match)
    // OUTPUT:
    // true

    // Above, we used a string pattern directly,
    // but for more complex regex tasks,
    // you’ll need to Compile a Regexp object.
    r, _ := regexp.Compile("p([a-z]+)ch")

    // I recommend using MustCompile when declaring regular expressions,
    // because they will throw a Panic if the regular expression is invalid.
    // Also, they're cleaner to look at since we don't need an "_"
    //
    r2 := regexp.MustCompile("p([a-z]+)ch")

    // Many methods are available for regexp objects.
    // Here’s a match test like we saw earlier:
    fmt.Println(r.MatchString("peach"))
    // OUTPUT:
    // true

    // This finds the first match for the regexp.
    fmt.Println(r2.FindString("peach punch"))
    // OUTPUT:
    // peach

    // This also finds the first match, but returns the
    // start and end indexes instead of the matching text.
    fmt.Println(r2.FindStringIndex("peach punch"))
    // OUTPUT:
    // [0 5]

    // The "All" variants of these functions apply to all matches in the string,
    // not just the first. Providing a non-negative integer as the second argument to
    // these functions will limit the number of matches, while a negative will return ALL matches
    // (see below)
    fmt.Println(r2.FindAllString("peach punch pick pinch", -1))
    // OUTPUT:
    // [peach punch pinch]

    fmt.Println(r2.FindAllString("peach punch pick pinch", 2))
    // OUTPUT:
    // [peach punch]

    fmt.Println(r2.FindAllStringSubmatchIndex("peach punch pinch", -1))
    // OUTPUT:
    // [[0 5 1 3] [6 11 7 9] [12 17 13 15]]

    // The regexp package can also be used to replace subsets of strings with other values.
    fmt.Println(r2.ReplaceAllString("a peach", "pineapple"))
    // OUTPUT:
    // a pineapple

    // This function allows you to transform matched text with the results of a given function:
    in := []byte("a peach")
    out := r2.ReplaceAllFunc(in, bytes.ToUpper)
    fmt.Println(string(out))
    // OUTPUT:
    // a PEACH
}