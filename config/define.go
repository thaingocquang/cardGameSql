package config

// Database ...
type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

// Jwt ...
type Jwt struct {
	SecretKey string
}

// ENV ...
type ENV struct {
	// AppPort ...
	AppPort string

	// Database ...
	Database Database

	// Jwt
	Jwt Jwt
}
