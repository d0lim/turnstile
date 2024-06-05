package config

import "os"

type JwtConfig struct {
	AccessSecret  string
	RefreshSecret string
}

func NewJwtConfig() *JwtConfig {
	return &JwtConfig{
		AccessSecret:  os.Getenv("ACCESS_SECRET"),
		RefreshSecret: os.Getenv("REFRESH_SECRET"),
	}
}
