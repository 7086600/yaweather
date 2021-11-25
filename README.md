# yaweather
Yandex Weather API-client

Get Yandex Weather API-key from https://developer.tech.yandex.ru/

Get latitude and longitude of place from maps.

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/7086600/yaweather"
)

const (
	API_KEY = "630e3b87-****-****-****-0664ffef72f7"
	LAT     = 53.0
	LON     = 70.0
)

func main() {
	yaw, err := yaweather.NewClient(API_KEY, time.Second*10)
	if err != nil {
		log.Fatal(err)
	}
	w, err := yaw.GetWeather(LAT, LON)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(w.Geo_object.GetLocationInfo())
	fmt.Println(w.Fact.GetFactWeatherInfo())
}
