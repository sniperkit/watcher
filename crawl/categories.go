package crawl

import (
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/vsouza/watcher/models"
)

var catModel = new(models.Categories)

func Runner(document *goquery.Document) {
	var wg sync.WaitGroup

	wg.Add(4)
	go GetMajorCategories(document, &wg)
	go GetSubCategories(document, &wg)
	go GetNestedCategories(document, &wg)
	go GetNestedSubCategories(document, &wg)

	wg.Wait()
	// DONE
}

func GetMajorCategories(document *goquery.Document, wg *sync.WaitGroup) error {
	var err error
	defer wg.Done()
	h1 := document.Find("h1")
	for i := range h1.Nodes {
		cat := category{}
		cat.Name = h1.Eq(i).Text()
		// SAVE DATA
		go catModel.SaveData(cat)
		// Extract Links
	}
}

func GetSubCategories(document *goquery.Document, wg *sync.WaitGroup) error {
	var err error
	defer wg.Done()
	h1 := document.Find("h1")
	for i := range h1.Nodes {
		cat := category{}
		cat.Name = h1.Eq(i).Text()
		// SAVE DATA
		go catModel.SaveData(cat)

		// Extract Links
	}
}

func GetNestedCategories(document *goquery.Document, wg *sync.WaitGroup) error {
	var err error
	defer wg.Done()
	h1 := document.Find("h1")
	for i := range h1.Nodes {
		cat := category{}
		cat.Name = h1.Eq(i).Text()
		// SAVE DATA
		go catModel.SaveData(cat)

		// Extract Links
	}
}

func GetNestedSubCategories(document *goquery.Document, wg *sync.WaitGroup) error {
	var err error
	defer wg.Done()
	h1 := document.Find("h1")
	for i := range h1.Nodes {
		cat := category{}
		cat.Name = h1.Eq(i).Text()
		// SAVE DATA
		go catModel.SaveData(cat)

		// Extract Links
	}
}
