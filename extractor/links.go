package extractor

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/vsouza/watcher/utils"
)

type Repo struct {
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
		Carthage  bool `bson:"carthage"`
		SPM       bool `bson:"spm"`
		CocoaPods bool `bson:"cocoa_pods"`
	} `json:"package_managers" bson:"package_managers,omitempty"`
}

var throttle = make(chan int, 50)

func ExtractLinks(node *goquery.Selection, categoryName string) {
	var wg sync.WaitGroup
	node.NextFiltered("ul").Each(func(i int, s *goquery.Selection) {
		s.Find("li").Each(func(j int, t *goquery.Selection) {
			t.Find("a").Each(func(y int, u *goquery.Selection) {
				url, exists := u.Attr("href")
				if exists {
					wg.Add(1)
					throttle <- 1
					go f(url, categoryName, &wg, throttle)
				}
			})
		})
	})
	wg.Wait()
}

func f(url string, categoryName string, wg *sync.WaitGroup, throttle chan int) {
	defer wg.Done()
	rep, err := getGithubData(url, categoryName)
	if err != nil {
		log.Println(err)
	}
	if rep != nil {
		// sendo to model
		// go processRepoData(rep)
		fmt.Println(rep)
	}
	<-throttle
}

func getGithubData(url, category string) (*Repo, error) {

	if !strings.Contains(url, "github.com") || strings.Contains(url, "gist") {
		return &Repo{Type: "url", URL: url, Category: category}, nil
	}

	stringSlice := strings.Split(url, "/")
	if len(stringSlice) != 5 {
		return nil, errors.New("item url size error")
	}

	var buffer bytes.Buffer
	buffer.WriteString("https://api.github.com/repos/")
	buffer.WriteString(stringSlice[3])
	buffer.WriteString("/")
	buffer.WriteString(stringSlice[4])

	req, err := utils.DoReq(buffer.String())
	defer req.Close()
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(req)
	rep := Repo{
		Type:     "repo",
		Category: category,
	}
	if err := decoder.Decode(&rep); err != nil {
		return nil, err
	}

	data, err := GetPkgManagers(&rep)
	if err != nil {
		return nil, err
	}
	return data, err
}
