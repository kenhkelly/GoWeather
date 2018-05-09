package ipinfo

const (
	API_URL = "http://ipinfo.io/geo"
)

type Info struct {
	Zip     string `json:"postal"`
	City    string `json:"city"`
	Region  string `json:"region"`
	Country string `json:"country"`
	Loc     string `json:"loc"`
	Lat     string `json:"lat"`
	Lng     string `json:"lng"`
}

type Location Info
