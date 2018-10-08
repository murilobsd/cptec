# Go CPTEC

This Go package includes a set of tools for scraping weather data from INPE/CPTEC.

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