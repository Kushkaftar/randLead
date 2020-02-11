package main

import (
	"flag"
	"fmt"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/config.toml", "path to config file")
}

func main() {
	flag.Parse()
	fmt.Println("helo world")
}
