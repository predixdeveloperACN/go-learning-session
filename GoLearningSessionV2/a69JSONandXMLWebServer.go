package main

import (
    "time"
    "io/ioutil"
    "encoding/xml"
    "net/http"
    "log"
    "encoding/json"
    "strings"
    "os"
    "fmt"
)

// our valid URLs
var GetURL = "/get/"
var PatchURL = "/patch/"
var WriteURL = "/write/"

// mock database object
var Flights = FlightList{}
var XMLFileName = "flights.xml"

type Flight struct {
    FlightNumber string    `xml:"flightnumber"`
    Departure    time.Time `xml:"departure"`
    Arrival      time.Time `xml:"arrival"`
}

type FlightList struct {
    Flight []Flight `xml:"flight"`
}

func main() {
    // read from XML file
    xmlData, err := ioutil.ReadFile(XMLFileName)
    checkError(err)

    // convert XML data to Struct
    err = xml.Unmarshal(xmlData, &Flights)
    checkError(err)

    // handle URLs after "localhost:8080" that match our predefined ones
    http.HandleFunc(GetURL, Flights.returnJSONFlightList)
    http.HandleFunc(PatchURL, Flights.patchFlight)
    http.HandleFunc(WriteURL, Flights.writeXMLFlightList)

    // -------------------------------------------

    log.Print("Listening on Port 8080...")

    // blocks until the program is terminated
    http.ListenAndServe(":8080", nil)
}

// returns JSON Flight Slice
func (fl *FlightList) returnJSONFlightList(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(fl)
}

// Updates a Flight one field at a time.
// Also automatically writes to file once done.
//
// NOTE: Only handles FlightNumbers at the moment.
//       The first number is the FlightNumber to Update,
//       and the second one is the Update value.
//
// Sample URL: localhost:8080/patch/flightnumber=001,111
//
func (fl *FlightList) patchFlight(w http.ResponseWriter, r *http.Request) {
    updateStatement := r.URL.Path[len(PatchURL):]
    fields := strings.FieldsFunc(updateStatement, Delimeters)

    if fields[0] == "flightnumber" {
        for i := range fl.Flight {
            if fl.Flight[i].FlightNumber == fields[1] {
                fl.Flight[i].FlightNumber = fields[2]
                http.Redirect(w, r, WriteURL, http.StatusFound)
            }
        }
    }
}

func Delimeters(r rune) bool {
    match := (r == '=' || r == ',')

    return match
}

// converts FlightList to XML and writes it to an XML file
func (fl *FlightList) writeXMLFlightList(w http.ResponseWriter, r *http.Request) {
    strXML, err := xml.Marshal(fl)
    checkError(err)

    xmlFile, err := os.Create(XMLFileName)
    checkError(err)
    defer xmlFile.Close()

    xmlFile.Write(strXML)

    w.Write([]byte("FlightList saved successfully!"))
}

func checkError(err error) {
    if err != nil {
        fmt.Print(err)
    }
}
