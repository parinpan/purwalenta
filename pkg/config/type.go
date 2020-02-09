package config

type Config struct {
	App                AppConfig
	Database           DatabaseConfig
	UserAuthentication UserAuthenticationConfig
	SMTP               SMTPConfig
}

type AppConfig struct {
	Name             string
	Version          string
	SignUpEmailAgent string
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

type SMTPConfig struct {
	Identity string
	Username string
	Password string
	Outgoing OutgoingSMTPConfig
	Incoming IncomingSMTPConfig
}

type OutgoingSMTPConfig struct {
	Server string
	Port   int
}

type IncomingSMTPConfig struct {
	Server string
	Port   int
}
