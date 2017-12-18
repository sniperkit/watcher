package models

import (
	"log"
	"strconv"

	"github.com/sniperkit/watcher/db"
)

type AwesomeItem struct {
	ID       int64  `json:"id"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	Category string `json:"category"`
	FullName string `json:"full_name"`
	URL      string `json:"html_url"`
	Owner    struct {
		Login     string `json:"login"`
		Id        int64  `json:"id"`
		AvatarURL string `json:"avatar_url"`
		Type      string `json:"type"`
	} `json:"owner"`
	Description     string `json:"description"`
	Homepage        string `json:"homepage"`
	Stars           int32  `json:"stargazers_count"`
	Watchers        int32  `json:"watchers_count"`
	Language        string `json:"language"`
	HasWiki         bool   `json:"has_wiki"`
	HasIssues       bool   `json:"has_issues"`
	OpenIssuesCount int32  `json:"open_issues_count"`
	ForksCount      int32  `json:"forks"`
	DefaultBranch   string `json:"default_branch"`
	CreatedAt       string `json:"created_at,omitempty"`
	UpdatedAt       string `json:"updated_at,omitempty"`
	PackageManagers struct {
		Carthage  bool `json:"carthage"`
		SPM       bool `json:"spm"`
		CocoaPods bool `json:"cocoa_pods"`
	} `json:"package_managers"`
}

func (a AwesomeItem) SaveData(awesomeItem *AwesomeItem) error {

	var err error
	client := db.GetDB()

	if awesomeItem.Type == "url" {
		body := map[string]interface{}{
			"url":      awesomeItem.URL,
			"category": awesomeItem.Category,
		}
		_, err = client.Index().
			Index("awesome_items").
			Type("link").
			Id(awesomeItem.URL).
			BodyJson(body).
			Refresh(true).
			Do()
		return err
	}
	_, err = client.Index().
		Index("awesome_items").
		Type("github_repo").
		Id(strconv.FormatInt(awesomeItem.ID, 10)).
		BodyJson(&awesomeItem).
		Refresh(true).
		Do()
	return err
}

func (a AwesomeItem) Flush() {
	client := db.GetDB()
	_, err := client.Flush().Index("awesome_items").Do()
	if err != nil {
		log.Println("error on flush db")
	}
}
