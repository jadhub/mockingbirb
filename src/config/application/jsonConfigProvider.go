package application

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"go.aoe.com/mockingbirb/src/config/domain"
)

type (
	JsonConfigProvider struct{}
)

const (
	CONFIGDIR = "config/mockconfig"
)

func (p *JsonConfigProvider) GetConfigTree() domain.ConfigTree {
	return p.loadConfig(CONFIGDIR)
}

func (p *JsonConfigProvider) loadConfig(path string) domain.ConfigTree {
	var files []string

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})

	if err != nil {
		panic(err)
	}

	res := domain.ConfigTree{}

	for _, file := range files {
		fileInfo, err := os.Stat(file)
		if err != nil {
			panic(err)
		}

		switch mode := fileInfo.Mode(); {
		case mode.IsDir():
			continue
		case mode.IsRegular():
			read, err := ioutil.ReadFile(file)
			if err != nil {
				panic(err)
			}

			fileTree := &domain.ConfigTree{}

			err = json.Unmarshal(read, fileTree)
			if err != nil {
				panic(err)
			}

			for _, configEntry := range *fileTree {
				res = append(res, configEntry)
			}
		}
	}

	return res
}
