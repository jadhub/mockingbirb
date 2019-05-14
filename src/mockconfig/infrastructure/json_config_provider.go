package infrastructure

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	configDomain "mockingbirb/src/mockconfig/domain"
	configDto "mockingbirb/src/mockconfig/infrastructure/dto"
)

type (
	// JSONConfigProvider defines the configProvider implementation
	JSONConfigProvider struct {
		cfg configDomain.ConfigTree
	}
)

const (
	// CONFIGDIR ses the directory to read mock config from
	CONFIGDIR = "mock_config"
)

var (
	// check interface implementation
	_ configDomain.ConfigProvider = (*JSONConfigProvider)(nil)
)

// NewJSONConfigProvider returns a JSONConfigProvider
func NewJSONConfigProvider() *JSONConfigProvider {
	p := new(JSONConfigProvider)

	currentDir, err := os.Getwd()
	if err != nil {
		panic("shit")
	}

	p.cfg = p.LoadConfig(currentDir + "/../" + CONFIGDIR)
	return p
}

// GetConfigTree returns the current ConfigTree
func (p *JSONConfigProvider) GetConfigTree() configDomain.ConfigTree {
	return p.cfg
}

// LoadConfig loads the Config and returns its ConfigTree
func (p *JSONConfigProvider) LoadConfig(path string) configDomain.ConfigTree {
	log.Print("loading config")

	var files []string

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil

	})

	if err != nil {
		panic(err)
	}

	res := configDto.ConfigTree{}

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

			fileTree := &configDto.ConfigTree{}

			err = json.Unmarshal(read, fileTree)
			if err != nil {
				panic(err)
			}

			for _, configEntry := range *fileTree {
				res = append(res, configEntry)
			}
		}
	}

	return p.mapDtoToConfig(res)
}

func (p *JSONConfigProvider) mapDtoToConfig(dtoConfig configDto.ConfigTree) configDomain.ConfigTree {
	res := configDomain.ConfigTree{}

	resMockConfig := &configDomain.MockConfig{
		Responses: []configDomain.Response{},
	}

	for _, dtoConfigPart := range dtoConfig {
		for _, response := range dtoConfigPart.Responses {
			resMockConfig.Responses = append(
				resMockConfig.Responses,
				configDomain.Response{
					MatcherConfig: struct {
						URI    string
						Method string
						Params struct {
							GET  map[string]string
							POST map[string]string
						}
					}{
						URI:    response.MatcherConfig.URI,
						Method: response.MatcherConfig.Method,
						Params: struct {
							GET  map[string]string
							POST map[string]string
						}{
							GET:  response.MatcherConfig.Params.GET,
							POST: response.MatcherConfig.Params.POST,
						},
					},
					ResponseConfig: struct {
						StatusCode int
						Headers    map[string]string
						Body       interface{}
					}{
						StatusCode: response.ResponseConfig.StatusCode,
						Headers:    response.ResponseConfig.Headers,
						Body:       response.ResponseConfig.Body,
					},
				},
			)
		}
	}

	res = append(res, resMockConfig)

	return res
}
