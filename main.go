package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// GetPort defines a port for wild environment
func GetPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "3000"
		fmt.Println("INFO: No PORT environment detected", port)
	}
	return ":" + port
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch, please send an email to <a href=\"mailto:support@fotok.com\">support@fotok.com</a>.")
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>This is the page of FAQ</h1>")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>404 Page not Found</h1>")
}

func main() {

	// Router from the standard library
	// mux := &http.ServeMux{}
	// mux.HandleFunc("/", handlerFunc)

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	r.NotFoundHandler = http.HandlerFunc(notFound)

	err := http.ListenAndServe(GetPort(), r)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
