package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	homeTemplate, contactTemplate *template.Template
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
	if err := homeTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactTemplate.Execute(w, nil); err != nil {
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

	var err error
	homeTemplate, err = template.ParseFiles(
		"views/home.gohtml",
		"views/layouts/footer.gohtml",
	)
	if err != nil {
		panic(err)
	}

	contactTemplate, err = template.ParseFiles(
		"views/contact.gohtml",
		"views/layouts/footer.gohtml",
	)
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	r.NotFoundHandler = http.HandlerFunc(notFound)

	err = http.ListenAndServe(GetPort(), r)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
