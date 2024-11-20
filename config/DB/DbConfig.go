package DB

import (
	 "gopkg.in/yaml.v2"
	 "os"
)

type DbConfig struct {
	Username string
	Password string
	Host     string
	Db       string
}

func LoadConfig() (DbConfig, error) {
	fileData, err := os.ReadFile("config/local.yml")
	if err != nil {
		return DbConfig{}, err
	}

	var dbConfig DbConfig
	err = yaml.Unmarshal(fileData, &dbConfig)
	if err != nil {
		return DbConfig{}, err
	}

	return dbConfig, nil
}