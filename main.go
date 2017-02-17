package main

import (
	"flag"
	"log"

	"github.com/bamzi/jobrunner"
	"github.com/gin-gonic/gin"
	"github.com/vsouza/watcher/config"
	"github.com/vsouza/watcher/db"
	"github.com/vsouza/watcher/document"
	"github.com/vsouza/watcher/extractor"
	"github.com/vsouza/watcher/status"
)

var enviroment = flag.String("e", "development", "which environment do you wanna start ?")

func main() {
	flag.Parse()
	if *enviroment == "" {
		log.Fatal("enviroment must be set")
	}
	jobrunner.Start()
	jobrunner.Schedule("@every 10m", Scrapper{})
	routes := gin.Default()
	routes.LoadHTMLGlob("views/Status.html")
	routes.GET("/jobrunner/html", status.JobHtml)
	routes.Run(":8888")
}

type Scrapper struct {
}

func (s Scrapper) Run() {
	log.Println("============================== STARTED ==========================")
	config.Init(*enviroment)
	db.Init()
	data, err := document.Init()
	if err != nil {
		log.Println(err)
	}
	extractor.Runner(data)
	log.Println("============================== ENDED ==========================")
}
