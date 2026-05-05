package config

type Config struct {
	HTTPAddr       string
	AppriseBaseUrl string
	AppriseKey     string
}

func Load() Config {
	return Config{
		HTTPAddr:       ":8080",
		AppriseBaseUrl: "http://localhost:8000",
		AppriseKey:     "myapp",
	}
}
