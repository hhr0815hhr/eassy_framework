package read

import (
	"game_framework/src/eassy/db/mysql/login"
	"github.com/jinzhu/gorm"
)

func FindUserByPhone(db *gorm.DB, phone string) (acc *login.Account) {
	acc = new(login.Account)
	db.Where("phone = ?", phone).First(acc)
	return
}
