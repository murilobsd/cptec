package cptec

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
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

// New ...
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
		UserAgent: "Go-CPTEC-0.1",
	}
	c.Station = &StationService{client: c}

	return c
}

func (c *CPTEC) newRequest(method, path string, data interface{}, header interface{}) (req *http.Request, err error) {
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)

	if method == "POST" && data != nil {
		formData, ok := data.(map[string]string)
		if !ok {
			return nil, errors.New("Invalid format type data")
		}
		form := url.Values{}
		for k, v := range formData {
			form.Add(k, v)
		}
		req, err = http.NewRequest(method, u.String(), strings.NewReader(form.Encode()))
		if err != nil {
			return nil, err
		}
		// req.PostForm = form
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	} else {
		req, err = http.NewRequest(method, u.String(), nil)
		if err != nil {
			return nil, err
		}
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
