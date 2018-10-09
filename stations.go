package cptec

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
)

// These are the constants that can be use
const (
	// pathStations link to get the list of stations.
	pathStations = "/PCD/SITE/novo/site/historico/index.php"
	// pathFullStations link to get full information of stations.
	pathFullStations = "/PCD/SITE/novo/site/historico/passo2.php"
)

// These are the constants that can be use
var (
	rgStations     = `<option value=(?P<ID>\d+)>\d+-(?P<UF>[A-Z]{2})-(?P<Locality>.*?)</option>`
	rgFullStations = `Latitude:\s+(?P<Latitude>.*?),\s+?Longitude:\s+(?P<Longitude>.*?),`
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
	ID        string `json:"id"`
	UF        string `json:"uf"`
	Locality  string `json:"locality"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

// Stations array ...
type Stations []*Station

// ToBytes convert station to byte
func (s *Station) ToBytes() []byte {
	jsonBytes, _ := json.Marshal(s)
	return jsonBytes
}

// Hash return md5 of station
func (s *Station) Hash() string {
	stationBytes := s.ToBytes()
	stationHash := fmt.Sprintf("%x", md5.Sum(stationBytes))
	return stationHash
}

func (s *Station) String() string {
	return fmt.Sprintf("ID: %s UF: %s Locality: %s Latitude: %s Longitude: %s", s.ID, s.UF, s.Locality, s.Latitude, s.Longitude)
}

// GetAll ...
func (s *StationService) GetAll() ([]*Station, error) {
	var stations []*Station
	req, err := s.client.NewRequest("GET", pathStations, nil, nil)

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

// Get coleta todos
func (s *StationService) Get(station *Station) error {
	form := make(map[string]string)
	form["id"] = station.ID
	req, err := s.client.NewRequest("POST", pathFullStations, form, nil)
	if err != nil {
		return err
	}

	html, err := s.client.do(req)

	if err != nil {
		return err
	}

	re := regexp.MustCompile(rgFullStations)

	for _, match := range re.FindAllStringSubmatch(string(html), -1) {
		station.Latitude = match[1]
		station.Longitude = match[2]
	}
	return nil
}

func worker(s *StationService, id int, jobs <-chan *Station, results chan<- *Station) {
	for j := range jobs {
		s.Get(j)
		results <- j
	}
}

// GetFull ...
func (s *StationService) GetFull() ([]*Station, error) {
	stations, err := s.GetAll()
	var stationsFuture []*Station
	if err != nil {
		return nil, err
	}

	fmt.Println("Número de Estações: ", len(stations))

	jobs := make(chan *Station, len(stations))
	results := make(chan *Station, len(stations))

	for w := 1; w <= 5; w++ {
		go worker(s, w, jobs, results)
	}

	for _, station := range stations {
		jobs <- station
	}
	close(jobs)

	for range stations {
		station := <-results
		stationsFuture = append(stationsFuture, station)
	}

	if err := Save("./data/stations.json", stationsFuture); err != nil {
		log.Fatalln(err)
	}

	return stationsFuture, nil
}
