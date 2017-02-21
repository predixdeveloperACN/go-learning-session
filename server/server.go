package server

import (
	"github.com/predixdeveloperACN/go-learning-session/vcap"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"fmt"
)

var router *mux.Router

const base_path = "/api/v1/"
var response_path string

func SetupServer() {
	router = mux.NewRouter()
	router.Methods("GET").Path(base_path  + "ping").HandlerFunc(HandlePing)
	router.Methods("POST").Path(base_path  + "dog").HandlerFunc(HandleAddDog) // C
	router.Methods("GET").Path(base_path  + "dogs").HandlerFunc(HandleGetDogs) // R
	router.Methods("PUT").Path(base_path  + "dog").HandlerFunc(HandleUpdateDog) // U
	router.Methods("DELETE").Path(base_path  + "dog").HandlerFunc(HandleDeleteDogs) // D

}

func StartRestServer(portString string) {

	// Get Cloud Foundry assigned port
	port := os.Getenv("PORT")
	if port == "" {
		port = portString
	}

	vcapAppMap, _ := vcap.LoadApplication()

	var uri string
	if len(vcapAppMap.URIs) > 0 {
		uri = vcapAppMap.URIs[0]
	}
	if uri == "" {
		response_path = "localhost:" + port + base_path
	} else {
		response_path = "https://" + uri + base_path
	}

	log.Println(fmt.Sprintf("Starting REST Server Interface for", uri))
	log.Println(fmt.Sprintf("response_path = %s", response_path))
	log.Println("Port = 0x%s", port)

	// Start listening on the configured port.
	// ListenAndServe is not expected to return, so we wrap it in a log.Fatal
	err:= http.ListenAndServe(":"+port, router)
	if(err != nil){
		os.Exit(1)
	}

}
