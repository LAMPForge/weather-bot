package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type WeatherResponse struct {
	Weather []Weather `json:"weather"`
	Sys     Sys       `json:"sys"`
	Rain    Rain      `json:"rain"`
	Name    string    `json:"name"`
}

type Rain struct {
	OneHour float64 `json:"1h"`
}

type Weather struct {
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Sys struct {
	Country string `json:"country"`
	Sunrise int    `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

func (c *Controller) request() *WeatherResponse {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s", c.config.LAT, c.config.LONG, c.config.API_KEY)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to get response: %s", resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	var weatherData WeatherResponse
	err = json.Unmarshal([]byte(body), &weatherData)
	if err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}
	return &weatherData

}

func formatWeatherHTML(data WeatherResponse) string {
	return fmt.Sprintf(
		"🌦 <b>Weather Report for %s</b>\n\n"+
			"<b>🌡️ Condition:</b> <i>%s</i> - %s\n"+
			"🌧️ <b>Rainfall:</b> %.2f mm (last hour)\n"+
			"🕐 <i>Report generated at:</i> %s\n",
		data.Name,
		data.Weather[0].Main,
		data.Weather[0].Description,
		data.Rain.OneHour,
		formatTime(data.Sys.Sunset),
	)
}

func formatTime(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format("15:04:05 MST, Jan 2 2006") // Customize the time format as needed
}
