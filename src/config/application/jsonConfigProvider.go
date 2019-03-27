package application

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"go.aoe.com/mockingbirb/src/config/domain"
)

type (
	JsonConfigProvider struct {
		cfg domain.ConfigTree
	}
)

const (
	CONFIGDIR = "config/json_config"
)

func NewJsonConfigProvider() *JsonConfigProvider {
	p := new(JsonConfigProvider)
	p.cfg = p.LoadConfig(CONFIGDIR)
	return p
}

func (p *JsonConfigProvider) GetConfigTree() domain.ConfigTree {
	return p.cfg
}

func (p *JsonConfigProvider) LoadConfig(path string) domain.ConfigTree {
	log.Print("loading config")

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

		// skip non-json files here
		if !strings.HasSuffix(fileInfo.Name(), ".json") {
			continue
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
