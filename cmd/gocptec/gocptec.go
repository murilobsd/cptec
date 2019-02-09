package main

import (
	"fmt"

	cptec "github.com/murilobsd/cptec/internal/pkg/cptec"
)

func main() {
	client := cptec.NewClient(nil)

	// get locations
	locs, _, err := client.Locality.Get("ribeirão preto")
	if err != nil {
		panic(err)
	}
	for _, loc := range locs.Localities {
		fmt.Println(loc.String())
	}
}
