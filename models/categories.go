package models

import "github.com/vsouza/watcher/db"

type Categories struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name"`
	Parent     string `json:"parent,omitempty"`
	MainParent string `json:"main_parent,omitempty"`
	UpdatedAt  int64  `json:"updated_at,omitempty"`
}

func (c *Categories) SaveData(category *Categories) error {
	client := db.GetDB()
	var err error
	_, err = client.Index().
		Index("categories").
		Type("category").
		Id(category.ID).
		BodyJson(category).
		Refresh(true).
		Do()
	return err
}
