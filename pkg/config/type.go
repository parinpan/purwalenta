package config

type Config struct {
	App                AppConfig
	Database           DatabaseConfig
	UserAuthentication UserAuthenticationConfig
}

type AppConfig struct {
	Name    string
	Version string
}

type DatabaseConfig struct {
	Postgre PostgreConfig
}

type PostgreConfig struct {
	DSN         string
	MaxLifeTime int
	MaxIdle     int
	MaxOpen     int
}

type UserAuthenticationConfig struct {
	SecretToken string
	MaxLifeTime int
}
