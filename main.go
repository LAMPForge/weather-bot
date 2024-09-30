package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	_ "time/tzdata"

	"gopkg.in/robfig/cron.v2"
)

type ControllerIF interface {
	request() *WeatherResponse
	sendMessageToTelegram(message string)
	set_cron(cronExpr string, cronFunc func())
	sendWeatherUpdate()
}

type Controller struct {
	config        *Config
	cronScheduler *cron.Cron
	mu            sync.Mutex
}

func InitController(config *Config) Controller {
	return Controller{config: config, cronScheduler: cron.New()}
}

func run(c ControllerIF) {
	c.set_cron("TZ=Asia/Bangkok 0 16 * * *", func() {
		c.sendWeatherUpdate()
	})
}

func main() {
	config := getEnv()
	controller := InitController(config)
	run(&controller)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (c *Controller) sendWeatherUpdate() {
	c.mu.Lock()
	defer c.mu.Unlock()

	weatherData := c.request()
	message := formatWeatherHTML(*weatherData)
	c.sendMessageToTelegram(message)
	fmt.Println("Weather update sent.")
}
