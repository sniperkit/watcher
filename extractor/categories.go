package extractor

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/vsouza/watcher/models"
)

var catModel = new(models.Category)

func Runner(document *goquery.Document) {
	ProcessCategories(document)
}

func ProcessCategories(document *goquery.Document) {
	h1 := document.Find("h1")
	for i := range h1.Nodes {
		cat := models.Category{}
		cat.Name = h1.Eq(i).Text()
		catModel.SaveData(&cat)
		ExtractLinks(h1.Eq(1), cat.Name)
		GetSubCategories(h1.Eq(i))
	}
	catModel.Flush()
}

func GetSubCategories(node *goquery.Selection) {
	h3 := node.NextFilteredUntil("h3", "h1")
	for i := range h3.Nodes {
		cat := models.Category{}
		cat.Name = h3.Eq(i).Text()
		cat.Parent = node.Text()
		ExtractLinks(h3.Eq(i), cat.Name)
		GetNestedCategories(node, h3.Eq(i))
		GetNestedSubCategories(node, h3.Eq(i))
	}
}

func GetNestedCategories(mainNode, node *goquery.Selection) {
	h4 := node.NextFilteredUntil("h4", "h3")
	for i := range h4.Nodes {
		cat := models.Category{}
		cat.Name = h4.Eq(i).Text()
		cat.Parent = node.Text()
		cat.MainParent = mainNode.Text()
		ExtractLinks(h4.Eq(1), cat.Name)
	}
}

func GetNestedSubCategories(mainNode, node *goquery.Selection) {
	h5 := node.NextFilteredUntil("h5", "h3")
	for i := range h5.Nodes {
		cat := models.Category{}
		cat.Name = h5.Eq(i).Text()
		cat.Parent = node.Text()
		cat.MainParent = mainNode.Text()
		ExtractLinks(h5.Eq(1), cat.Name)
	}
}
