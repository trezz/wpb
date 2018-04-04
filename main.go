package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
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

		locations := make(map[string]map[string]string)

		// Load received values
		for key, values := range r.Form {
			for _, value := range values {
				if key == "name" {
					p.Name = value
				} else if key == "desc" {
					p.Description = value
				} else if strings.HasPrefix(key, "locations") {
					items := strings.Split(key, "]") // 0: index, 1: field
					// create map if not present
					if _, ok := locations[items[0]]; !ok {
						locations[items[0]] = make(map[string]string)
					}
					// set the field's value
					locations[items[0]][items[1]] = value
				} else {
					fmt.Println("parse error:", key, value)
				}
			}
		}

		// Add locations to the person object
		for _, fields := range locations {
			loc := DatedLocation{}
			loc.Location = LatLng{}
			for key, value := range fields {
				if strings.Contains(key, "latitude") {
					loc.Location.Latitude, _ = strconv.ParseFloat(value, 64)
				} else if strings.Contains(key, "longitude") {
					loc.Location.Longitude, _ = strconv.ParseFloat(value, 64)
				} else if strings.Contains(key, "date") {
					layout := "01/02/2006"
					loc.Time, _ = time.Parse(layout, value)
				} else if strings.Contains(key, "refs") {
					loc.Refs = value
				} else {
					fmt.Println("load error:", key, value)
				}
			}
			p.Locations = append(p.Locations, loc)
		}

		// Print persons
		persons = append(persons, p)

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
