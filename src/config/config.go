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
