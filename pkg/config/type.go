package config

type Config struct {
	App      AppConfig
	Database DatabaseConfig
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
