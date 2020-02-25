package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	ID          uint32  `gorm:"primary_key;AUTO_INCREMENT" json:"user_id"`
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	Email       string  `json:"email"`
	PhoneNumber *string `json:"phone_number"`
	Password    string  `json:"password"`
	//CreatedAt   time.Time  `gorm:"DEFAULT:now()" json:"created_at"`
	//UpdatedAt   *time.Time `gorm:"DEFAULT:now()" json:"updated_at"`
	DeletedAt *time.Time `gorm:"DEFAULT:NULL" json:"deleted_at"`
}

func (user *User) Create() error {
	return DB.Create(user).Error
}

func (user *User) GetByID() User {
	userModel := User{}
	DB.Where("id = ?", user.ID).Find(&userModel)
	return userModel
}

func (user *User) Updates(setting map[string]interface{}) error {
	if err := DB.Model(&user).Where("id = ?", user.ID).Updates(setting).Error; err != nil {
		return err
	}
	return nil
}

func (user *User) Delete() error {
	userModel := User{}
	DB.Where("id = ?", user.ID).Find(&userModel)

	if err := DB.Delete(&userModel).Error; err != nil {
		return err
	}

	return nil
}

func (user *User) UnDelete() error {
	setting := map[string]interface{}{
		"deleted_at": gorm.Expr("null"),
	}

	if err := DB.Table("users").Where("id = ?", user.ID).Update(setting).Error; err != nil {
		return err
	}

	return nil
}

func (user *User) DeleteByName() error {
	return DB.Where("first_name = ?", &user.FirstName).Delete(&user).Error
}

func (user *User) GetAll() ([]User, error) {
	userModel := []User{}
	if err := DB.Find(&userModel).Error; err != nil {
		return nil, err
	}

	return userModel, nil
}
