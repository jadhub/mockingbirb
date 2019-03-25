package domain

type (
	ConfigProvider interface {
		GetConfigTree() ConfigTree
	}
)
