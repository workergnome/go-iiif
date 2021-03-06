package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
        Level	    LevelConfig	      `json:"level"`
        Graphics    GraphicsConfig    `json:"graphics"`
        Features    FeaturesConfig    `json:"features"`
	Images      ImagesConfig      `json:"images"`
	Derivatives DerivativesConfig `json:"derivatives"`
}

type LevelConfig struct {
     Compliance string `json:"compliance"`
}

type FeaturesConfig struct {
     Enable FeaturesToggle `json:"enable"`
     Disable FeaturesToggle `json:"disable"`
     Append FeaturesAppend `json:"append"`
}

type FeaturesToggle map[string][]string

type FeaturesAppend map[string]map[string]FeaturesDetails

type FeaturesDetails struct {
	Syntax    string `json:"syntax"`
	Required  bool   `json:"required"`
	Supported bool   `json:"supported"`
	Match     string `json:"match,omitempty"`
}

type ImagesConfig struct {
	Source SourceConfig `json:"source"`
	Cache  CacheConfig  `json:"cache"`
}

type DerivativesConfig struct {
	Cache CacheConfig `json:"cache"`
}

type GraphicsConfig struct {
	Source SourceConfig `json:"source"`
}

type SourceConfig struct {
	Name string `json:"name"`
	Path string `json:"path,omitempty"`
}

type CacheConfig struct {
	Name string `json:"name"`
	Path string `json:"path,omitempty"`
	TTL int `json:"ttl,omitempty"`
	Limit int `json:"limit,omitempty"`
}

func NewConfigFromFile(file string) (*Config, error) {

	body, err := ioutil.ReadFile(file)

	if err != nil {
		return nil, err
	}

	c := Config{}
	err = json.Unmarshal(body, &c)

	if err != nil {
		return nil, err
	}

	return &c, nil
}
