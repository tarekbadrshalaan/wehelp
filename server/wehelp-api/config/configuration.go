package config

import (
	"fmt"
	"sync"

	"github.com/kelseyhightower/envconfig"
	"github.com/tarekbadrshalaan/goStuff/configuration"
)

var (
	readConfigOnce sync.Once
	internalConfig Config
)

// Config : application configuration
type Config struct {
	DBConnectionString string `json:"WEHELP_API_DB_CONNECTION_STRING" envconfig:"WEHELP_API_DB_CONNECTION_STRING"`
	DBEngine           string `json:"WEHELP_API_DB_ENGINE" envconfig:"WEHELP_API_DB_ENGINE"`
	WebAddress         string `json:"WEHELP_API_ADDRESS" envconfig:"WEHELP_API_ADDRESS"`
	WebPort            int    `json:"WEHELP_API_PORT" envconfig:"WEHELP_API_PORT"`
	SigendKey          string `json:"WEHELP_API_SIGNED_KEY" envconfig:"WEHELP_API_SIGNED_KEY"`
	ValidationDuration int    `json:"WEHELP_API_VALIDATION_DURATION" envconfig:"WEHELP_API_VALIDATION_DURATION"`
}

// Configuration : get configuration based on json file or environment variables
func Configuration() Config {
	readConfigOnce.Do(func() {
		err := configuration.JSON("config.json", &internalConfig)
		if err == nil {
			return
		}
		fmt.Println(err)

		err = envconfig.Process("", &internalConfig)
		if err != nil {
			err = fmt.Errorf("Error while initiating app configuration : %v", err)
			panic(err)
		}
	})
	return internalConfig
}
