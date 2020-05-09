package dispatch

func init() {
	ServiceRoute = make(map[int]string)
	regProto()
}

//注册路由
func regProto() {
	ServiceRoute[10001] = "Login"
	ServiceRoute[10002] = "SetNick"

}
