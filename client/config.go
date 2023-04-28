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

func getSecretValue(secret string) (string, error) {
	// TODO: mb change it.
	return secret, nil
}

func isCfgValid(cfg *Config) bool {
	if cfg == nil {
		return false
	}

	return cfg.GetApiUrl() != "" // TODO: mb change it.
}
