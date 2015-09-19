package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Need to enter zip code")
		os.Exit(3)
	}

	arg := os.Args[1]

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&units=imperial", arg)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to get url")
		os.Exit(3)
	}
	defer resp.Body.Close()
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		fmt.Println("Read error")
		os.Exit(3)
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

	var f WeatherResponse

	errr := json.Unmarshal(body, &f)
	if errr != nil {
		fmt.Println("Failed to parse body")
		os.Exit(3)
	}

	fmt.Printf("%-15s%-15s%-15s%-20s\n", "Current temp", "Today's high", "Today's low", "Condition")
	fmt.Printf("%-15.2f%-15.2f%-15.2f%-20s\n\n", f.Main.Temp, f.Main.Temp_Max, f.Main.Temp_Min, f.Weather[0].Description)
}
