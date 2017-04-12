package main

import (
	"fmt"
	"net/http"
	"html/template"
	"log"
	"strings"
	"encoding/json"
	//"golang.org/x/crypto/ssh/test"
)

type Context struct {
	FirstName string
	LastName string
}

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("val: ", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello Scott!")
}

func login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tmpl, err := template.ParseFiles("/home/shendrickson/GO/src/test/templates/test.html")
		if err == nil {
			tmpl.Execute(w, nil)
		} else {
			fmt.Println(err)
			json.NewEncoder(w).Encode("Nope")
		}
	default:
		r.ParseForm()
		fname := strings.Join(r.Form["firstName"], "")
		lname := strings.Join(r.Form["lastName"], "")
		tmpl, err := template.ParseFiles("/home/shendrickson/GO/src/test/templates/test.html")
		if err == nil {
			context := Context{fname, lname}
			tmpl.Execute(w, context)
		} else {
			fmt.Println(err)
			json.NewEncoder(w).Encode("Nope")
		}
	}
}

func main() {
	myMux := http.NewServeMux()
	myMux.HandleFunc("/", sayHelloName)
	myMux.HandleFunc("/login", login)
	myMux.Handle("./home/shendrickson/GO/src/test/templates/", http.StripPrefix("./home/shendrickson/GO/src/test/templates/", http.FileServer(http.Dir("templates"))))
	err := http.ListenAndServe(":8080", myMux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}