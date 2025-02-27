package domain

type User struct {
	ID      uint `gorm:"primaryKey"`
	Name    string
	Profile UserProfile `gorm:"foreignKey:UserID"`
	Data    UserData    `gorm:"foreignKey:UserID"`
}

func (u *User) TableName() string {
	return "user"
}

type UserProfile struct {
	UserID    uint `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Phone     string
	Address   string
	City      string
}

func (u *UserProfile) TableName() string {
	return "user_profile"
}

type UserData struct {
	UserID uint `gorm:"primaryKey"`
	School string
}

func (u *UserData) TableName() string {
	return "user_data"
}
