package models

import (
	"log"

	"github.com/vsouza/watcher/db"
)

type Category struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	UpdatedAt  string `json:"updated_at"`
	Parent     string `json:"parent"`
	MainParent string `json:"main_parent"`
}

func (c Category) SaveData(category *Category) error {
	client := db.GetDB()
	indexType := "nested"
	if len(category.Parent) == 0 {
		indexType = "main"
	}
	var err error
	_, err = client.Index().
		Index("categories").
		Type(indexType).
		Id(category.ID).
		BodyJson(&category).
		Refresh(true).
		Do()
	return err
}

func (c Category) Flush() {
	client := db.GetDB()
	_, err = client.Flush().Index("categories").Do()
	if err != nil {
		log.Println("error on flush db")
	}
}
