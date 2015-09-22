package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
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
	days     int
	emptyLoc bool
)

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

	flag.StringVar(&unit, "unit", "imperial", "Imperial, metric, or kelvin units of measurement")
	flag.IntVar(&days, "days", 1, "Shows forecasts for number of days (1-16)" )
	helpPtr := flag.Bool("help", false, "Shows this help")
	flag.Parse()

	if *helpPtr == true {
		help()
		os.Exit(0)
	}

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
	List []ListType
}
type ListType struct {
	Dt int64
	Temp TempType
	Weather []WeatherType
}
type TempType struct {
	Day float64
	Min float64
	Max float64
}
type WeatherType struct {
	Description string
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
		row_1 := " %-15s%-15s%-15s%-20s\n"
		row_2 := " %-15.2f%-15.2f%-15.2f%-20s\n\n"

		fmt.Println(parseTime(f.List[i].Dt))
		fmt.Printf(row_1, "Current temp", "Today's high", "Today's low", "Condition")
		fmt.Printf(row_2, f.List[i].Temp.Day, f.List[i].Temp.Max, f.List[i].Temp.Min, f.List[i].Weather[0].Description)
	}
}

func main() {
	sendRequest()
}