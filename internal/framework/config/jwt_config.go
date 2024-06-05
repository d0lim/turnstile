package config

import "os"

type JwtConfig struct {
	AccessSecret  []byte
	RefreshSecret []byte
}

func NewJwtConfig() *JwtConfig {
	return &JwtConfig{
		AccessSecret:  []byte(os.Getenv("ACCESS_SECRET")),
		RefreshSecret: []byte(os.Getenv("REFRESH_SECRET")),
	}
}
