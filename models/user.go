package models

import (
	"onecho_sso_backend/pkg/enums"
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	CreatedAt time.Time
	UpdatedAt time.Time

	UID       string       `gorm:"column:uid;primaryKey"`
	Phone     string       `gorm:"column:phone"`
	Email     string       `gorm:"column:email"`
	Password  string       `gorm:"column:password"`
	NickName  string       `gorm:"column:nick_name;unique"`
	Name      string       `gorm:"column:name"`
	JoinTime  time.Time    `gorm:"column:join_time"`
	AvatarURL string       `gorm:"column:avatar_url"`
	Gender    enums.Gender `gorm:"column:gender"`

	LarkID string `gorm:"column:lark_id"`
}

func (user User) TableName() string {
	return "sso_user"
}

func GetUser(user *User) (*User, error) {
	result := new(User)
	if err := db.Where(user).First(result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return result, nil
}

func GetUserByNickName(name string) (*User, error) {
	return GetUser(&User{NickName: name})
}

func GetUserByEmail(email string) (*User, error) {
	return GetUser(&User{
		Email: email,
	})
}

func AddUser(user *User) error {
	return db.Create(user).Error
}
