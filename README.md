# Go CPTEC

[![GoDoc](https://godoc.org/github.com/murilobsd/cptec?status.svg)](https://godoc.org/github.com/murilobsd/cptec) [![Go Report Card](https://goreportcard.com/badge/github.com/murilobsd/cptec)](https://goreportcard.com/report/github.com/murilobsd/cptec) [![Coverage Status](https://coveralls.io/repos/github/murilobsd/cptec/badge.svg?branch=master)](https://coveralls.io/github/murilobsd/cptec?branch=master)


<p align="center">
  <img src="https://i.imgur.com/f8UohJw.png"/ width=200>
</p>


This Go package includes a set of tools for collecting weather data from [INPE/CPTEC](https://www.cptec.inpe.br/).

## Installation

```markdown
go get github.com/murilobsd/cptec
```

## Usage

Get stations:

```go
cptec := cptec.New(nil)
stations, err := cptec.Station.GetAll()
if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
}

for _, station := range stations {
    fmt.Println(station.String())
}

# Output
# ID: 32549 - UF: TO - Locality: UHE Isamu Ikeda Montante
# ID: 32619 - UF: TO - Locality: Xambioa
```