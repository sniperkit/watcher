package main

import (
	"flag"
	"log"
	"time"

	"github.com/jasonlvhit/gocron"
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
	gocron.Every(2).Hours().Do(run)
	<-gocron.Start()
}

func run() {
	start := time.Now()
	log.Printf("Started at: %s", start.String())
	config.Init(*enviroment)
	db.Init()
	data, err := document.Init()
	if err != nil {
		log.Println(err)
	}
	extractor.Runner(data)
	elapsed := time.Since(start)
	log.Printf("Awesome-iOS Parser took %s", elapsed)
	ended := time.Now()
	log.Printf("Ended at: %s", ended.String())
}
