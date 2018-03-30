package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var persons []Person

func rootHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "web/index.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		// Create a new Person object
		p := Person{
			Name:        r.FormValue("PersonName"),
			Description: r.FormValue("PersonDescription")}
		persons = append(persons, p)
		for _, p := range persons {
			jsonData, _ := json.Marshal(p)
			fmt.Println(string(jsonData))
		}

		http.ServeFile(w, r, "web/index.html")
	default:
		fmt.Fprintf(os.Stdout, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	/*
		p := Person{
			Name:        "Vincent Camus",
			Description: "C'est moi",
			Picture:     "https://ih0.redbubble.net/image.394665545.8115/flat,800x800,075,f.jpg"}
		jsonData, _ := json.Marshal(p)
		fmt.Println(string(jsonData))
	*/

	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
