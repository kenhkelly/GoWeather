package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

const (
	ErrEmptyIPInfoResp = "empty postal code in response from ipinfo.io"
)

func main() {
	var zip string

	if len(os.Args) < 2 {
		// attempt to guess the zip code
		var err error
		zip, err = DetermineZip()

		if err != nil {
			fmt.Println("Please provide zip code as argument.")
			os.Exit(3)
		}
	} else {
		zip = os.Args[1]
	}

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&units=imperial", url.QueryEscape(zip))

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to get url")
		os.Exit(3)
	}
	defer resp.Body.Close()

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

	var f WeatherResponse

	err = json.NewDecoder(resp.Body).Decode(&f)
	if err != nil {
		fmt.Println("Failed to parse body")
		os.Exit(3)
	}

	fmt.Printf("%-15s%-15s%-15s%-20s\n", "Current temp", "Today's high", "Today's low", "Condition")
	fmt.Printf("%-15.2f%-15.2f%-15.2f%-20s\n\n", f.Main.Temp, f.Main.Temp_Max, f.Main.Temp_Min, f.Weather[0].Description)
}

// DetermineZip calls ipinfo.io to determine the user's zip code
// via inspecting the ip address
func DetermineZip() (string, error) {
	resp, err := http.Get("http://ipinfo.io/geo")
	if err != nil {
		log.Println("error attempting to determine zip code ", err)
		return "", err
	}
	defer resp.Body.Close()

	var info struct {
		// omiting ip, city, region, loc, and phone area code
		Zip string `json:"postal"`
	}

	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		log.Println("unable to parse ip info to determine zip code ", err)
		return "", err
	}

	if info.Zip == "" {
		log.Println("unable to get ip info to determine zip code ", err)
		return "", fmt.Errorf(ErrEmptyIPInfoResp)
	}

	return info.Zip, nil
}
