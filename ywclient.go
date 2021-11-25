package yaweather

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// HTTTP-client structure
type Client struct {
	client *http.Client
	apikey string
}

func NewClient(apikey string, timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("timeout can't be zero")
	}
	return &Client{
		client: &http.Client{
			Timeout: timeout,
		},
		apikey: apikey,
	}, nil
}

// Get weather method
func (c Client) GetWeather(lat, lon float64) (WeatherResponse, error) {
	url := fmt.Sprintf("https://api.weather.yandex.ru/v2/forecast?lat=%f&lon=%f&limit=1&hours=true&extra=false", lat, lon)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Add Header to request
	req.Header.Add("X-Yandex-API-Key", c.apikey)

	// Receive response
	resp, err := c.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return WeatherResponse{}, err
	}

	var w WeatherResponse

	if err = json.Unmarshal(body, &w); err != nil {
		return WeatherResponse{}, err
	}
	return w, nil
}
