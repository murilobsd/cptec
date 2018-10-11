<p align="center">
  <img src="https://i.imgur.com/f8UohJw.png"/ alt="gocptec" width=150>
</p>

> Weather data for geekers

[![GoDoc](https://godoc.org/github.com/murilobsd/cptec?status.svg)](https://godoc.org/github.com/murilobsd/cptec) [![Go Report Card](https://goreportcard.com/badge/github.com/murilobsd/cptec)](https://goreportcard.com/report/github.com/murilobsd/cptec) [![Coverage Status](https://coveralls.io/repos/github/murilobsd/cptec/badge.svg?branch=master)](https://coveralls.io/github/murilobsd/cptec?branch=master) [![Build Status](https://travis-ci.org/murilobsd/cptec.svg?branch=master)](https://travis-ci.org/murilobsd/cptec) 
[![License](https://img.shields.io/badge/license-BSD-green.svg)](https://raw.githubusercontent.com/murilobsd/cptec/master/LICENSE)

[`Go Cptec`](https://github.com/murilobsd/cptec) is an [INPE/CPTEC](https://www.cptec.inpe.br/) weather data parser. The data range from weather stations, given (temperature, rain, etc.), alerts, forecast.

> The project has no link with [INPE/CPTEC](https://www.cptec.inpe.br/).

# Table of Contents

- [Features](#features)
- [Installing](#installing)
- [Usage](#usage)
- [License](#License)

## Features

- Station list
- Weather data
- Forecast

## Installing

### From source (always latest)

Make sure to have [go](https://golang.org/) (1.9+) installed, then do:

```bash
go get -u github.com/murilobsd/cptec
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
    fmt.Println(station)
}

# Output
# ID: 32549 - UF: TO - Locality: UHE Isamu Ikeda Montante
# ID: 32619 - UF: TO - Locality: Xambioa
```

## License

Released under the [BSD](./LICENSE) license.