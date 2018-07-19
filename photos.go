package pexels

type Src struct {
	Original  string `json:"original"`
	Large2X   string `json:"large2x"`
	Large     string `json:"large"`
	Medium    string `json:"medium"`
	Small     string `json:"small"`
	Portrait  string `json:"portrait"`
	Square    string `json:"square"`
	Landscape string `json:"landscape"`
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
