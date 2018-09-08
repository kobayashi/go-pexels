package pexels

import (
	"net/http"
)

type Src struct {
	Original  string `json:"original"`
	Large2X   string `json:"large2x"`
	Large     string `json:"large"`
	Medium    string `json:"medium"`
	Small     string `json:"small"`
	Portrait  string `json:"portrait"`
	Square    string `json:"square"`
	LandScape string `json:"landscape"`
	Tiny      string `json:"tiny"`
}

type Photo struct {
	ID           int    `json:"id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	URL          string `json:"url"`
	Photographer string `json:"photographer"`
	Src          Src    `json:"src"`
}

type Results struct {
	TotalResults int     `json:"total_results"`
	Page         int     `json:"page"`
	PerPage      int     `json:"per_page"`
	Photos       []Photo `json:"photos"`
	NextPage     string  `json:"next_page"`
}

// SearchPhotos returns photos based requested query.
func (s *Service) SearchPhotos(params *SearchParams) (Results, *http.Response, error) {
	photos := new(Results)
	pexelsError := new(PexelsError)
	resp, err := s.sling.New().Get("v1/search").QueryStruct(params).Receive(photos, pexelsError)
	if err == nil {
		err = pexelsError
	}
	return *photos, resp, err
}

// CuratedPhotos returns selection of trending photos.
func (s *Service) CuratedPhotos(params *TrendSearchParams) ([]Photo, *http.Response, error) {
	photos := new([]Photo)
	pexelsError := new(PexelsError)
	resp, err := s.sling.New().Get("v1/curated").QueryStruct(params).Receive(photos, pexelsError)
	if err == nil {
		err = pexelsError
	}
	return *photos, resp, err
}

// GetPhoto returns a specific photo
func (s *Service) GetPhoto(id string) (Photo, *http.Response, error) {
	photo := new(Photo)
	pexelsError := new(PexelsError)
	resp, err := s.sling.New().Get("v1/photos/"+id).Receive(photo, pexelsError)
	if err == nil {
		err = pexelsError
	}
	return *photo, resp, err
}
