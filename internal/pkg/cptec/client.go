// Copyright 2013 The go-cptec AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cptec

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/google/go-querystring/query"
	"golang.org/x/net/html/charset"
)

const (
	defaultBaseURL = "http://servicos.cptec.inpe.br/XML/"
	userAgent      = "go-cptec"
)

var (
	errTraling = errors.New("BaseURL must have a trailing slash")
)

// Client manages communication
type Client struct {
	client    *http.Client // HTTP client used to communicate
	BaseURL   *url.URL
	UserAgent string

	commom service

	Locality *LocalityService
	Station  *StationService
}

type service struct {
	client *Client
}

// addQuery add paramenters to url
func addQuery(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}
	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}
	qs, err := query.Values(opt)

	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

// NewClient returns new Cptec client
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(defaultBaseURL)
	c := &Client{
		client:    httpClient,
		BaseURL:   baseURL,
		UserAgent: userAgent,
	}
	c.commom.client = c
	c.Locality = &LocalityService{client: c}
	c.Station = &StationService{client: c}
	return c
}

// NewRequest TODO
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, errTraling
	}

	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/xml")
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}

// Response TODO
type Response struct {
	*http.Response
}

func newResponse(resp *http.Response) *Response {
	response := &Response{Response: resp}
	return response
}

// ResponseError TODO
type ResponseError struct {
	Response *http.Response
}

func (rr *ResponseError) Error() string {
	return fmt.Sprintf("%v %v: %v", rr.Response.Request.Method, rr.Response.Request.URL, rr.Response.StatusCode)
}

func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	response := newResponse(resp)
	err = CheckResponse(resp)
	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			decoder := xml.NewDecoder(resp.Body)
			decoder.CharsetReader = charset.NewReaderLabel
			decErr := decoder.Decode(v)
			if decErr == io.EOF {
				decErr = nil
			}
			if decErr != nil {
				err = decErr
			}
		}
	}
	return response, err
}

func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	// errorResponse := &ResponseError{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	fmt.Println(data)
	// if err == nil && data != nil {
	// 	json.Unmarshal(data, errorResponse)
	// }
	return nil
}
