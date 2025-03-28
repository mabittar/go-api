package handlers

import (
	"fmt"
	"net/http"
    "strings"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	headers := r.Header
	var headersString strings.Builder

    // Convert headers map to a string
    for key, values := range headers {
        for _, value := range values {
            headersString.WriteString(fmt.Sprintf("%s: %s\n", key, value))
        }
    }
	response := "Hello, Golang API\n" + headersString.String()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, response)
}
