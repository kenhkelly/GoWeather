package ipinfo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func GetLocation() (location Location, err error) {

	resp, err := http.Get(API_URL)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&location)
	if err != nil {
		return
	}

	if location.Loc == "" {
		return location, fmt.Errorf("unable to determine location")
	}

	latLng := strings.Split(location.Loc, ",")
	location.Lat = latLng[0]
	location.Lng = latLng[1]

	return location, nil

}
