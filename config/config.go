package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

// The Config struct represents the unmarshalled config file.
type Config struct {
	Port int `json:"port"`
}

// Load loads the given config file and maps it to an internal Config struct.
func Load(filename string) (cfg *Config, err error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.New("Config file `" + filename + "` could not be opened.")
	}

	c := &Config{}
	json.Unmarshal(file, c)
	return c, nil
}
