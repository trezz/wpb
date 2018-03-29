package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	p := Person{
		Name:        "Vincent Camus",
		Description: "C'est moi",
		Picture:     "https://ih0.redbubble.net/image.394665545.8115/flat,800x800,075,f.jpg"}
	jsonData, _ := json.Marshal(p)
	fmt.Println(string(jsonData))
}
