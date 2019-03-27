package domain

type (
	ConfigProvider interface {
		GetConfigTree() ConfigTree
		LoadConfig(path string) ConfigTree
	}
)
