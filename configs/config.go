package configs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"projects/review-finder/tools/alerts"
	"projects/review-finder/tools/email"
)

type Config struct {
	Alerts        []alerts.AlertSettings `json:"alert_settings"`
	EmailSettings email.ServiceConfig    `json:"email_settings"`
}

func Init() (Config, error) {
	cfg := Config{}

	cfgFile, err := ioutil.ReadFile("configs/cfg.json")

	if err != nil {
		return cfg, fmt.Errorf("error reading config file, err: %w", err)
	}

	err = json.Unmarshal(cfgFile, &cfg)
	if err != nil {
		return cfg, fmt.Errorf("error unmarshaling config file data, err: %w", err)
	}

	return cfg, nil
}
