package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Stages map[string]Stage
}
type Stage struct {
	Fields []string
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func readConfig() Config {
	f, err := os.Open("src/config/config.yml")
	var config Config
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		processError(err)
	}
	return config
}
func IsStage(stage string) bool {
	config := readConfig().Stages
	for k := range config {
		if k == stage {
			return true
		}
	}
	return false
}
func IsFieldInStage(stage string, field string) bool {
	config := readConfig().Stages
	for _, v := range config[stage].Fields {
		if v == field {
			return true
		}
	}
	return false
}
func HandleConfig(params ...string) {
	switch len(params) {
	case 0:
		fmt.Println(`hafifa config is used to handle your credentials.
		Usage:
			hafifa config show
			hafifa config set [field]`)
	case 1:
		if params[0] == "show" {
			// Implement showing the .env file here
		} else if params[0] == "set" {
			fmt.Println(`Please specify a value you want to set.
			The possible values are:
			username`)
		} else {
			fmt.Println("This is not a correct usage of hafifa config.")
		}
	case 2:
		if params[0] == "set" && params[1] == "username" {
			// Implement here setting a username in env file
		} else {
			fmt.Println("This is not a correct usage of hafifa config.")
		}
	}

}
