package extractor

import (
	"encoding/json"
	"strings"

	"github.com/vsouza/watcher/utils"
)

func GetPkgManagers(repo *Repo) (*Repo, error) {

	req, err := utils.DoReq(utils.MountGHURL(repo.Owner.Login, repo.Name))
	defer req.Close()
	if err != nil {
		return nil, err
	}

	var c Content
	decoder := json.NewDecoder(req)
	if err := decoder.Decode(&c); err != nil {
		return nil, err
	}

	for _, node := range c.Tree {
		if node.Path == "Package.swift" {
			repo.PackageManagers.SPM = true
		}
		if strings.Contains(node.Path, "podspec") {
			repo.PackageManagers.CocoaPods = true
		}
		if node.Path == "Cartfile" {
			repo.PackageManagers.Carthage = true
		}
	}

	return repo, nil
}
