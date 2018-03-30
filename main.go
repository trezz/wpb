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
		// Create a new Person object
		p := Person{
			Name:        r.FormValue("PersonName"),
			Description: r.FormValue("PersonDescription"),
		}
		persons = append(persons, p)
		for _, p := range persons {
			jsonData, _ := json.Marshal(p)
			fmt.Println(string(jsonData))
		}
	default:
		fmt.Fprintf(os.Stdout, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.HandleFunc("/addPerson", addPersonHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
