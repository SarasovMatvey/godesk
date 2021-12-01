package client

const DefaultBaseUrl = "https://api.chat2desk.com/v1/"

type Config struct {
	ApiToken string
	BaseUrl  string
}

func NewConfig(apiToken string, baseUrl string) Config {
	config := Config{
		ApiToken: apiToken,
		BaseUrl:  baseUrl,
	}

	if baseUrl == "" {
		config.BaseUrl = DefaultBaseUrl
	}

	return config
}
