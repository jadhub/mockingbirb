package dto

type (
	// ConfigTree encapsulates MockConfigs
	ConfigTree []*MockConfig

	// MockConfig holds a list of Responses
	MockConfig struct {
		Responses []Response `json:"responses"`
	}

	// Response structures the MatcherConfig and ResponseConfig for a single request and response
	Response struct {
		MatcherConfig struct {
			URI    string `json:"uri"`
			Method string `json:"method"`
			Params struct {
				GET  map[string]string `json:"GET,omitempty"`
				POST map[string]string `json:"POST,omitempty"`
			} `json:"params,omitempty"`
		} `json:"matcherconfig"`
		ResponseConfig struct {
			StatusCode int               `json:"statusCode"`
			Headers    map[string]string `json:"headers"`
			Body       interface{}       `json:"body,omitempty"`
		} `json:"responseconfig"`
	}
)
