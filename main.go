package main

import (
	"flag"
	"log"

	"github.com/vsouza/watcher/config"
	"github.com/vsouza/watcher/db"
)

var enviroment = flag.String("e", "development", "which environment do you wanna start ?")

func main() {
	flag.Parse()
	if *enviroment == "" {
		log.Fatal("enviroment must be set")
	}
	config.Init(*enviroment)
	db.Init()
}
