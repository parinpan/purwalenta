package config

type Config struct {
	App                AppConfig
	Database           DatabaseConfig
	UserAuthentication UserAuthenticationConfig
	Assessment         AssessmentConfig
	SMTP               SMTPConfig
	Oauth2             Oauth2Config
}

type AppConfig struct {
	Name             string
	Version          string
	SignUpEmailAgent string
}

type DatabaseConfig struct {
	Postgre PostgreConfig
	Redis   RedisConfig
}

type PostgreConfig struct {
	DSN         string
	MaxLifeTime int
	MaxIdle     int
	MaxOpen     int
}

type RedisConfig struct {
	Address  string
	Password string
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

type Oauth2Config struct {
	Google GoogleOauth2Config
}

type GoogleOauth2Config struct {
	ClientID     string
	ClientSecret string
	CallbackURL  string
	Scopes       []string
}

type AssessmentConfig struct {
	PersonalityQuestionsJSON string
}
