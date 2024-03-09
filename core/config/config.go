package config

import "tutu-gin/core/config/parserConfig"

type ServiceConfig struct {
	Redis        Redis                    `mapstructure:"redis" json:"redis" yaml:"redis"`
	DataBase     DataBase                 `json:"dataBase" yaml:"dataBase" mapstructure:"database"`
	YtDl         YtDl                     `json:"yt-dl" yaml:"yt-dl" mapstructure:"yt-dl"`
	ParserConfig parserConfig.ParseConfig `yaml:"parser" json:"parser" mapstructure:"parser"`
}
