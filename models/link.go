package models

import (
	"log"
	"strconv"

	"github.com/vsouza/watcher/db"
)

type Link struct {
	ID       int64  `json:"-" bson:"_id,omitempty"`
	Type     string `json:"-" bson:"type"`
	Name     string `json:"name" bson:"name,omitempty"`
	Category string `json:"-" bson:"category,omitempty"`
	FullName string `json:"full_name" bson:"full_name,omitempty"`
	URL      string `json:"html_url" bson:"url,omitempty"`
	Owner    struct {
		Login     string `json:"login"`
		Id        int64  `json:"id"`
		AvatarURL string `json:"avatar_url"`
		Type      string `json:"type"`
	} `json:"owner" bson:"owner,omitempty"`
	Description     string `json:"description" bson:"description,omitempty"`
	Homepage        string `json:"homepage" bson:"homepage,omitempty"`
	Stars           int32  `json:"stargazers_count" bson:"stargazers_count,omitempty"`
	Watchers        int32  `json:"watchers_count" bson:"watchers_count,omitempty"`
	Language        string `json:"language" bson:"language,omitempty"`
	HasWiki         bool   `json:"has_wiki" bson:"has_wiki,omitempty"`
	HasIssues       bool   `json:"has_issues" bson:"has_issues,omitempty"`
	OpenIssuesCount int32  `json:"open_issues_count" bson:"open_issues_count,omitempty"`
	ForksCount      int32  `json:"forks" bson:"forks_count,omitempty"`
	DefaultBranch   string `json:"default_branch" bson:"default_branch,omitempty"`
	CreatedAt       string `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt       string `json:"updated_at" bson:"updatedAt,omitempty"`
	PackageManagers struct {
		Carthage  bool `json:"carthage"`
		SPM       bool `json:"spm"`
		CocoaPods bool `json:"cocoa_pods"`
	} `json:"package_managers"`
}

func (l Link) SaveData(link *Link) error {
	client := db.GetDB()

	var err error
	_, err = client.Index().
		Index("links").
		Type(link.Type).
		Id(strconv.Itoa(link.ID)).
		BodyJson(&link).
		Refresh(true).
		Do()
	return err
}

func (l Link) Flush() {
	client := db.GetDB()
	_, err = client.Flush().Index("links").Do()
	if err != nil {
		log.Println("error on flush db")
	}
}
