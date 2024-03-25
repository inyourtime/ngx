package util

import (
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
)

const (
	EnvProduction  = "production"
	EnvDevelopment = "development"
)

type Config struct {
	Environment string `mapstructure:"ENVIRONMENT"`

	PostgresSource string `mapstructure:"POSTGRES_SOURCE"`

	ServerType string `mapstructure:"SERVER_TYPE"`
	ServerPort string `mapstructure:"SERVER_PORT"`

	LogType string `mapstructure:"LOG_TYPE"`
	DBType  string `mapstructure:"DB_TYPE"`

	TokenSymmetricKey string `mapstructure:"TOKEN_SYMMETRIC_KEY"`

	TestRepo string `mapstructure:"TEST_REPO"`

	GoogleClientID     string `mapstructure:"GOOGLE_CLIENT_ID"`
	GoogleClientSecret string `mapstructure:"GOOGLE_CLIENT_SECRET"`
	GoogleRedirectURL  string `mapstructure:"GOOGLE_REDIRECT_URL"`
	GoogleUserInfoURL  string `mapstructure:"GOOGLE_USER_INFO_URL"`

	GithubClientID     string `mapstructure:"GITHUB_CLIENT_ID"`
	GithubClientSecret string `mapstructure:"GITHUB_CLIENT_SECRET"`
	GithubRedirectURL  string `mapstructure:"GITHUB_REDIRECT_URL"`
	GithubUserInfoURL  string `mapstructure:"GITHUB_USER_INFO_URL"`
}

func (c Config) IsProduction() bool {
	return c.Environment == EnvProduction
}

func (c Config) IsTestAllRepo() bool {
	return c.TestRepo == "all"
}

func (c Config) GetGoogleCfg() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     c.GoogleClientID,
		ClientSecret: c.GoogleClientSecret,
		RedirectURL:  c.GoogleRedirectURL,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
}

func (c Config) GetGithubCfg() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     c.GithubClientID,
		ClientSecret: c.GithubClientSecret,
		RedirectURL:  c.GithubRedirectURL,
		Scopes:       []string{"user:email"},
		Endpoint:     github.Endpoint,
	}
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
