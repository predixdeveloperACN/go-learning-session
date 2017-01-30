package main

import (
    "fmt"
    "strconv"
)

// The purpose of Interfaces is to basically have a guarrantee that a certain object
// implements certain functions, meaning we can call said object's functions and expect
// the same Type of result, without having to worry about HOW said object arrived at said result.
//
// The "Efficient Go" section on the official site describes it as:
//
//     ‘if something can do this, then it can be used here.’
//
// Interfaces are always implemented implicitly.
// There is NO explicit way to say that something "implements" an interface.

type Human interface {
    Add(a, b int) string
}

type Child struct {}

// NOTE: You can't directly cast Numbers to Strings.
//       You must use the strconv package's "Itoa" (int to string) function.
//
func (c Child) Add(a, b int) string {
    // children just repeat the first number they hear
    return "Child: " + strconv.Itoa(a + 0)
}

type Adult struct {}

func (ad Adult) Add(a, b int) string {
    // an adult knows how to add properly
    return "Adult: " + strconv.Itoa(a + b)
}

type SeniorCitizen struct {}

func (s SeniorCitizen) Add(a, b int) string {
    // senior citizens just return 0
    return "SeniorCitizen: " + strconv.Itoa(0)
}

func main() {
    // Here's how to declare a variable of a Type that
    // implements a certain Interface:
    child1 := Human(Child {})

    fmt.Println(child1.Add(2, 3))
    // OUTPUT:
    // Child: 2

    // Here's how to declare a slice of structs
    // that implement a certain Interface:
    //
    humans := []Human {Child{}, Adult {}, SeniorCitizen{}}

    for _, human := range humans {
        fmt.Println(human.Add(3, 4))
    }
    // OUTPUT:
    // Child: 3
    // Adult: 7
    // SeniorCitizen: 0
}
