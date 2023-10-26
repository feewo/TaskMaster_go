package config

import (
	"os"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server Server
	Api    map[string]Service
}
type Service struct {
	Method string
	Url    string
}
type Server struct {
	Port    string
	Timeout string
}

var config Config

func Get() *Config {
	return &config
}

var isEnv = regexp.MustCompile(`\$\{(.*?)\}`)

func init() {
	file, err := os.ReadFile("api.yml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}

	if isEnv.MatchString(config.Server.Port) {
		find := isEnv.FindStringSubmatch((config.Server.Port))
		sub := strings.SplitN(find[1], ":", 2)
		envOs := os.Getenv(sub[0])
		if len(envOs) > 0 {
			config.Server.Port = envOs
		} else {
			config.Server.Port = sub[1]
		}
	}

}
