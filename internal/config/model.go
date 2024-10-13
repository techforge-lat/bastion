package config

type Root struct {
	ServerPort     uint
	Env            string
	AllowedDomains string
	AllowedMethods string
	Database       Database
}

type Database struct {
	Driver   string
	Host     string
	Port     uint
	User     string
	Password string
	Name     string
	SSLMode  string
}
