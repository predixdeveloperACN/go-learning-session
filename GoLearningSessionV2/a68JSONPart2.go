package main

import (
    "fmt"
    "os"
    "encoding/json"
)

// We'll use these two structs to demonstrate
// encoding and decoding of custom types below.

type Response1 struct {
    Page   int
    Fruits []string
}

// Go lets you use Tags for JSON Struct Fields. Here are some:
//
// Field is ignored by this package.
//     Field int `json:"-"`
//
// Field appears in JSON as key "myName".
//     Field int `json:"myName"`
//
// Field appears in JSON as key "myName" and the field is omitted from
// the object if its value is empty, as defined above.
//     Field int `json:"myName,omitempty"`
//
// Field appears in JSON as key "Field" (the default), but the field is skipped if empty.
// Also note the leading comma.
//      Field int `json:",omitempty"`
//
// XML format:
//      Field int `xml:"myName,attr"`
//
// Handle both JSON and XML:
//      Field int `xml:"myName,attr" json:"myName"`

// Note how we use a backtick (`) instead of ' or " for our tags.
// This is because ' is only for single characters,
// while " would require several character escape sequences
// that would make our code unreadable.
//
// NOTE: ` is usually located at the top-left of your keyboard,
//       right below Esc and on the left side of the 1 key
type Response2 struct {
    Page   int      `json:"page"`
    Fruits []string `json:"fruits"`
}

func main() {
    // First we'll look at encoding basic data types to JSON strings.
    // Here are some examples for simple values:

    bolB, _ := json.Marshal(true)
    fmt.Println(string(bolB))   
    // OUTPUT:
    // true
    
    intB, _ := json.Marshal(1)
    fmt.Println(string(intB))
    // OUTPUT:
    // 1

    fltB, _ := json.Marshal(2.34)
    fmt.Println(string(fltB))
    // OUTPUT:
    // 2.34
    
    strB, _ := json.Marshal("gopher")
    fmt.Println(string(strB))
    // OUTPUT:
    // "gopher"

    // ---------------------------------------------
    //
    // Here are some examples for slices and maps,
    // which encode to JSON arrays and objects as you'd expect.

    slcD := []string{"apple", "peach", "pear"}
    slcB, _ := json.Marshal(slcD)
    fmt.Println(string(slcB))
    // OUTPUT:
    // ["apple","peach","pear"]

    mapD := map[string]int{"apple": 5, "lettuce": 7}
    mapB, _ := json.Marshal(mapD)
    fmt.Println(string(mapB))
    // OUTPUT:
    // {"apple":5,"lettuce":7}

    // ----------------------------------------
    //
    // The JSON package can also encode your custom data types (like Structs).
    // It will only include Public fields in the encoded output and will
    // by default use those names as the JSON keys.
    res1D := &Response1{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"},
    }

    res1B, _ := json.Marshal(res1D)
    fmt.Println(string(res1B))
    // OUTPUT:
    // {"Page":1,"Fruits":["apple","peach","pear"]}

    // You can use tags when you declare your struct fields to customize the
    // encoded JSON key names. Check the definition of `Response2` above
    // to see an example of such tags.
    res2D := &Response2{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"},
    }

    res2B, _ := json.Marshal(res2D)
    fmt.Println(string(res2B))
    // OUTPUT:
    // {"page":1,"fruits":["apple","peach","pear"]}

    // ---------------------------------------------
    //
    // Now let's look at decoding JSON data into Go values.
    // Here's an example for a generic data structure,
    // composed of a float and a slice of strings.
    //
    // NOTE: We need to put our JSON string into an array of bytes
    //       to be able to Unmarshal them later.
    //
    byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

    // We need to provide a variable where the JSON package can put the decoded data.
    // We'll use a map that will hold keys of Strings and values of Unknown Data Types.
    //
    // Note how we once again use the Blank Interface to handle unknown types.
    //
    var dat map[string]interface{}

    // Here, we put each key-value pair of our JSON byte slice into our map.
    if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }

    fmt.Println(dat)
    // OUTPUT:
    // map[num:6.13 strs:[a b]]

    // In order to use the values in our map, we'll need to cast them to the correct type.
    // Here, we cast "num" to float64
    num := dat["num"].(float64)
    fmt.Println(num)
    // OUTPUT:
    // 6.13

    // Accessing elements from nested data (e.g. slices) requires two casts --
    // one to cast to a slice of unknown data types,
    // and another one to cast to the correct data type.
    //
    strs := dat["strs"].([]interface{})
    fmt.Println(strs[0].(string))
    fmt.Println(strs[1].(string))
    // OUTPUT:
    // a
    // b

    // ---------------------------------------------
    //
    // We can also decode JSON into custom data types (like Structs).
    // This has the advantage of adding type-safety to our programs and
    // eliminating the need for type assertions when accessing the decoded data.

    str := `{"page": 1, "fruits": ["apple", "peach"]}`
    res := Response2{}

    json.Unmarshal([]byte(str), &res)

    fmt.Println(res)
    // OUTPUT:
    // {1 [apple peach]}

    fmt.Println(res.Fruits[0])
    // OUTPUT:
    // apple

    // ---------------------------------------
    //
    // In the examples above, we always used bytes and strings as
    // intermediates between the data and JSON representation on standard out.
    //
    // We can also stream JSON encodings directly to os.Writers like "os.Stdout"
    // or even HTTP response bodies:
    //
    enc := json.NewEncoder(os.Stdout)

    d := map[string]int{"apple": 5, "lettuce": 7}

    enc.Encode(d)
    // OUTPUT:
    // {"apple":5,"lettuce":7}
}