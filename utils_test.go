package cptec

import (
	"testing"
	"time"
)

// TestFixCityName correct city and state name
func TestFixCityName(t *testing.T) {
	cityName := "Ribeirão Preto / SP |"
	city, state := FixCityState(cityName)
	if city != "Ribeirão Preto" {
		t.Errorf("city name no match: %s expected: %s", city, "Ribeirão Preto")
	}
	if state != "SP" {
		t.Errorf("state name no match: %s expected: %s", state, "SP")
	}
}

// TestParseDate testing pase date
func TestParseDate(t *testing.T) {
	t1 := time.Date(2018, time.October, 9, 0, 0, 0, 0, time.UTC)
	t2, _ := ParserDate("09/10/2018")
	if t1.Day() != t2.Day() {
		t.Errorf("parse date: %s expected: %s", t1.String(), t2.String())
	}
}

// TestParseTime testing parse time
func TestParseTime(t *testing.T) {
	t1 := "05:44"
	texp, _ := ParserTime(t1)
	if texp.Hour() != 5 {
		t.Errorf("parse time hour: %s expected: %d", t1, texp.Hour())
	}
	if texp.Minute() != 44 {
		t.Errorf("parse time minute: %s expected: %d", t1, texp.Minute())
	}
}

// TestCleanString testing clean string
func TestCleanString(t *testing.T) {
	city := "Ribeirão Preto / SP | "
	exp := "Ribeirão Preto / SP  "
	clean := CleanString(city)
	if clean != exp {
		t.Errorf("clean string: %s expected: %s", clean, exp)
	}
}
