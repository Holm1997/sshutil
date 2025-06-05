package utils

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type HostsConfig struct {
	Hosts []struct {
		Name  string `yaml:"name"`
		Login string `yaml:"login"`
		Pass  string `yaml:"pass"`
	} `yaml:"hosts"`
}

func ParseYamlFile(file *string) HostsConfig {
	hc := HostsConfig{}
	f, err := os.ReadFile(*file)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(f, &hc)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return hc
}
