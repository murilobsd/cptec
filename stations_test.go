package cptec

import "testing"

// TestHashStation check hash is equal
func TestHashStation(t *testing.T) {
	testOne := &Station{
		ID:        "12345",
		Locality:  "Jaboticabal",
		UF:        "SP",
		Latitude:  "",
		Longitude: "",
	}
	compare := &Station{
		ID:        "12345",
		Locality:  "Jaboticabal",
		UF:        "SP",
		Latitude:  "",
		Longitude: "",
	}

	if testOne.Hash() != compare.Hash() {
		t.Error("Different hash station")
	}
}
