package extractor

import (
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/sniperkit/watcher/models"
)

var catModel = new(models.Category)

func shouldIgnore(item string) bool {
	ignored := []string{
		"Contributing and License",
		"About",
		"How to Use",
		"Content",
	}
	set := make(map[string]struct{}, len(ignored))
	for _, s := range ignored {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func Runner(document *goquery.Document) {
	ProcessCategories(document)
}

func ProcessCategories(document *goquery.Document) {
	if document == nil {
		log.Fatal("cannot open document")
	}
	h1 := document.Find("h1")
	for i := range h1.Nodes {
		cat := models.Category{}
		cat.Name = h1.Eq(i).Text()
		if err := catModel.SaveData(&cat); err != nil {
			log.Printf("category error: %s \n", err)
		}
		if !shouldIgnore(cat.Name) {
			ExtractLinks(h1.Eq(i), cat.Name)
			GetMainSubCategories(h1.Eq(i))
		}
	}
}

func GetMainSubCategories(node *goquery.Selection) {
	h2 := node.NextFilteredUntil("h2", "h1")
	for i := range h2.Nodes {
		cat := models.Category{}
		cat.Name = h2.Eq(i).Text()
		if err := catModel.SaveData(&cat); err != nil {
			log.Printf("main sub category error: %s \n", err)
		}
		cat.MainParent = node.Text()
		if !shouldIgnore(cat.Name) {
			ExtractLinks(h2.Eq(i), cat.Name)
			GetSubCategories(h2.Eq(i))
		}
	}
}

func GetSubCategories(node *goquery.Selection) {
	h4 := node.NextFilteredUntil("h4", "h2")
	for i := range h4.Nodes {
		cat := models.Category{}
		cat.Name = h4.Eq(i).Text()
		cat.Parent = node.Text()
		if err := catModel.SaveData(&cat); err != nil {
			log.Printf("sub category error: %s \n", err)
		}
		if !shouldIgnore(cat.Name) {
			ExtractLinks(h4.Eq(i), cat.Name)
		}
	}
}
