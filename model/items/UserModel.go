package items

import (
	"code/01/model"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	ID int `json:"id"`
	Name string `json:"name"`
	Nickname string `json:"nickname"`
	Status int `json:"status" default:"1"`
}

func InnertUser(user *User)(err error)  {

	if err = model.DB.Create(&user).Error;err != nil{
		return err
	}
	return


}