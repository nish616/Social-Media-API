package users

import (
	"nishin/social-media-API/common"
	"time"
)

type User struct {
	ID        uint64    `gorm:"primary_key"`
	Username  string    `gorm:"column:username"`
	Email     string    `gorm:"column:email;unique_index"`
	CreatedAt time.Time `gorm:"column:createdAt;not_null"`
	UpdatedAt time.Time `gorm:"column:updatedAt;not_null"`
}

func AutoMigrate() {
	db := common.GetDB()

	db.AutoMigrate(&User{})

}

func (user *User) CreateUser() {

	db := common.GetDB()

	db.Create(&user)

}

func (user *User) GetUser(id uint64) int {
	db := common.GetDB()

	result := db.Find(&user, id)

	return int(result.RowsAffected)
}

func (user *User) UpdateUser() int {
	db := common.GetDB()

	result := db.Save(&user)

	return int(result.RowsAffected)
}

func (user *User) DeleteUser() int {
	db := common.GetDB()

	result := db.Delete(&user)

	return int(result.RowsAffected)
}
