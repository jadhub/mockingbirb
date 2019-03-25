package domain

type (
	ConfigTree []*MockConfig

	MockConfig struct {
		Route     string     `json:"route"`
		Methods   []string   `json:"methods"`
		Responses []Response `json:"responses"`
	}

	Response struct {
		URI        string      `json:"uri"`
		Method     string      `json:"method"`
		StatusCode int         `json:"statusCode"`
		Headers    Headers     `json:"headers"`
		Body       interface{} `json:"body,omitempty"`
	}

	Headers struct {
		ContentType string `json:"Content-Type"`
	}
)
