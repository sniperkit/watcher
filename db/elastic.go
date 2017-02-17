package db

import (
	"log"

	"github.com/vsouza/watcher/config"

	"gopkg.in/olivere/elastic.v3"
)

var db *elastic.Client

func Init() {
	c := config.GetConfig()
	client, err := elastic.NewClient(
		elastic.SetURL(c.GetString("db.host")),
	)
	client.Start()
	if err != nil {
		log.Fatalf("error connecting on database. %s", err)
	}
	db = client
}

func GetDB() *elastic.Client {
	return db
}
