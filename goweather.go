package main

import (
	"fmt"
	"net/http"
	"net/url"
	"encoding/json"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Need to enter zip code")
		os.Exit(3)
	}

	arg := os.Args[1]

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?zip=%s&units=imperial", url.QueryEscape(arg))

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
