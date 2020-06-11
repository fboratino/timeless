package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch, please send an email to <a href=\"mailto:support@fotok.com\">support@fotok.com</a>.")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>We could not find the page you were looking for :(</h1><p>Please email us if you keep being sent to an invalid page.</p>")
	}
}

func main() {
	mux := &http.ServeMux{}
	mux.HandleFunc("/", handlerFunc)
	// http.HandleFunc("/", handlerFunc)
	err := http.ListenAndServe(GetPort(), mux)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

// GetPort defines a port for wild environment
func GetPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "3000"
		fmt.Println("INFO: No PORT environment detected", port)
	}
	return ":" + port
}
