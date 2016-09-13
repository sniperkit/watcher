package extractor

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/vsouza/watcher/models"
	"github.com/vsouza/watcher/utils"
)

type Throttler chan bool

func NewThrottler(n int) Throttler {
	return make(chan bool, n)
}

func (t Throttler) Fill() Throttler {
	t <- true
	return t
}

func (t Throttler) Drain() {
	<-t
}

var awsModel = new(models.AwesomeItem)

func ExtractLinks(node *goquery.Selection, categoryName string) {
	var wg sync.WaitGroup
	th := NewThrottler(2)
	node.NextFiltered("ul").Each(func(i int, s *goquery.Selection) {
		s.Find("li").Each(func(j int, t *goquery.Selection) {
			t.Find("a").Each(func(y int, u *goquery.Selection) {
				url, exists := u.Attr("href")
				if last := len(url) - 1; last >= 0 && url[last] == '/' {
					url = url[:last]
				}
				if exists {
					wg.Add(1)
					go func() {
						defer wg.Done()
						defer th.Fill().Drain()
						time.Sleep(1 * time.Millisecond) //to simulate work
						rep, err := getGithubData(url, categoryName)
						if err != nil {
							log.Println(err)
						}

						if rep != nil {
							if err := awsModel.SaveData(rep); err != nil {
								log.Println(err)
							}
						}

					}()
				}
			})
		})
	})
	wg.Wait()
	return
}

func getGithubData(url, category string) (*models.AwesomeItem, error) {
	if strings.Contains(url, "github.com") && !strings.Contains(url, "gist") {
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
		aws := models.AwesomeItem{
			Type:     "repo",
			Category: category,
		}
		if err := decoder.Decode(&aws); err != nil {
			return nil, err
		}

		data, err := GetPkgManagers(&aws)
		if err != nil {
			return nil, err
		}
		return data, err
	}
	return &models.AwesomeItem{Type: "url", URL: url, Category: category}, nil
}
