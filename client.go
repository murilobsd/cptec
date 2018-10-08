package cptec

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

// These are the errors that can be returned
var (
	ErrStatusCode = errors.New("Http invalid status code")
)

// CPTEC ...
type CPTEC struct {
	Client    *http.Client
	BaseURL   *url.URL
	UserAgent string
	Station   *StationService
}

func New(httpClient *http.Client) *CPTEC {
	if httpClient == nil {
		cookieJar, _ := cookiejar.New(nil)
		httpClient = http.DefaultClient
		httpClient.Jar = cookieJar
		httpClient.Timeout = time.Duration(10 * time.Second)
	}

	url, _ := url.Parse("http://sinda.crn.inpe.br")

	c := &CPTEC{
		Client:    httpClient,
		BaseURL:   url,
		UserAgent: "Go-INPE-0.1",
	}
	c.Station = &StationService{client: c}

	return c
}

func (c *CPTEC) newRequest(method, path string) (req *http.Request, err error) {
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)

	req, err = http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
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
	return []byte(""), ErrStatusCode

}
