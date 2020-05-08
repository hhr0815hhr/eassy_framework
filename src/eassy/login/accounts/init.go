package accounts

import (
	"game_framework/src/eassy/db/mysql"
	"github.com/jinzhu/gorm"
)

type account struct{}

var (
	Acc *account
	r   *gorm.DB
	w   *gorm.DB
)

func init() {
	Acc = &account{}
}

func InitDbConn() {
	r, w = mysql.Connect("account")
}
