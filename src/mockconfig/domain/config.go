package domain

type (
	// ConfigTree encapsulates MockConfigs
	ConfigTree []*MockConfig

	// MockConfig holds a list of Responses
	MockConfig struct {
		Responses []Response
	}

	// Response structures the MatcherConfig and ResponseConfig for a single request and response
	Response struct {
		MatcherConfig struct {
			URI    string
			Method string
			Params struct {
				GET  map[string]string
				POST map[string]string
			}
		}
		ResponseConfig struct {
			StatusCode int
			Headers    map[string]string
			Body       interface{}
		}
	}
)
