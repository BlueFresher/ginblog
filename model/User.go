package model

import (
	"encoding/base64"
	"ginblog/utils/errmsg"
	"log"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(20);not null " json:"password" validate:"required,min=6,max=20" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色"`
	// Avatar string
}

// 查询用户是否存在
func CheckUser(username string) int {
	var user User
	db.Select("id").Where("username = ?", username).First(&user)
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USE
	}
	return errmsg.SUCCESS
}

// 更新查询
func CheckUpdateUser(id int, username string) int {
	var user User
	db.Select("id, username").Where("username = ?", username).First(&user)
	if user.ID == uint(id) {
		return errmsg.SUCCESS
	}
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USE
	}
	return errmsg.SUCCESS
}

//　新增用户
func CreateUser(data *User) int {
	data.Password = ScryptPw(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询单个用户
func GetUser(id int) (User, int) {
	var user User
	err = db.Where("ID = ?", id).First(&user).Error
	if err != nil {
		return user, errmsg.ERROR
	}
	return user, errmsg.SUCCESS
}

func GetUsers(username string, pagSize int, pagNum int) ([]User, int) {
	var users []User
	var total int

	if username != "" {
		db.Select("id, username, role").Where("username LIKE ?", username+"%").Find(&users).Count(&total).Limit(pagSize).Offset((pagNum - 1) * pagSize)
		return users, total
	}
	db.Select("id, username, role").Find(&users).Count(&total).Limit(pagSize).Offset((pagNum - 1) * pagSize)
	if err == gorm.ErrRecordNotFound {
		return nil, 0
	}
	return users, total
}

func DeleteUser(id int) int {
	var user User
	err = db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func EditUser(id int, data *User) int {
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err = db.Model(&User{}).Where("id=?", id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 钩子函数
// func (u *User) BeforeSave() {
// 	u.Password = ScryptPw(u.Password)
// }

func ScryptPw(password string) string {
	const KeyLen = 10
	// salt := make([]byte, 8)
	salt := []byte{12, 21, 4, 6, 66, 22, 222, 11}
	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}

func CheckLogin(username, password string) int {
	var user User
	db.Where("username=?", username).First(&user)
	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}

	if ScryptPw(password) != user.Password {
		return errmsg.ERROR_PASSWORD_ERROR
	}
	if user.Role != 1 {
		return errmsg.ERROR_USER_NO_RIGHT
	}
	return errmsg.SUCCESS
}

// CheckLoginFront 前台登录
func CheckLoginFront(username string, password string) (User, int) {
	var user User

	db.Where("username = ?", username).First(&user)

	if ScryptPw(password) != user.Password {
		return user, errmsg.ERROR_PASSWORD_ERROR
	}
	if user.Role != 1 {
		return user, errmsg.ERROR_USER_NO_RIGHT
	}
	return user, errmsg.SUCCESS
}
