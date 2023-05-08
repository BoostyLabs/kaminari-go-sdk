package client

type Config struct {
	ApiKey string
	ApiUrl string
}

func (c Config) GetApiKey() string {
	return c.ApiKey
}

func (c Config) GetApiUrl() string {
	return c.ApiUrl
}

func isCfgValid(cfg *Config) bool {
	if cfg == nil {
		return false
	}

	return cfg.GetApiUrl() != "" // TODO: mb change it.
}