package gate

import (
	game "game_framework/src/eassy/game/service"
	pb "game_framework/src/eassy/proto"
	"game_framework/src/eassy/service/msgService"
	user "game_framework/src/eassy/user/service"
)

var gs *game.GameService
var us *user.UserService

func init() {
	msgService.GetMsgService().Register(10001, us.Login, "Login", &pb.C2S_Login_10001{}, &pb.S2C_Login_10001{})
}
