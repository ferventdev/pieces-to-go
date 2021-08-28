package main

import (
	"log"
	"time"

	"github.com/BurntSushi/toml"
)

type tomlConfig struct {
	Title   string
	Owner   ownerInfo
	DB      database `toml:"database"`
	Servers map[string]server
	Clients clients
}

type ownerInfo struct {
	Name string
	Org  string `toml:"organization"`
	Bio  string
	DOB  time.Time
}

type database struct {
	Server  string
	Ports   []int
	ConnMax int `toml:"connection_max"`
	Enabled bool
}

type server struct {
	IP string
	DC string
}

type clients struct {
	Data  [][]interface{}
	Hosts []string
}

func main() {
	var cfg tomlConfig
	if _, err := toml.DecodeFile("config.toml", &cfg); err != nil {
		log.Fatalln("config file parsing error:", err)
	}

	log.Println(cfg)
}
