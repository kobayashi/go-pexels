package pexels

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

const baseURL = "https://api.pexels.com/"

type Service struct {
	sling *sling.Sling
}

type PexelsError struct {
	APIError string `json:"error"`
}

type SearchParams struct {
	Query   string `url:"query,omitempty"`
	PerPage int    `url:"per_page,omitempty"`
	Page    int    `url:"page,omitempty"`
}

type TrendSearchParams struct {
	PerPage int `url:"per_page,omitempty"`
	Page    int `url:"page,omitempty"`
}

// NewCLient returns a new Client
func NewClient(httpClient *http.Client, apikey string) *Service {
	return &Service{
		sling: sling.New().Client(httpClient).Base(baseURL).Set("Authorization", apikey),
	}
}

// Error returns api error
func (e PexelsError) Error() string {
	return fmt.Sprintf("pexles: %s", e.APIError)
}
