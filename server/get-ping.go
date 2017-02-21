package server

import (
	"fmt"
	"net/http"
	"strings"
)

func HandlePing(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// Always set content type and status code
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	message := []string{
		"{ }",
	}

	fmt.Fprintf(w, strings.Join(message, ""))
}
