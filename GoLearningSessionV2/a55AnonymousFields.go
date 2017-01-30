package main

import (
    "fmt"
)

// Anonymous fields are useful when you want to embed a Struct within another Struct.
// They let you access Properties (as well as Functions) of the Anonymous field
// as if they were properties of the Parent Struct.

type Room struct {
    numOfWindows int
}

func (r Room) OpenWindow() {
    fmt.Println("A window was opened")
}

type Hotel struct {
    numOfDoors int

    Room // this Anonymous field lets us access
         // numOfWindows as if it belonged to the Hotel struct.
         // It also lets us use functions tied to
         // the Room struct directly from a Hotel object.
}

// ---------------------------------------------
//
type Bedroom struct {
    numOfLamps int
    numOfBeds int
}

type DiningRoom struct {
    numOfPlates int
}

type Kitchen struct {
    numOfPlates int
}

type House struct {
    numOfLamps int // This field "hides" the numOfLamps in Bedroom.
                   // If we want to access Bedroom's numOfLamps, we need to do so Explicitly.

    Bedroom
    DiningRoom     // Since DiningRoom and Kitchen both have numOfPlates, we need to
    Kitchen        // Explicitly state which Anonymous field we plan to get numOfPlates from.
                   // Otherwise, we will get an "ambiguous selector" error.
}

func main() {
    // we need to Explicitly use the Type name when initializing Anonymous fields
    hotel := Hotel{1, Room{2}}

    fmt.Printf("The hotel has %v windows and %v door \n", hotel.numOfWindows, hotel.numOfDoors)
    // OUTPUT:
    // The hotel has 2 windows and 1 door

    hotel.OpenWindow()
    // OUTPUT:
    // A window was opened

    // ---------------------------------------------
    //
    h := House{2, Bedroom{3,1}, DiningRoom{4}, Kitchen{10}}

    // However, it's always good practice to explicitly access the
    // properties you want to avoid confusion (and errors) later on
    // when your code grows and the possibility of having same-name fields rises.
    //
    fmt.Printf("The house has %v lamps \n", h.numOfLamps)
    // OUTPUT:
    // The house has 2 lamps

    fmt.Printf("The bedroom has %v lamps and %v bed \n", h.Bedroom.numOfLamps, h.Bedroom.numOfBeds)
    // OUTPUT:
    // The bedroom has 3 lamps and 1 bed

    fmt.Printf("The dining room has %v plates \n", h.DiningRoom.numOfPlates)
    // OUTPUT:
    // The dining room has 4 plates

    fmt.Printf("The kitchen has %v plates \n", h.Kitchen.numOfPlates)
    // OUTPUT:
    // The kitchen has 10 plates
}
