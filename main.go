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
	masterView, homeView, contactView *views.View
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

func master(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := masterView.Template.ExecuteTemplate(w, masterView.Layout, nil)
	if err != nil {
		panic(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(contactView.Render(w, nil))
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
	masterView = views.NewView("masterBootstrap", "views/master.gohtml")
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")

	r := mux.NewRouter()
	r.HandleFunc("/", master)
	r.HandleFunc("/timeless", home)
	r.HandleFunc("/timeless/contact", contact)
	r.HandleFunc("/timeless/faq", faq)
	r.NotFoundHandler = http.HandlerFunc(notFound)

	err := http.ListenAndServe(GetPort(), r)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func must(err erro) {
	if err != nil {
		panic(err)
	}
}
