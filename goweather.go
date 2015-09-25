package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	m "math"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

const (
	API = "http://api.openweathermap.org/data/2.5/forecast/daily"
)
var (
	key      string
	val      string
	unit     string
	wind     bool
	days     int
	emptyLoc bool
)
var windDirection = map[float64]string{
	0:     "North",
	22.5:  "North north east",
	45:    "North east",
	67.5:  "East north east",
	90:    "East",
	112.5: "East south east",
	135:   "South east",
	157.5: "South south east",
	180:   "South",
	202.5: "South south west",
	225:   "South west",
	247.5: "West south west",
	270:   "West",
	292.5: "West north west",
	315:   "North west",
	337.5: "North north west",
}


func help() {
	fmt.Printf("Usage: goweather [flags] location\n")
	fmt.Printf("location: city name or zip code\nflags:\n")
	flag.PrintDefaults()
}


func exitHelp() {
	help()
	os.Exit(3)
}


func init() {

	flag.Usage = help

	flag.IntVar(&days, "days", 1, "Shows forecasts for number of days (1-16)" )
	flag.StringVar(&unit, "unit", "imperial", "Imperial, metric, or kelvin units of measurement")
	flag.BoolVar(&wind, "wind", false, "Show wind")
	flag.Parse()

	if days < 1 || days > 16 {
		fmt.Println("Days must be between 1 and 16")
		exitHelp()
	}

	if unit != "imperial" && unit != "metric" && unit != "kelvin" {
		fmt.Println("Units must be imperial, metric, or kelvin")
		exitHelp()
	}

	val = flag.Arg(0)
	_, err := strconv.Atoi(val)

	if err == nil {
		if len(val) != 5 {
			exitHelp()
		}
		key = "zip"
	} else if len(val) != 0 {
		key = "q"
	} else {
		zip, err := determineZip()
		if err != nil {
			exitHelp()
		}
		key = "zip"
		val = zip
		emptyLoc = true
	}

}


type info struct {
	Zip     string `json:"postal"`
	City    string `json:"city"`
	Region  string `json:"region"`
	Country string `json:"country"`
}
var location info


func determineZip() (string, error) {
	resp, err := http.Get("http://ipinfo.io/geo")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&location)
	if err != nil {
		return "", err
	}

	if location.Zip == "" {
		return "", fmt.Errorf("unable to determine zip code")
	}

	return location.Zip, nil
}


func escape(s string) string {
	return url.QueryEscape(s)
}


func sendRequest() {
	params := fmt.Sprintf("?%s=%s&units=%s&cnt=%d", key, escape(val), escape(unit), days)
	resp, err := http.Get(API + params)
	if err != nil {
		fmt.Println("Failed to get url")
		os.Exit(3)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println("Failed to get data")
		os.Exit(3)
	}

	handleResponse(resp.Body)
}


type WeatherResponse struct {
	List []ListType `json:"list"`
}
type ListType struct {
	Dt            int64         `json:"dt"`
	Temp          TempType      `json:"temp"`
	Weather       []WeatherType `json:"weather"`
	WindSpeed     float64       `json:"speed"`
	WindDirection float64       `json:"deg"`
}
type TempType struct {
	Day float64 `json:"day"`
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}
type WeatherType struct {
	Description string `json:"description"`
}


func parseTime(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return fmt.Sprintf("%s, %s %02d, %d", t.Weekday(), t.Month(), t.Day(), t.Year())
}


func handleResponse(s io.ReadCloser ) {
	var f WeatherResponse

	err := json.NewDecoder(s).Decode(&f)
	if err != nil {
		fmt.Println("Failed to parse body", err)
		os.Exit(3)
	}

	if emptyLoc {
		fmt.Printf("Determined location: %s, %s, %s\n", location.City, location.Region, location.Zip)
	}

	for i := range f.List {

		t := f.List[i]

		row_1 := " %-15s%-15s%-15s%-20s"
		row_2 := " %-15.2f%-15.2f%-15.2f%-20s"

		if wind {
			row_1 = row_1 + "%-15s%-15s"
			row_2 = row_2 + "%-15.2f%-15s"
		}

		fmt.Println(parseTime(f.List[i].Dt))

		if wind {
			fmt.Printf(row_1 + "\n", "Current temp", "Today's high", "Today's low", "Condition", "Wind speed", "Wind direction")
			fmt.Printf(row_2 + "\n\n", t.Temp.Day, t.Temp.Max, t.Temp.Min, t.Weather[0].Description, t.WindSpeed, getWindDirection(t.WindDirection))
		} else {
			fmt.Printf(row_1 + "\n", "Current temp", "Today's high", "Today's low", "Condition")
			fmt.Printf(row_2 + "\n\n", t.Temp.Day, t.Temp.Max, t.Temp.Min, t.Weather[0].Description)
		}
	}
}


func getWindDirection(deg float64) string {
	return windDirection[getClosestNum(deg)]
}


func getClosestNum(deg float64) float64 {
	// var myNumber float64 = 490 //deg

	var numbers = make([]float64, len(windDirection))
	for k := range windDirection {
		numbers = append(numbers, k)
	}

	distance := m.Abs(numbers[0] - deg)
	idx := 0
	for c := 1; c < len(numbers); c++ {
	    cdistance := m.Abs(numbers[c] - deg)
	    if cdistance < distance {
	        idx = c
	        distance = cdistance
	    }
	}
	theNumber := numbers[idx]
	return theNumber
}


func main() {
	sendRequest()
}