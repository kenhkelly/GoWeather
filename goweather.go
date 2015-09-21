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
)

const (
	API = "http://api.openweathermap.org/data/2.5/weather"
)

var (
	key  string
	val  string
	unit string
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

	flag.StringVar(&unit, "unit", "imperial", "Imperial or metric units of measurement")
	flag.Parse()

	val = flag.Arg(0)
	_, err := strconv.Atoi(val)

	if err == nil {
		if len(val) != 5 {
			exitHelp()
		}
		key = "zip"
	} else if val != "" {
		key = "q"
	} else {
		exitHelp()
	}

}

func escape(s string) string {
	return url.QueryEscape(s)
}

func sendRequest() {

	params := fmt.Sprintf("?%s=%s&units=%s", key, escape(val), escape(unit))
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
	Main struct {
		Temp     float64
		Temp_Min float64
		Temp_Max float64
	}
	Weather []struct {
		Description string
	}
}

func handleResponse(s io.ReadCloser ) {

	var f WeatherResponse

	err := json.NewDecoder(s).Decode(&f)
	if err != nil {
		fmt.Println("Failed to parse body")
		os.Exit(3)
	}

	row_1 := "%-15s%-15s%-15s%-20s\n"
	row_2 := "%-15.2f%-15.2f%-15.2f%-20s\n\n"

	fmt.Printf(row_1, "Current temp", "Today's high", "Today's low", "Condition")
	fmt.Printf(row_2, f.Main.Temp, f.Main.Temp_Max, f.Main.Temp_Min, f.Weather[0].Description)

}

func main() {
	sendRequest()
}

