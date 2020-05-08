package login

type Account struct {
	Id            int64
	AccountId     int64
	Phone         string
	Pwd           string
	CreatedTime   int64
	LastLoginTime int64
	LastIP        int64
}
