package extractor

import (
	"encoding/json"
	"strings"

	"github.com/vsouza/watcher/document"
	"github.com/vsouza/watcher/models"
	"github.com/vsouza/watcher/utils"
)

func GetPkgManagers(aws *models.AwesomeItem) (*models.AwesomeItem, error) {

	req, err := utils.DoReq(utils.MountGHURL(aws.Owner.Login, aws.Name))
	defer req.Close()
	if err != nil {
		return nil, err
	}

	var c document.Content
	decoder := json.NewDecoder(req)
	if err := decoder.Decode(&c); err != nil {
		return nil, err
	}

	for _, node := range c.Tree {
		if node.Path == "Package.swift" {
			aws.PackageManagers.SPM = true
		}
		if strings.Contains(node.Path, "podspec") {
			aws.PackageManagers.CocoaPods = true
		}
		if node.Path == "Cartfile" {
			aws.PackageManagers.Carthage = true
		}
	}

	return aws, nil
}
