// Copyright 2013 The go-cptec AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cptec

import (
	"encoding/xml"
	"fmt"
)

type StationService service

// Station ...
type Station struct {
	XMLName          xml.Name `xml:"metar"`
	Codigo           string   `json:"nome" xml:"nome"`
	Atualizacao      string   `json:"atualizacao" xml:"atualizacao"`
	Pressao          float64  `json:"pressao" xml:"pressao"`
	Temperatura      float64  `json:"temperatura" xml:"temperatura"`
	Tempo            string   `json:"tempo" xml:"tempo"`
	TempoDescricao   string   `json:"tempo_desc" xml:"tempo_desc"`
	Umidade          float64  `json:"umidade" xml:"umidade"`
	VentoDirecao     float64  `json:"vento_dir" xml:"vento_dir"`
	VentoIntensidade float64  `json:"vento_int" xml:"vento_int"`
	Visibilidade     string   `json:"visibilidade" xml:"visibilidade"`
}

func (s *StationService) Get(code string) (*Station, *Response, error) {
	u := fmt.Sprintf("estacao/%s/condicoesAtuais.xml", code)

	req, err := s.client.NewRequest("GET", u, nil)

	if err != nil {
		fmt.Println(err)
	}

	var station *Station

	resp, err := s.client.Do(req, &station)

	if err != nil {
		return nil, resp, err
	}
	return station, resp, nil
}
