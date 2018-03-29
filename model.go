package main

import (
	"fmt"
	"time"
)

// LatLng Latitude/Longitude structure holding a location
type LatLng struct {
	Latitude  float32 `json:"latitude,omitempty"`
	Longitude float32 `json:"longitude,omitempty"`
}

// DatedLocation A geographic location associated to a date. Used to locate a
// person in time.
type DatedLocation struct {
	Time     time.Time `json:"time,omitempty"`
	Location LatLng    `json:"location,omitempty"`
}

// Person A person with registered locations over its life
type Person struct {
	Name        string          `json:"name,omitempty"`
	Description string          `json:"desc,omitempty"`
	Picture     string          `json:"picture,omitempty"`
	Locations   []DatedLocation `json:"locations,omitempty"`
}

func hello() {
	fmt.Println("hello")
}
