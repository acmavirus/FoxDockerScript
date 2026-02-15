// Copyright by AcmaTvirus
package system

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type NotificationSettings struct {
	TelegramEnabled bool   `json:"telegram_enabled"`
	TelegramToken   string `json:"telegram_token"`
	TelegramChatID  string `json:"telegram_chat_id"`
	DiscordEnabled  bool   `json:"discord_enabled"`
	DiscordWebhook  string `json:"discord_webhook"`
}

const settingsFile = "data/notifications.json"

func GetNotificationSettings() (NotificationSettings, error) {
	file, err := os.ReadFile(settingsFile)
	if err != nil {
		if os.IsNotExist(err) {
			return NotificationSettings{}, nil
		}
		return NotificationSettings{}, err
	}
	var settings NotificationSettings
	err = json.Unmarshal(file, &settings)
	return settings, err
}

func SaveNotificationSettings(settings NotificationSettings) error {
	data, err := json.MarshalIndent(settings, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(settingsFile, data, 0644)
}

func SendAlert(message string) {
	settings, err := GetNotificationSettings()
	if err != nil {
		fmt.Printf("Failed to get notification settings: %v\n", err)
		return
	}

	if settings.TelegramEnabled {
		sendTelegram(settings.TelegramToken, settings.TelegramChatID, message)
	}
	if settings.DiscordEnabled {
		sendDiscord(settings.DiscordWebhook, message)
	}
}

func sendTelegram(token, chatID, message string) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)
	payload := map[string]string{
		"chat_id": chatID,
		"text":    "🦊 *FoxDocker Alert*\n\n" + message,
		"parse_mode": "Markdown",
	}
	jsonPayload, _ := json.Marshal(payload)
	http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
}

func sendDiscord(webhookURL, message string) {
	payload := map[string]string{
		"content": "🦊 **FoxDocker Alert**\n" + message,
	}
	jsonPayload, _ := json.Marshal(payload)
	http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonPayload))
}
