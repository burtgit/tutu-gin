package config

type DataBase struct {
	Username    string `yaml:"username" mapstructure:"username"`
	Password    string `yaml:"password" mapstructure:"password"`
	Port        string `yaml:"port" mapstructure:"port"`
	Host        string `yaml:"host" mapstructure:"host"`
	Dbname      string `yaml:"dbname" mapstructure:"dbname"`
	TablePrefix string `yaml:"tablePrefix" mapstructure:"tablePrefix"`
}
