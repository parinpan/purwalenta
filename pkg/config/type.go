package config

type Config struct {
	App AppConfig
}

type AppConfig struct {
	Name    string
	Version string
}
