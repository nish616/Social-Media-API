package users

import (
	"nishin/social-media-API/common"
	"time"
)

type User struct {
	ID        uint32    `gorm:"primary_key"`
	Username  string    `gorm:"column:username"`
	Email     string    `gorm:"column:email;unique_index"`
	CreatedAt time.Time `gorm:"column:createdAt;not_null"`
	UpdatedAt time.Time `gorm:"column:updatedAt;not_null"`
}

func AutoMigrate() {
	db := common.GetDB()

	db.AutoMigrate(&User{})

}

func (user *User) Create() {

	db := common.GetDB()

	db.Create(&user)

}
