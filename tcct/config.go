package tcct

import (
	"os"
	"encoding/json"
	"log"
)

type Config struct {
	// Basic config structure
	Name string
	Server string
	Api_key string
	Channel string
	Port string
}

func GetConfig() (string, string, string, string, string) {
	// This will return the config variables
	// Code copy-pasted from stackoverflow
	// Link: https://stackoverflow.com/questions/16465705/how-to-handle-configuration-in-go
	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	conf := Config{}
	err := decoder.Decode(&conf)
	if err != nil {
		log.Panicln(err)
	}
	return conf.Name, conf.Server,
		   conf.Api_key, conf.Channel,
		   conf.Port
}