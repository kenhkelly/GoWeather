package forecast

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetForecast(key, lat, lng string) (*Forecast, error) {
	var fc *Forecast

	url := fmt.Sprintf(API_URL, key, lat, lng)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(&fc)
	if err != nil {
		return nil, err
	}

	return fc, nil

}
