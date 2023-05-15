package Model

import "log"

type User struct {
	UId       int    `gorm:"primary_key;type:bigint;column:uid" json:"userId"`
	Name      string `gorm:"NOT NULL UNIQUE" json:"userName"`
	Password  string `gorm:"NOT NULL" json:"password"`
	UserEmail string `json:"email"`
}

func (user *User) CheckUser() bool {
	if user.Name == "" || user.Password == "" {
		return false
	}
	result := GlobalConn.Where(&User{Name: user.Name, Password: user.Password}).Find(user)
	if result.Error != nil || result.RowsAffected == 0 {
		return false
	}
	return true
}

func (user *User) CheckUserExist() bool {
	if user.Name == "" {
		return false
	}
	result := GlobalConn.Where(&User{Name: user.Name}).Find(user)
	if result.Error != nil || result.RowsAffected == 0 {
		return true
	}
	return false
}

func (user *User) RegisterUser() {
	GlobalConn.Table("users").Create(user)
}

func (user *User) ChangePassword(newPassword string) error {
	result := GlobalConn.Model(&User{}).Where(&User{Name: user.Name}).Update("password", newPassword)
	return result.Error
}

func (user *User) GetUserByName(UserName string) {
	result := GlobalConn.Table("users").Where("name=?", user.Name).First(user)
	if result.Error != nil {
		log.Println(result.Error)
	}
}
