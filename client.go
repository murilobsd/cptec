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

const (
	defaultBaseURL = "http://sinda.crn.inpe.br"
	userAgent      = "go-cptec"

	// Headers

	// headerAccept header Accept
	headerAccept = "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"

	// headerContentTypePost header content type when post data
	headerContentTypePost = "application/x-www-form-urlencoded"
)

// These are the errors that can be returned
var (
	ErrStatusCode = errors.New("Http invalid status code")
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
	Station *StationService
}

// New returns a new CPTEC client. If a nil httpClient is
// provided, http.DefaultClient will be used.
func New(httpClient *http.Client) *CPTEC {
	if httpClient == nil {
		cookieJar, _ := cookiejar.New(nil)
		httpClient = http.DefaultClient
		httpClient.Jar = cookieJar
		httpClient.Timeout = time.Duration(10 * time.Second)
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &CPTEC{
		Client:    httpClient,
		BaseURL:   baseURL,
		UserAgent: userAgent,
	}

	c.Station = &StationService{client: c}

	return c
}

// NewRequest ...
func (c *CPTEC) NewRequest(method, path string, data interface{}, header interface{}) (req *http.Request, err error) {
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
		req.Header.Add("Content-Type", headerContentTypePost)
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

	req.Header.Set("Accept", headerAccept)
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
