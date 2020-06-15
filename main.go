package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"timeless/views"

	"github.com/gorilla/mux"
)

var (
	homeView, contactView *views.View
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
	err := homeView.Template.ExecuteTemplate(w, homeView.Layout, nil)
	if err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := contactView.Template.ExecuteTemplate(w, contactView.Layout, nil)
	if err != nil {
		panic(err)
	}
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>This is the page of FAQ</h1>")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>404 Page not Found</h1>")
}

func main() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")

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
