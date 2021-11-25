package yaweather

import (
	"fmt"
	"time"
)

// Main response structure
type WeatherResponse struct {
	Now        int                `json:"now"`
	Now_dt     time.Time          `json:"now_dt"`
	Info       LocalityInfo       `json:"info"`
	Geo_object Geo_ObjectInfo     `json:"geo_object"`
	Fact       WeatherFact        `json:"fact"`
	Forecasts  []WeatherForecasts `json:"forecasts"`
}

// Coordinates
type LocalityInfo struct {
	Url             string       `json:"url"`
	Lat             float64      `json:"lat"`
	Lon             float64      `json:"lon"`
	Tzinfo          TimezoneInfo `json:"tzinfo"`
	Def_Pressure_Mm int          `json:"def_pressure_mm"`
	Def_Pressure_Pa int          `json:"def_pressure_pa"`
}

// Time zone info
type TimezoneInfo struct {
	// Offset int32  `json:"offset"`
	Name string `json:"name"`
	Abbr string `json:"abbr"`
	Dst  bool   `json:"dst"`
}

// Location info
type Geo_ObjectInfo struct {
	District GeoInfo `json:"district"`
	Locality GeoInfo `json:"locality"`
	Province GeoInfo `json:"province"`
	Country  GeoInfo `json:"country"`
}

type GeoInfo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (g Geo_ObjectInfo) GetLocationInfo() string {
	return fmt.Sprintf("Местоположение: %s, %s, %s, %s", g.Country.Name, g.Province.Name, g.Locality.Name, g.District.Name)
}

// Current weather structure
type WeatherFact struct {
	Temp        int    `json:"temp"`
	Feels_Like  int    `json:"feels_like"`
	Icon        string `json:"icon"` // https://yastatic.net/weather/i/icons/funky/dark/<value from the icon field>.svg
	Condition   string `json:"condition"`
	Wind_Dir    string `json:"wind_dir"`
	Pressure_Mm int    `json:"pressure_mm"`
	Humidity    int    `json:"humidity"`
}

func (c WeatherFact) GetFactWeatherInfo() string {
	return fmt.Sprintf("Температура воздуха: %d\nОщущается как: %d\nВетер: %s\nДавление: %d mm\nВлажность: %d %%",
		c.Temp, c.Feels_Like, c.Wind_Dir, c.Pressure_Mm, c.Humidity)
}

// Weather forecast struct
type WeatherForecasts struct {
	Date  string            `json:"date"`
	Parts PartOfDayForecast `json:"parts"`
	Hours []HourForecast    `json:"hours"`
}

// Forecasts by time of day
type PartOfDayForecast struct {
	Day         partForecast    `json:"day"`
	Evening     partForecast    `json:"evening"`
	Morning     partForecast    `json:"morning"`
	Night       partForecast    `json:"night"`
	Night_Short HalfDayForecast `json:"night_short"`
	Day_Short   HalfDayForecast `json:"day_short"`
}

type partForecast struct {
	Temp_Min    int16  `json:"temp_min"`
	Temp_Max    int16  `json:"temp_max"`
	Icon        string `json:"icon"`
	Condition   string `json:"condition"`
	Wind_Dir    string `json:"wind_dir"`
	Pressure_Mm int16  `json:"pressure_mm"`
	Humidity    int16  `json:"humidity"`
}

// 12-hour forecast for the day
type HalfDayForecast struct {
	Temp        int    `json:"temp"`
	Icon        string `json:"icon"`
	Condition   string `json:"condition"`
	Wind_Dir    string `json:"wind_dir"`
	Pressure_Mm int    `json:"pressure_mm"`
	Humidity    int    `json:"humidity"`
}

// hourly forecast
type HourForecast struct {
	Hour        string `json:"hour"`
	Temp        int    `json:"temp"`
	Icon        string `json:"icon"`
	Condition   string `json:"condition"`
	Wind_Dir    string `json:"wind_dir"`
	Pressure_Mm int    `json:"pressure_mm"`
	Humidity    int    `json:"humidity"`
}
