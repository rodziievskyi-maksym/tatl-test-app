package entities

type Auth struct {
	ID     int
	ApiKey string
}

func (a *Auth) TableName() string {
	return "auth"
}
