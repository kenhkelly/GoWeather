package ipinfo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func GetLocation() (string, string, error) {

	resp, err := http.Get(API_URL)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	var location Location
	err = json.NewDecoder(resp.Body).Decode(&location)
	if err != nil {
		return "", "", err
	}

	if location.Loc == "" {
		return "", "", fmt.Errorf("unable to determine location")
	}

	latLng := strings.Split(location.Loc, ",")

	return latLng[0], latLng[1], nil

}
