package forecast

import (
	"encoding/json"
	"fmt"
	"net/http"
	"errors"
)

func GetForecast(key, lat, lng string) (*Forecast, error) {
	var fc *Forecast

	url := fmt.Sprintf(API_URL, key, lat, lng)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Forbidden (most likely invalid token)")
	}

	err = json.NewDecoder(resp.Body).Decode(&fc)
	if err != nil {
		return nil, err
	}

	return fc, nil

}
