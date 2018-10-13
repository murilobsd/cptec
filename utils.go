package cptec

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var lock sync.Mutex

// Marshal is a function that marshals the object into an
// io.Reader.
// By default, it uses the JSON marshaller.
var Marshal = func(v interface{}) (io.Reader, error) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

// Save saves a representation of v to the file at path.
func Save(path string, v interface{}) error {
	lock.Lock()
	defer lock.Unlock()
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	r, err := Marshal(v)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, r)
	return err
}

// Unmarshal is a function that unmarshals the data from the
// reader into the specified value.
// By default, it uses the JSON unmarshaller.
var Unmarshal = func(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

// Load loads the file at path into v.
// Use os.IsNotExist() to see if the returned error is due
// to the file being missing.
func Load(path string, v interface{}) error {
	lock.Lock()
	defer lock.Unlock()
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return Unmarshal(f, v)
}

// Trim space
func Trim(value string) string {
	return strings.TrimSpace(value)
}

// FixCityState ajusta o nome da cidade e do estado
func FixCityState(cityValue string) (string, string) {
	cityClean := CleanString(cityValue)
	splitCity := strings.Split(cityClean, "/")
	if len(splitCity) != 2 {
		log.Fatalf("split city name error: %s", cityValue)
	}
	return Trim(splitCity[0]), Trim(splitCity[1])
}

// ParserDate parsea o dado de data
func ParserDate(date string) (*time.Time, error) {
	layout := "02/01/2006"
	t, err := time.Parse(layout, date)
	if err != nil {
		return nil, ErrParseDate
	}
	return &t, nil
}

// ParserTime parsea o dade de útil para o horário do nascer
// do sol e o por do sol
func ParserTime(timestr string) (*time.Time, error) {
	timestr = Trim(timestr)
	layout := "15:04"
	t, err := time.Parse(layout, timestr)
	if err != nil {
		return nil, nil
	}
	return &t, nil
}

// PeriodToTime convert period to time
func PeriodToTime(period string) time.Time {
	if period == "Manha" {
		return time.Time{}
	} else if period == "Tarde" {
		return time.Time{}
	} else {
		return time.Time{}
	}
}

// CleanString Remove no caracteres
func CleanString(value string) string {
	re := regexp.MustCompile(`[|°%]+`)
	s := re.ReplaceAllLiteralString(value, "")
	return s
}

//FindLat pesquisa a latitude
func FindLat(value string) {}

//FindLon pesquisa a longitude
func FindLon(value string) {}

//ToFloat converte string para float64, caso retorne erro o valor -999.0 é retornado.
func ToFloat(number string) float64 {
	number = Trim(CleanString(number))
	n, err := strconv.ParseFloat(number, 64)
	if err != nil {
		return float64(-999.0)
	}
	return n
}

//ToInt converte string para int64, caso retorne erro o valor -999 é retornado.
func ToInt(number string) int64 {
	number = Trim(CleanString(number))
	n, err := strconv.ParseInt(number, 10, 64)
	if err != nil {
		return int64(-999)
	}
	return n
}
