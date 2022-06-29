package parserConfig

type GetSpare struct {
	Token      string `yaml:"token" mapstructure:"token"`
	RequestUrl string `yaml:"requestUrl" mapstructure:"requestUrl"`
}
