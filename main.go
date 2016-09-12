package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/vsouza/watcher/config"
	"github.com/vsouza/watcher/db"
	"github.com/vsouza/watcher/document"
	"github.com/vsouza/watcher/extractor"
)

var enviroment = flag.String("e", "development", "which environment do you wanna start ?")

func main() {
	flag.Parse()
	if *enviroment == "" {
		log.Fatal("enviroment must be set")
	}
	config.Init(*enviroment)
	db.Init()
	data, err := document.Init()
	if err != nil {
		fmt.Println(err)
	}
	extractor.Runner(data)
	fmt.Println(data)
}
