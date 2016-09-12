package extractor

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/vsouza/watcher/models"
)

var catModel = new(models.Categories)

type category struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name"`
	UpdatedAt  int64  `json:"updated_at,omitempty"`
	Parent     string `json:"parent,omitempty"`
	MainParent string `json:"main_parent,omitempty"`
}

func Runner(document *goquery.Document) {
	ProcessCategories(document)
}

func ProcessCategories(document *goquery.Document) {
	h1 := document.Find("h1")
	for i := range h1.Nodes {
		cat := category{}
		cat.Name = h1.Eq(i).Text()
		fmt.Printf("- %s \n", cat.Name)
		ExtractLinks(h1.Eq(1), cat.Name)
		GetSubCategories(h1.Eq(i))
	}
}

func GetSubCategories(node *goquery.Selection) {
	h3 := node.NextFilteredUntil("h3", "h1")
	for i := range h3.Nodes {
		cat := category{}
		cat.Name = h3.Eq(i).Text()
		cat.Parent = node.Text()
		fmt.Printf("-- %s \n", cat.Name)
		ExtractLinks(h3.Eq(i), cat.Name)
		GetNestedCategories(node, h3.Eq(i))
		GetNestedSubCategories(node, h3.Eq(i))
	}
}

func GetNestedCategories(mainNode, node *goquery.Selection) {
	h4 := node.NextFilteredUntil("h4", "h3")
	for i := range h4.Nodes {
		cat := category{}
		cat.Name = h4.Eq(i).Text()
		cat.Parent = node.Text()
		cat.MainParent = mainNode.Text()
		ExtractLinks(h4.Eq(1), cat.Name)
		fmt.Printf("--- %s \n", cat.Name)
	}
}

func GetNestedSubCategories(mainNode, node *goquery.Selection) {
	h5 := node.NextFilteredUntil("h5", "h3")
	for i := range h5.Nodes {
		cat := category{}
		cat.Name = h5.Eq(i).Text()
		cat.Parent = node.Text()
		cat.MainParent = mainNode.Text()
		ExtractLinks(h5.Eq(1), cat.Name)
		fmt.Printf("---- %s \n", cat.Name)
	}
}
