package config

// DATABASE type contains database connection values
type Database struct {
	Engine   string
	HOST     string
	Port     string
	Username string
	Password string
	DBName   string
}
