package main

import (
    "fmt"
)

type animal struct {
    name string
    sound string
}

func (a animal) Speak() {
    fmt.Println(a.name + ": " + a.sound + "!")
}

type Vertex struct {
    Lat, Long float64
}

type Dog struct {
    ID     int
    Breed  string
}

func main() {
    // declare a slice of animals
    animals := []animal{
        {"dog", "woof"},
        {"cat", "meow"},
        {"cow", "moo"},
    }

    for _, animal := range animals {
        animal.Speak()
    }
    // OUTPUT:
    // dog: woof!
    // cat: meow!
    // cow: moo!

    // ---------------------------------------------
    //
    m := make(map[string]Vertex)

    m["Bell Labs"] = Vertex{
        40.68433, -74.39967,
    }

    fmt.Println(m["Bell Labs"])
    // OUTPUT:
    // {40.68433 -74.39967}

    // ---------------------------------------------
    //
    // You can also skip the "make" function altogether by
    // initializing your map like this:
    //
    // Note how Intellij IDEA dims the text for "Vertex".
    // This is because specifying the Struct Data Type is redundant.
    // You can omit it, as shown in the next section of this file.
    //
    n := map[string]Vertex{
        "Bell Labs": Vertex{
            40.68433, -74.39967,
        },
        "Google": Vertex{
            37.42202, -122.08408,
        },
    }

    fmt.Println(n)
    // OUTPUT:
    // map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]

    // ---------------------------------------------
    //
    // Here's the *Recommended* approach:
    //
    o := map[string]Vertex{
        "Amazon":    {11.22, -33.44},
        "Microsoft": {55.66, -77.88},
    }

    fmt.Println(o)
    // OUTPUT:
    // map[Amazon:{11.22 -33.44} Microsoft:{55.66 -77.88}]

    // ---------------------------------------------
    //
    // You can't use Maps to update Struct fields because they're not "addressable".
    //
    // This code will throw an error:
    //
    //     o["Amazon"].Lat = 10.20
    //
    // BAD OUTPUT:
    // cannot assign to struct field o["Amazon"].Lat in map

    // You can use Slices to update Struct fields instead.
    animals[0].sound = "Bow wow"

    animals[0].Speak()
    // OUTPUT:
    // dog: Bow wow!

    // ---------------------------------------------
    //
    // If you still want to use maps, one solution is to use a temporary variable:
    //
    // First, update the temporary variable.
    // Then, give it back to the map via Key.
    //
    v := o["Amazon"]
    v.Lat = 10.20
    o["Amazon"] = v

    fmt.Println(o["Amazon"])
    // OUTPUT:
    // {10.2 -33.44}

    // ---------------------------------------------
    //
    // Another solution is to use a map of Pointers:
    //
    p := map[string]*Vertex{
        "Japan":       {36.2048, 138.2529},
        "Philippines": {12.8797, 121.7740},
    }

    p["Japan"].Lat = 12345.678

    fmt.Println(p["Japan"])
    // OUTPUT: &{12345.678 138.2529}
    //
    // Note the "&" character. This is because we're using
    // a map of Pointers to Vertex Structs.

    // ---------------------------------------------
    //
    // You can use Struct objects as keys to maps:
    //
    m1 := make(map[Dog]string)

    d1 := Dog{1, "poodle"}

    m1[d1] = "Alice's dog"

    fmt.Println(m1[d1])
    // OUTPUT:
    // Alice's dog

    // ---------------------------------------------
    //
    // If you want to use Struct fields as keys,
    // use a Basic Data Type instead:

    m2 := make(map[int]string)

    d2 := Dog{2, "pomeranian"}

    m2[d2.ID] = "Bob's dog"

    fmt.Println(m2[d2.ID])
    // OUTPUT:
    // Bob's dog
}
