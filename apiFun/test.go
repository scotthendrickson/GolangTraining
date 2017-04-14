package main

import (
	"net/http"
	"text/template"
	"encoding/json"
	"fmt"
)

type Person struct {
	ID        string `json:"id,omitempty"`
	Firstname string `json:"Firstname,omitempty"`
	Lastname  string `json:"Lastname,omitempty"`
	Address   *Address `json:"Lastname,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func PersonEndpoint(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	switch req.Method {
	case "GET":
		params := req.URL.Query().Get("id")
		for _, item := range people {
			if item.ID == params {
				json.NewEncoder(w).Encode(item)
				return
			}
		}
		json.NewEncoder(w).Encode(&Person{})
	case "POST":
		var person Person
		_ = json.NewDecoder(req.Body).Decode(&person)
		people = append(people, person)
		json.NewEncoder(w).Encode(people)
	case "PUT":
		var person Person
		_ = json.NewDecoder(req.Body).Decode(&person)
		for index, item := range people {
			if item.ID == person.ID {
				if people[index].Firstname !=  person.Firstname {
					people[index].Firstname = person.Firstname
				}
				if people[index].Lastname !=  person.Lastname {
					people[index].Lastname = person.Lastname
				}
				if people[index].Address !=  person.Address {
					people[index].Address = person.Address
				}
			}
		}
		json.NewEncoder(w).Encode(people)
	case "DELETE":
		params := req.URL.Query().Get("id")
		for index, item := range people {
			if item.ID == params {
				people = append(people[:index], people[index + 1:]...)
			}
		}
		json.NewEncoder(w).Encode(people)
	default:
		json.NewEncoder(w).Encode("No one found")
	}

}

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func main() {
	myMux := http.NewServeMux()
	people = append(people, Person{ID: "1", Firstname: "Scott", Lastname: "Hendrickson", Address: &Address{City: "Spanish Fork", State: "Utah"}})
	people = append(people, Person{ID: "2", Firstname: "Megan", Lastname: "Hendrickson", Address: &Address{City: "Spanish Fork", State: "Utah"}})
	people = append(people, Person{ID: "3", Firstname: "Thomas", Lastname: "Hendrickson", Address: &Address{City: "Spanish Fork", State: "Utah"}})
	myMux.HandleFunc("/", myHandlerFunc)
	myMux.HandleFunc("/test", myTestFunc)
	myMux.HandleFunc("/people", GetPeopleEndpoint)
	myMux.HandleFunc("/person", PersonEndpoint)
	http.ListenAndServe(":8080", myMux)
}

func myHandlerFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content Type", "text/html")
	tmpl, err := template.New("anyNameForTemplate").Parse(doc)
	if err == nil {
		tmpl.Execute(w, nil)
	}
}

func myTestFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content Type", "text/html")
	tmpl, err := template.New("anyNameForTemplate").Parse(docEr)
	if err == nil {
		tmpl.Execute(w, nil)
	}
}

const doc = `
<!DOCTYPE html>
<html>
<head lang="en">
	<meta charset="UTF-8">
	<title>First Template</title>
</head>
<body>
	<h1 style="color:blue">Hello Jamaica</h1>
	<input type="text" placeholder="{{">
</body>
</html>
`

const docEr = `
<!DOCTYPE html>
<html>
<head lang="en">
	<meta charset="UTF-8">
	<title>First Template</title>
</head>
<body>
	<h1 style="color:blue">Hello Derrick</h1>
	<input type="text" placeholder="Stuff">
</body>
</html>
`
//const docPath = `/home/shendrickson/GO/src/test/templates/test.html`