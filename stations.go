package cptec

import (
	"fmt"
	"regexp"
)

// These are the constants that can be use
const (
	// pathStations link to get the list of stations.
	pathStations = "/PCD/SITE/novo/site/historico/index.php"
)

// These are the constants that can be use
var (
	rgStations = `<option value=(?P<ID>\d+)>\d+-(?P<UF>[A-Z]{2})-(?P<Locality>.*?)</option>`
)

// StationService ...
type StationService struct {
	client *CPTEC
}

// Station structure representing the CPTEC/INPE weather station.
//
// Information such as location and identification is part of this
// structure.
type Station struct {
	ID       string `json:"id"`
	UF       string `json:"uf"`
	Locality string `json:"locality"`
}

func (s Station) String() string {
	return fmt.Sprintf("ID: %s - UF: %s - Locality: %s", s.ID, s.UF, s.Locality)
}

// GetAll ...
func (s *StationService) GetAll() ([]*Station, error) {
	var stations []*Station
	req, err := s.client.newRequest("GET", pathStations)

	if err != nil {
		return nil, err
	}
	html, err := s.client.do(req)

	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile(rgStations)

	for _, match := range re.FindAllStringSubmatch(string(html), -1) {
		station := &Station{
			ID:       match[1],
			UF:       match[2],
			Locality: match[3],
		}
		stations = append(stations, station)
	}
	return stations, err
}
