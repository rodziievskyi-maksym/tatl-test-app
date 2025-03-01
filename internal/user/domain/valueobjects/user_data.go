package valueobjects

type UserData struct {
	UserID uint `gorm:"primaryKey"`
	School string
}

func (u *UserData) TableName() string {
	return "user_data"
}
