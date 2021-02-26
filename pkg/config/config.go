package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

var AppConfig = LoadConfig()

type Config struct {
	RunMode 			string  `yaml:"run_mode"`
	Server struct{
		Port			string	`yaml:"port"`
		ReadTimeout		int		`yaml:"read_timeout"`
		WriteTimeout	int		`yaml:"write_timeout"`
	}
	DataBase struct {
		Type 			string	`yaml:"type"`
		User 			string	`yaml:"user"`
		Password		string  `yaml:"password"`
		Host 			string  `yaml:"host"`
		Name 			string  `yaml:"name"`
	}
	Oauth struct{
		GithubClientID 		string  `yaml:"github_client_id"`
		GithubClientSecret	string  `yaml:"github_client_secret"`
	}

}

func LoadConfig() Config {
	var config Config
	fp, err:=os.ReadFile("config/config.yml")
	if err != nil{
		log.Println(err)
	}
	err = yaml.Unmarshal(fp, &config)
	if err != nil{
		log.Println(err)
	}
	return config
}