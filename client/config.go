package client

type Config struct {
	ApiKey    string
	SecretKey string
	ApiUrl    string
}

func isCfgValid(cfg *Config) bool {
	if cfg == nil {
		return false
	}

	return cfg.ApiKey != "" && cfg.ApiUrl != ""
}
