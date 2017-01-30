package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "strings"
    "math/big"
)

type Tour struct{
    Name, Price string
}

func main() {
    url := "http://services.explorecalifornia.org/json/tours.php"

    content := contentFromServer(url)

    tours := toursFromJson(content)

    for _, tour := range tours {
        // base 10, 2 decimal places, Rounding Mode
        price, _, _ := big.ParseFloat(tour.Price, 10, 2, big.ToZero)
        fmt.Printf("%v ($%.2f)\n", tour.Name, price)
    }

    // EXCERPT OF RESULTS:
    //
    // 2 Days Adrift the Salton Sea ($256.00)
    // A Week of Wine ($768.00)
    // Amgen Tour of California Special ($4096.00)
    // Avila Beach Hot springs ($768.00)
    // Big Sur Retreat ($512.00)
    // Channel Islands Excursion ($128.00)
    // Coastal Experience ($1024.00)
    // Cycle California: My Way ($1024.00)
    // Day Spa Package ($512.00)
    // Endangered Species Expedition ($512.00)
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

// reads JSON from the url and returns a string version
func contentFromServer(url string) string {
    resp, err := http.Get(url)
    checkError(err)

    defer resp.Body.Close()
    bytes, err := ioutil.ReadAll(resp.Body)
    checkError(err)

    return string(bytes)
}

func toursFromJson(content string) []Tour {
    tours := make([]Tour, 0)

    // create a new Decoder object and pass our JSON string to it
    decoder := json.NewDecoder(strings.NewReader(content))
    
    // Tell our Decoder to remove initial bracket from our JSON string
    //
    // NOTE: The Token function is actually doing more than that.
    //       It is actually making sure that our JSON string is properly formatted and nested.
    //
    _, err := decoder.Token()

    checkError(err)

    var tour Tour

    // loop through the decoder's JSON contents
    for decoder.More() {
        // try to decode each JSON element into our Tour Struct
        // by finding names that match our Tour Fields
        err := decoder.Decode(&tour)
        checkError(err)

        tours = append(tours, tour)
    }

    return tours
}
