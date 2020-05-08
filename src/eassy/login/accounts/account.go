package accounts

import (
	"game_framework/src/eassy/db/mysql/login"
	read "game_framework/src/eassy/db/mysql/login/read"
	write "game_framework/src/eassy/db/mysql/login/write"
)

func (p *account) FindAccount(phone string) (ret bool, acc *login.Account) {
	acc = read.FindUserByPhone(r, phone)
	if acc != nil && acc.Id > 0 {
		ret = true
	}
	return
}

func (p *account) Register(phone, pwd string) bool {
	return write.CreateAccount(w, phone, pwd) == nil
	//return true
}

func (p *account) UpdateInfo(acc *login.Account) {
	write.UpdateAccount(w, acc)
}
