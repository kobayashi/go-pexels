package pexels

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

const (
	baseURL = "https://api.pexels.com/v1/search/"
)

// Option is providing APIKEY
type Option struct {
	APIKey string
}

// Queries is Query options
type Queries struct {
	Query   []string
	PerPage int
	Page    int
}

// Client creats http client with Option
type Client struct {
	Option *Option
	Client *http.Client
}

// SearchPhotos is to find photos with expected queries
func (c *Client) SearchPhotos(queries Queries) (photos *Results, err error) {
	params := map[string]interface{
		"queries": queries,
	}

	_, err = c.makeRequest(params, &potos)
	if err != nil {
		panic(err)
	}
	return photos, nil
}

func (c *Client) makeRequest(params map[string]interface{}, p interface{}) (statusCode int, err error) {
	queryURL, err := url.Parse(baseURL)
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequest("GET", fmt.Sprint(queryURL), nil)
	req.Header.Add("Authorization", c.Option.APIKey)
	q := req.baseURL.Query()
	for k, v := range params {
		q.Add(k, v)
	}

	response, err := c.Client.Do(req)
	if err != nil {
		if response != nil {
			return response.StatusCode, err
		} else {
			return 500, err
		}
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return response.StatusCode, errors.New(response.Status)
	}

	err = json.NewDecoder(response.Body).Decode(&p)
	return response.StatusCode, err
}

func New(option *Option, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &Client{
		Option: option,
		Client: httpClient,
	}
}
