package client

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	defaultBaseURL = "http://servicos.cptec.inpe.br"
	userAgent      = "go-cptec"
)

// These are the errors that can be returned
var (
	errStatusCode    = errors.New("Http invalid status code")
	errInvalidHeader = errors.New("Invalid header")
)

// CPTEC manages communication with the CPTEC.
type CPTEC struct {
	Client *http.Client // HTTP client used to communicate with the API.

	// Base URL for CPTEC requests. BaseURL should always be specified with a
	// trailing slash.
	BaseURL *url.URL

	// User agent used when communicating with the CPTEC.
	UserAgent string

	// Services used for talking to different parts of the CPTEC.
	// Station  *StationService
	// Forecast *ForecastService
}

// New returns a new CPTEC client. If a nil httpClient is
// provided, http.DefaultClient will be used.
func New(httpClient *http.Client) *CPTEC {
	if httpClient == nil {
		httpClient = http.DefaultClient
		httpClient.Timeout = time.Duration(10 * time.Second)
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &CPTEC{
		Client:    httpClient,
		BaseURL:   baseURL,
		UserAgent: userAgent,
	}

	// c.Station = &StationService{client: c}
	// c.Forecast = &ForecastService{client: c}

	return c
}

// NewRequest ...
func (c *CPTEC) NewRequest(method, path string, data interface{}, header interface{}) (req *http.Request, err error) {
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)

	req, err = http.NewRequest(method, u.String(), nil)

	if err != nil {
		return nil, err
	}

	if header != nil {
		headerData, ok := header.(map[string]string)
		if !ok {
			return nil, errors.New("Invalid header")
		}
		for k, v := range headerData {
			req.Header.Add(k, v)
		}
	}

	req.Header.Set("User-Agent", c.UserAgent)

	return req, nil
}

func (c *CPTEC) do(req *http.Request) ([]byte, error) {
	resp, err := c.Client.Do(req)

	if err != nil {
		return []byte(""), err
	}

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return []byte(""), err
		}

		defer resp.Body.Close()

		return body, nil
	}
	return nil, errStatusCode

}
