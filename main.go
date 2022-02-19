package main

import (
	"encoding/json"
	"fmt"
	"github.com/hugolgst/rich-go/client"
	"github.com/pkg/errors"
	"os"
	"time"
)

type Config struct {
	ClientID string
	client.Activity
}

func main() {
	config, err := readConfigFile("config.json")
	if err != nil {
		panic(err)
	}
	//now := time.Now()
	//config.Timestamps.Start = &now

	err = client.Login(config.ClientID)
	if err != nil {
		panic(err)
	}

	err = client.SetActivity(config.Activity)
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}

	fmt.Println("all good. close me when you're done looking cool.")
	for {
		time.Sleep(time.Second * 100)
	}
}
func parseConfig(cfg *Config) error {
	if cfg.Timestamps.End == nil && cfg.Timestamps.Start == nil {
		now := time.Now()
		cfg.Timestamps.Start = &now
	}
	return nil
}
func readConfigFile(fileName string) (*Config, error) {
	dat, err := os.ReadFile(fileName)
	if err != nil {
		return nil, errors.Wrap(err, "failed opening file")
	}
	var config Config
	err = json.Unmarshal(dat, &config)
	if err != nil {
		return nil, errors.Wrap(err, "failed unmarshalling json")
	}
	return &config, parseConfig(&config)
}
