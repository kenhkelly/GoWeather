package forecast

const (
	API_URL = "https://api.forecast.io/forecast/%s/%s,%s"
)

type Forecast struct {
	Currently struct {
		Time        int64
		Summary     string
		Temperature float32
		Humidity    float32
		WindSpeed   float32
		WindBearing float32
	} `json="currently"`
	Daily struct {
		Summary string
		Data    []struct {
			Time           int64
			Summary        string
			TemperatureMin float32
			TemperatureMax float32
			Humidity       float32
			WindSpeed      float32
			WindBearing    float32
		} `json="data"`
	} `json="daily"`
}
