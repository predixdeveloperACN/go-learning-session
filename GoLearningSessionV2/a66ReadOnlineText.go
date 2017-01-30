package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

func main() {
    // this url has a JSON string we want
    url := "http://services.explorecalifornia.org/json/tours.php"

    // get JSON from url
    response, err := http.Get(url)
    if err != nil {
        panic(err)
    }

    defer response.Body.Close()

    // just to show that we are getting an http Response type:
    fmt.Printf("Response type: %T\n", response)
    // OUTPUT:
    // Response type: *http.Response

    // convert http response body to a slice of bytes
    bytes, err := ioutil.ReadAll(response.Body)
    if err != nil {
        panic(err)
    }

    // convert byte slice to strings and print
    content := string(bytes)
    fmt.Print(content)

    // EXCERPT OF RESULTS:
    //
    // [{"tourId":"14","packageId":"5","packageTitle":"From Desert to Sea",
    // "name":"2 Days Adrift the Salton Sea","blurb":"The Salton Sea, 25% saltier than the Pacific,
    // has been a tourist destination since the 1920s. See what attracts people to this desert oasis.",
    // "description":"The Salton Sea is saltier than the Pacific, an unusual feat for inland body of water.
    // And even though its salinity has risen over the years, due in part to lack of outflows and pollution
    // from agricultural runoff, it has attracted a small, but dedicated population. The sea itself offers
    // recreational opportunities including boating, camping, off-roading, hiking, use of personal watercraft,
    // photography and bird watching. The sea has been termed a \"crown jewel of avian biodiversity,\" being a
    // major resting stop on the Pacific Flyway, a migratory path for birds. 2 Days Adrift the Salton Sea includes
    // two nights accommodations at the Bombay Beach Inn, boat rental at the Salton City Harbor, and a guided
    // fishing tour.","price":"350","difficulty":"2","length":"2","graphic":"map_saltonsea.gif","region":"Southern
    // California"},{"tourId":"26","packageId":"9","packageTitle":"Taste of California","name":"A Week of Wine",
}
