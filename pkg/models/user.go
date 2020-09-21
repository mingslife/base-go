package models

import (
	"github.com/jinzhu/gorm"

	"base-go/pkg/utils"
)

type User struct {
	gorm.Model

	Username string `json:"username"`
	Password string `json:"password"`
}

var (
	usersFields = []string{"id", "username"}
	userFields  = []string{"id", "username"}
)

func GetUsers(limit, page int) (s []*User) {
	DB.Select(usersFields).
		Order("created_at").Limit(limit).Offset(Offset(limit, page)).
		Find(&s)
	return
}

func CountUsers() (count int) {
	DB.Model(&User{}).Count(&count)
	return
}

func GetUser(id uint) *User {
	var v User
	if err := DB.Select(userFields).First(&v, "id = ?", id).Error; err != nil {
		return nil
	}
	return &v
}

func (v *User) Save() error {
	v.Password = utils.GenerateFromPassword(v.Password)
	return DB.Save(v).Error
}

func (v *User) Update() error {
	return DB.Model(v).Update(map[string]interface{}{
		"username": v.Username,
		"password": utils.GenerateFromPassword(v.Password),
	}).Error
}

func (v *User) Delete() error {
	return DB.Delete(v).Error
}

func GetUserByUsernameAndPassword(username, password string) *User {
	var v User
	if err := DB.Select([]string{"id", "username", "password"}).First(&v, "username = ?", username).Error; err != nil {
		return nil
	}
	if !utils.MatchPassword(v.Password, password) {
		return nil
	}
	return &v
}
