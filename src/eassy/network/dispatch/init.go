package dispatch

import (
	pb "game_framework/src/eassy/proto"
)

var funcMap map[string]interface{}
var userFlag bool
var gameFlag bool

func init() {
	funcMap = make(map[string]interface{})
	userFlag = false
	gameFlag = false
}

func initUserFuncMap(cUser pb.UserServiceClient) {
	funcMap["Login"] = cUser.Login
	userFlag = true
	//funcMap["SendMail"] = cGame.SendMail
}

func initGameFuncMap(cGame pb.GameServiceClient) {
	funcMap["SendMail"] = cGame.SendMail

	//funcMap["SendMail"] = cGame.SendMail
	gameFlag = true
}
