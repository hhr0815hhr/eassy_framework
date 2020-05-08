package write

import (
	"game_framework/src/eassy/db/mysql/login"
	"game_framework/src/eassy/service/idService"
	"game_framework/src/eassy/util"
	"github.com/jinzhu/gorm"
	"time"
)

func CreateAccount(db *gorm.DB, phone string, AccPwd string) error {
	times := time.Now().Unix()
	acc := &login.Account{
		AccountId:     idService.GenerateID().Int64(),
		Phone:         phone,
		Pwd:           util.MD5(AccPwd),
		CreatedTime:   times,
		LastLoginTime: times,
		LastIP:        0,
	}
	//return db.NewRecord(acc)
	return db.Create(acc).Error
}

func UpdateAccount(db *gorm.DB, acc *login.Account) {
	db.Save(acc)
}
