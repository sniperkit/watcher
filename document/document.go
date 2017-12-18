package document

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday"
	"github.com/sniperkit/watcher/config"
	"github.com/sniperkit/watcher/utils"
)

func Init() (*goquery.Document, error) {
	cfg := config.GetConfig()
	var err error
	var document *goquery.Document
	Content, err := getRemoteContent()
	if err != nil {
		log.Printf("here %s", err)
		return nil, err
	}
	for _, node := range Content.Tree {
		if node.Path != cfg.GetString("github.repo.readme") {
			continue
		}
		blob, err := getBlob(node.URL)
		if err != nil {

			return nil, err
		}
		decodedStr, err := utils.DecodeBase64(blob.Content)
		if err != nil {
			return nil, err
		}
		unsafe := blackfriday.MarkdownCommon(decodedStr)
		document, err = goquery.NewDocumentFromReader(bytes.NewReader(unsafe))
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	return document, nil
}

type Content struct {
	Tree []struct {
		Path string `json:"path"`
		URL  string `json:"url"`
	} `json:"tree"`
}

func getRemoteContent() (*Content, error) {
	cfg := config.GetConfig()
	url := cfg.GetString("github.repo.url")
	r, err := utils.DoReq(url)
	defer r.Close()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var c Content
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(&c); err != nil {
		return nil, err
	}
	return &c, nil
}

type blob struct {
	URL     string `json:"url"`
	Content string `json:"Content"`
}

func getBlob(url string) (*blob, error) {
	var b blob
	nodeReq, err := utils.DoReq(url)
	defer nodeReq.Close()
	if err != nil {
		return nil, err
	}
	decoderBlob := json.NewDecoder(nodeReq)
	if err = decoderBlob.Decode(&b); err != nil {
		return nil, err
	}
	return &b, nil
}
