// Copyright 2017 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The simple command demonstrates a simple functionality which
// prompts the user for a GitHub username and lists all the public
// organization memberships of the specified username.
package main

import (
	"fmt"

	"github.com/murilobsd/cptec"
)

// FetchStations fetch all stations
func FetchStations() ([]*cptec.Station, error) {
	cptec := cptec.New(nil)
	stations, err := cptec.Station.GetAll()
	return stations, err
}

func main() {
	stations, err := FetchStations()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for _, station := range stations {
		fmt.Println(station.String())
	}
}
