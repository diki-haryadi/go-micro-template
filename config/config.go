package config

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
	"path/filepath"
	"runtime"
)

type Config struct {
	App AppConfig
}

var BaseConfig *Config

type AppConfig struct {
	AppEnv      string `json:"app_env" envconfig:"APP_ENV"`
	AppName     string `json:"app_name" envconfig:"APP_NAME"`
	ConfigOauth ConfigOauth
}

// Config stores all configuration options
type ConfigOauth struct {
	Oauth         OauthConfig
	Session       SessionConfig
	IsDevelopment bool
}

// OauthConfig stores oauth service configuration options
type OauthConfig struct {
	AccessTokenLifetime  int `json:"access_token_lifetime" envconfig:"OAUTH_ACCESS_TOKEN_LIFETIME"`
	RefreshTokenLifetime int `json:"refresh_token_lifetime" envconfig:"OAUTH_REFRESH_TOKEN_LIFETIME"`
	AuthCodeLifetime     int `json:"auth_code_lifetime" envconfig:"OAUTH_AUTH_CODE_LIFETIME"`
}

// SessionConfig stores session configuration for the web app
type SessionConfig struct {
	Secret string `json:"secret" envconfig:"SESSION_SECRET"`
	Path   string `json:"path" envconfig:"SESSION_PATH"`
	// MaxAge=0 means no 'Max-Age' attribute specified.
	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'.
	// MaxAge>0 means Max-Age attribute present and given in seconds.
	MaxAge int `json:"max_age" envconfig:"SESSION_MAX_AGE"`
	// When you tag a cookie with the HttpOnly flag, it tells the browser that
	// this particular cookie should only be accessed by the server.
	// Any attempt to access the cookie from client script is strictly forbidden.
	HTTPOnly bool `json:"http_only" envconfig:"SESSION_HTTP_ONLY"`
}

func LoadConfig() *Config {
	_, callerDir, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Error generating env dir")
	}

	// Define the possible paths to the .env file
	envPaths := []string{
		filepath.Join(filepath.Dir(callerDir), "..", "envs/.env"),
	}
	_ = godotenv.Overload(envPaths[0])
	var configLoader Config

	if err := envconfig.Process("BaseConfig", &configLoader); err != nil {
		log.Printf("error load config: %v", err)
	}

	BaseConfig = &configLoader
	spew.Dump(configLoader)
	return &configLoader
}
