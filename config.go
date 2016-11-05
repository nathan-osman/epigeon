package main

import (
	"encoding/json"
	"os"
	"reflect"
	"strings"
)

// Config stores the configuration for the application.
type Config struct {
	SMTPAddress           string `json:"smtp_address"`
	TwitterConsumerKey    string `json:"twitter_consumer_key"`
	TwitterConsumerSecret string `json:"twitter_consumer_secret"`
	TwitterAccessToken    string `json:"twitter_access_token"`
	TwitterAccessSecret   string `json:"twitter_access_secret"`
}

// Default returns the default configuration for the application.
func Default() *Config {
	return &Config{
		SMTPAddress: ":25",
	}
}

// LoadFromFile loads the configuration from a JSON file.
func LoadFromFile(name string) (*Config, error) {
	r, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	c := Default()
	if err := json.NewDecoder(r).Decode(c); err != nil {
		return nil, err
	}
	return c, nil
}

// LoadFromEnv loads configuration from environment variables.
func LoadFromEnv() *Config {
	var (
		c      = Default()
		st     = reflect.ValueOf(c).Elem()
		stType = st.Type()
	)
	for i := 0; i < st.NumField(); i++ {
		var (
			field     = st.Field(i)
			fieldType = stType.Field(i)
			key       = fieldType.Tag.Get("json")
		)
		if len(key) == 0 {
			continue
		}
		val := os.Getenv(strings.ToUpper(key))
		if len(val) != 0 {
			field.SetString(val)
		}
	}
	return c
}
