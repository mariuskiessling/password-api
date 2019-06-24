package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
)

// The Config struct represents the unmarshalled config file.
type Config struct {
	Port        int    `json:"port"`
	Environment string `json:"env"`
}

// Load loads the given config file and maps it to an internal Config struct.
func Load(filename string) (cfg *Config, err error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.New("Config file `" + filename + "` could not be opened.")
	}

	c := &Config{}
	json.Unmarshal(file, c)
	err = c.SetupEnvironment()
	return c, err
}

// SetupEnvironment configures the logger and rejects any config environment
// string that is not "production" or "development". In an production
// environment all logging is disabled in order to hide the generated
// passwords. This is not perfect but the best we can do without introducing a
// better logging framework.
func (cfg *Config) SetupEnvironment() (err error) {
	switch cfg.Environment {
	case "production":
		log.Printf("Using production environment. No logging output will be displayed.")
		log.SetFlags(0)
		log.SetOutput(ioutil.Discard)
		return nil
	case "development":
		log.Println("Using development environment.")
		// Do nothing because the logger outputs everything by default
		return nil
	default:
		return errors.New("Config includes invalid environment")
	}
}
