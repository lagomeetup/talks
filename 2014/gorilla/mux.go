package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"html/template"
	"log"
	"net/http"
)

var index *template.Template

type Page struct{ Title, Body string }

func init() {
	var err error
	index, err = template.ParseFiles("./assets/templates/index.html")
	if err != nil {
		log.Fatal(err)
	}
}

func router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/foo", FooHandler)
	r.HandleFunc("/foo/new", NewFooHandler).Methods("POST", "PUT")
	r.HandleFunc("/api/v1/foo/{idOrSomething}", APIHandler).
		Methods("GET").
		Headers("X-Api-Version", "1")

	r.HandleFunc("/contact/new", NewConctactHandler).Methods("POST", "PUT")
	http.Handle("/assets/", http.StripPrefix("/assets/",
		http.FileServer(http.Dir("./assets/"))))
	return r
}

func FooHandler(resp http.ResponseWriter, req *http.Request) {
	index.Execute(resp, Page{
		"Foo Handler",
		"You can view this page",
	})
}

type Foo struct {
	Bar string `schema:"value"`
}

func NewFooHandler(resp http.ResponseWriter, req *http.Request) {
	foo := Foo{}
	if err := req.ParseForm(); err != nil {
		log.Println("Got error parsing form: ", err)
	}
	decoder := schema.NewDecoder()
	if err := decoder.Decode(&foo, req.PostForm); err != nil {
		log.Println("Got error decoding form: ", err)
	}

	index.Execute(resp, Page{
		"New Foo " + foo.Bar,
		"Can only view this page by making a POST or PUT request here",
	})
}

func APIHandler(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	someId := vars["idOrSomething"]
	encoder := json.NewEncoder(resp)
	encoder.Encode(struct{ SomeId string }{someId})
}

type Phone struct{ Kind, Number string }

type Contact struct {
	Name  string
	Phone []Phone
} //

func NewConctactHandler(resp http.ResponseWriter, req *http.Request) {
	contact := Contact{}
	if err := req.ParseForm(); err != nil {
		log.Println("Got error parsing form: ", err)
	}
	decoder := schema.NewDecoder()
	if err := decoder.Decode(&contact, req.PostForm); err != nil {
		log.Println("Got error decoding form: ", err)
	}

	body := "With phone numbers"
	for _, phone := range contact.Phone {
		body += fmt.Sprintf(", %s (%s)", phone.Number, phone.Kind)
	}

	index.Execute(resp, Page{"You made a new contact " + contact.Name, body})
}
