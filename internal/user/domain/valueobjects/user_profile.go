package valueobjects

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
