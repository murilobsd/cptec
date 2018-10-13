package cptec

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// These are the constants that can be use
const (
	// Selectors to get data

	// City and state Ex. Jaboticabal / SP |
	cityState = ".pt-1 > strong:nth-child(1)"
	// Date getting data Ex. 10/09/2018
	date = "div.previsao:nth-child(1) > div:nth-child(1) > div:nth-child(1) > small:nth-child(3)"
	// Period Ex. Manhã, Tarde, Noite
	period = "div.previsao:nth-child(1) > div:nth-child(2) > div:nth-child(3) > span:nth-child(1)"

	// Forecast Selectors

	// Min Temperature  Ex. 15°
	forecastTempMin = "div.proximos-dias:nth-child(%s) > div:nth-child(2) > div:nth-child(1) > div:nth-child(1) > span:nth-child(1)"
	// Max Temperature Ex. 26°
	forecastTempMax = "div.proximos-dias:nth-child(%s) > div:nth-child(2) > div:nth-child(2) > div:nth-child(1) > span:nth-child(1)"
	// UV Index Ex. 11
	forecastIndexUV = "div.proximos-dias:nth-child(%s) > div:nth-child(3) > div:nth-child(1) > div:nth-child(1) > a:nth-child(1) > span:nth-child(1)"
	// Probability Rain Ex. 80%
	forecastRain = "div.proximos-dias:nth-child(%s) > div:nth-child(3) > div:nth-child(2) > div:nth-child(1) > span:nth-child(1)"
	// Date getting data Ex. 10/09/2018
	forecastDate = "div.proximos-dias:nth-child(%s) > div:nth-child(1) > div:nth-child(1) > small:nth-child(2)"
	// Sunrise Ex. 05:46
	forecastSunrise = "div.proximos-dias:nth-child(%s) > div:nth-child(4) > div:nth-child(1) > span:nth-child(2)"
	// Sunset Ex. 18:14
	forecastSunset = "div.proximos-dias:nth-child(%s) > div:nth-child(4) > div:nth-child(2) > span:nth-child(2)"
	// Link Icon description climate
	forecastPicture = "div.proximos-dias:nth-child(%s) > div:nth-child(1) > div:nth-child(1) > img:nth-child(3)"
	// Description ...
	forecastDescription = "div.proximos-dias:nth-child(%s) > div:nth-child(1) > div:nth-child(1) > img:nth-child(3)"
)

var (
	// Errors

	// ErrParseDate error parse date
	ErrParseDate = errors.New("Invalid parse date")
)

// ForecastService serviço que fornece o cliente para informações
type ForecastService struct {
	client *CPTEC
}

// Forecast estrutura do dado retornado da previsão
type Forecast struct {
	Date            *time.Time `json:"date"`
	Description     string     `json:"description"`
	PictureURL      string     `json:"picture"`
	TempMin         float64    `json:"temp_min"`
	TempMax         float64    `json:"temp_max"`
	Sunrise         *time.Time `json:"sunrise"`
	Sunset          *time.Time `json:"sunset"`
	UvIndex         int64      `json:"uv_index"`
	RainProbability float64    `json:"rain_probability"`
}

func (fo *Forecast) String() string {
	return fmt.Sprintf("TempMin: %f TempMax: %f Description: %s Picture: %s Sunrise: %v Sunset: %v UV: %d RainPro: %f", fo.TempMin, fo.TempMax, fo.Description, fo.PictureURL, fo.Sunrise, fo.Sunset, fo.UvIndex, fo.RainProbability)
}

// ForecastCity representa a previsão da cidade
type ForecastCity struct {
	City      string      `json:"city"`
	UF        string      `json:"uf"`
	Forecasts []*Forecast `json:"forecasts"`
}

// Get get weather forecast of city.
func (f *ForecastService) Get(cityName string) {
	u := "https://tempo.cptec.inpe.br/" + cityName

	resp, err := http.Get(u)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", resp.StatusCode, resp.Status)
	}
	//html, err := os.Open("testdata/forecast.html")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer html.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	defer resp.Body.Close()
	//doc, err := goquery.NewDocumentFromReader(html)

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(doc.Find(cityState).Contents().Text())
	// fmt.Println(doc.Find(date).Contents().Text())
	// fmt.Println(doc.Find(period).Contents().Text())
	// fmt.Println(doc.Find(icon).Attr("src"))
	// fmt.Println(doc.Find(tempMax).Contents().Text())
	// fmt.Println(doc.Find(tempMin).Contents().Text())
	// fmt.Println(doc.Find(rainPro).Contents().Text())
	// fmt.Println(doc.Find(uvIndex).Contents().Text())
	// fmt.Println(doc.Find(sunRise).Contents().Text())
	// fmt.Println(doc.Find(sunSet).Contents().Text())

	e, _ := f.extractData(doc, 5)
	fmt.Println(e)
}

// extractData extrai dados de previsao do html
func (f *ForecastService) extractData(doc *goquery.Document, limit int) (*ForecastCity, error) {
	var forecasts []*Forecast
	// forecastDays limit 5
	forecastDays := []string{"1", "2", "3", "4", "5"}
	if limit <= 0 && limit > 5 {
		limit = 5
	}

	city, state := FixCityState(doc.Find(cityState).Text())
	for num, day := range forecastDays {
		if num+1 < limit+1 {
			tempmin := doc.Find(fmt.Sprintf(forecastTempMin, day)).Text()
			tempmax := doc.Find(fmt.Sprintf(forecastTempMax, day)).Text()
			description, _ := doc.Find(fmt.Sprintf(forecastDescription, day)).Attr("title")
			icon, _ := doc.Find(fmt.Sprintf(forecastPicture, day)).Attr("src")
			sunrise, _ := ParserTime(doc.Find(fmt.Sprintf(forecastSunrise, day)).Text())
			sunset, _ := ParserTime(doc.Find(fmt.Sprintf(forecastSunset, day)).Text())
			uv := ToInt(doc.Find(fmt.Sprintf(forecastIndexUV, day)).Text())
			rain := ToFloat(doc.Find(fmt.Sprintf(forecastRain, day)).Text())
			forecast := &Forecast{
				TempMin:         ToFloat(tempmin),
				TempMax:         ToFloat(tempmax),
				Description:     description,
				PictureURL:      icon,
				Sunrise:         sunrise,
				Sunset:          sunset,
				UvIndex:         uv,
				RainProbability: rain,
			}
			forecasts = append(forecasts, forecast)
		}
	}
	forecastCity := &ForecastCity{
		City:      city,
		UF:        state,
		Forecasts: forecasts,
	}
	return forecastCity, nil
}
