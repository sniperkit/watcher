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
		elastic.SetMaxRetries(2),
	)
	if err != nil {
		log.Fatal("Error on connecting to database")
	}
	db = client
}

func GetDB() *elastic.Client {
	return db
}
