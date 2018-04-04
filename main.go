package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var persons []Person

func addPersonHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		// get the json string sent via POST
		receivedJSON := r.FormValue("person")

		// marshal the received JSON object into a Person struct
		person := Person{}
		json.Unmarshal([]byte(receivedJSON), &person)
		persons = append(persons, person)

		// Print the list of received Person objects
		jsonData, _ := json.Marshal(persons)
		fmt.Println("hello", string(jsonData))
	default:
		fmt.Fprintf(os.Stdout, "Sorry, only GET and POST methods are supported.")
	}
}

func saveDBHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
	default:
		fmt.Fprintf(os.Stdout, "Sorry, only POST methods are supported.")
	}
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.HandleFunc("/addPerson", addPersonHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
