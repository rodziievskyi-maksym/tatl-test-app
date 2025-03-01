package aggregates

import "github.com/rodziievskyi-maksym/tatl-test-app/internal/user/domain/valueobjects"

type User struct {
	ID       uint `gorm:"primaryKey"`
	Username string
	Profile  valueobjects.UserProfile `gorm:"foreignKey:UserID"`
	Data     valueobjects.UserData    `gorm:"foreignKey:UserID"`
}

func (u *User) TableName() string {
	return "user"
}
