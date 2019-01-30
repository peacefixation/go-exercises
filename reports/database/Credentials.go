package database

// Credentials the database credentials
type Credentials struct {
	User     string
	Password string
	Name     string
}

// NewCredentials create a new credentials instance
func NewCredentials(user, password, name string) Credentials {
	return Credentials{
		User:     user,
		Password: password,
		Name:     name,
	}
}

// SampleCredentials create sample credentials
func SampleCredentials() Credentials {
	return NewCredentials("TODO", "TODO", "TODO")
}
