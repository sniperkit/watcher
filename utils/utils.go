package utils

import (
	"bytes"
	b64 "encoding/base64"
	"io"
	"net/http"

	"github.com/sniperkit/watcher/config"
)

func DecodeBase64(content string) ([]byte, error) {
	raw, err := b64.StdEncoding.DecodeString(content)
	if err != nil {
		return nil, err
	}
	return raw, nil
}

func DoReq(url string) (io.ReadCloser, error) {
	cfg := config.GetConfig()
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", cfg.GetString("github.auth.app"))
	req.SetBasicAuth("vsouza", cfg.GetString("github.auth.token"))
	resp, err := client.Do(req)
	return resp.Body, err
}

func MountGHURL(owner, repoName string) string {
	var url bytes.Buffer
	url.WriteString("https://api.github.com/repos/")
	url.WriteString(owner)
	url.WriteString("/")
	url.WriteString(repoName)
	url.WriteString("/git/trees/HEAD")
	return url.String()
}
