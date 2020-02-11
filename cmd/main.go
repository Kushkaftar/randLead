package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"

	"github.com/Kushkaftar/randLead/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/config.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("helo world")
}
