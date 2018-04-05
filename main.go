package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

////////////////////////////////////////////////////////////////////////////////

// Persons The list of persons objects in memory.
var Persons []Person

// PersonsDBName Name of the database file used to store persons objects
const PersonsDBName = "persons.json"

////////////////////////////////////////////////////////////////////////////////

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
		Persons = append(Persons, person)
	default:
		fmt.Fprintf(os.Stdout, "Sorry, only GET and POST methods are supported.")
	}
}

func saveDBHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		fmt.Print("info: saving persons...")
		if f, err := os.Create(PersonsDBName); err != nil {
			// couldn't open file
			panic(err)
		} else {
			// write persons to file
			if b, err := json.Marshal(Persons); err != nil {
				// error while marshaling persons into JSON
				panic(err)
			} else {
				defer f.Close()
				if _, err := f.Write(b); err != nil {
					panic(err)
				} else {
					fmt.Println(" ok")
				}
			}
		}
	default:
		fmt.Fprintf(os.Stdout, "Sorry, only POST methods are supported.")
	}
}

////////////////////////////////////////////////////////////////////////////////

func main() {
	// Open persons db and load its content
	if dat, err := ioutil.ReadFile(PersonsDBName); err != nil {
		// problem loading persons. skip this and start with an empty db
		fmt.Println(err)
	} else {
		if err := json.Unmarshal(dat, &Persons); err != nil {
			// problem unmarshalling persons. again, keep going with an empty db
			fmt.Println(err)
		}
	}

	// Initialize the server
	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.HandleFunc("/addPerson", addPersonHandler)
	http.HandleFunc("/savePersons", saveDBHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
