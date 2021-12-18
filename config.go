package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type TemplatesConfig struct {
	Templates map[string]Template `yaml:"templates"`
}

type Template struct {
	Type    string   `yaml:"type"`
	Help    string   `yaml:"help"`
	Message string   `yaml:"message"`
	Actions []Action `yaml:"actions"`
}

func (t Template) GetTitle() string {
	switch t.Type {
	case "error":
		return "Ошибка"
	case "success":
		return "Успех"
	default:
		return ""
	}
}

type Action struct {
	URL  string `yaml:"url"`
	Text string `yaml:"text"`
}

var CurrentConfig = &TemplatesConfig{}

func LoadConfig() error {
	data, err := ioutil.ReadFile("config.yml")
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, CurrentConfig)
}
