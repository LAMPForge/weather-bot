package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func (c *Controller) sendMessageToTelegram(message string) {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", c.config.TELEGRAM_BOT_TOKEN)

	data := url.Values{}
	data.Set("chat_id", c.config.CHAT_ID)
	data.Set("text", message)
	data.Set("parse_mode", "HTML")

	resp, err := http.PostForm(apiURL, data)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Fatalf("Failed to send message: %s\nResponse Body: %s", resp.Status, string(body))
	}
}
