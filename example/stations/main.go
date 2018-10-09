// Copyright 2018 Murilo Ijanc. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/murilobsd/cptec"
)

// FetchStations fetch all stations
func FetchStations() ([]*cptec.Station, error) {
	cptec := cptec.New(nil)
	stations, err := cptec.Station.GetAll()
	for _, station := range stations {
		cptec.Station.Get(station)
		fmt.Println(station.String())
	}
	return stations, err
}

func main() {
	cptec := cptec.New(nil)
	cptec.Station.GetFull()
	// FetchStations()
	// stations, err := FetchStations()
	// if err != nil {
	// 	fmt.Printf("Error: %v\n", err)
	// 	return
	// }
}
