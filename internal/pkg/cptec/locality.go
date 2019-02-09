// Copyright 2013 The go-cptec AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cptec

import (
	"encoding/xml"
	"fmt"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type LocalityService service

// Locality ...
type Locality struct {
	XMLName xml.Name `xml:"cidade"`
	Nome    string   `json:"nome" xml:"nome"`
	UF      string   `json:"uf" xml:"uf"`
	ID      int      `json:"id" xml:"id"`
}

type OptionsLocality struct {
	City string `url:"city"`
}

type response struct {
	XMLName    xml.Name   `xml:"cidades"`
	Localities []Locality `xml:"cidade"`
}

func (l *Locality) String() string {
	return fmt.Sprintf("Nome: %s UF: %s ID: %d", l.Nome, l.UF, l.ID)
}

func (s *LocalityService) Get(cityName string) (*response, *Response, error) {
	u := "listaCidades"
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)

	cleanCityName, _, _ := transform.String(t, cityName)

	opt := &OptionsLocality{
		City: cleanCityName,
	}

	u, err := addQuery(u, opt)

	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)

	if err != nil {
		fmt.Println(err)
	}

	var loc *response

	resp, err := s.client.Do(req, &loc)

	if err != nil {
		return nil, resp, err
	}
	return loc, resp, nil
}
