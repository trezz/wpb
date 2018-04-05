package main

import (
	"time"
)

// DatedLocation A geographic location associated to a date. Used to locate a
// person in time.
type DatedLocation struct {
	Time      time.Time `json:"time,omitempty,string"`
	Latitude  float64   `json:"latitude,omitempty,string"`
	Longitude float64   `json:"longitude,omitempty,string"`
	Refs      string    `json:"refs,omitempty"`
}

// Person A person with registered locations over its life
type Person struct {
	Name        string          `json:"name,omitempty"`
	Description string          `json:"desc,omitempty"`
	Picture     string          `json:"picture,omitempty"`
	Locations   []DatedLocation `json:"locations,omitempty"`
}
