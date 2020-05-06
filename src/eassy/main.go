package main

import (
	"game_framework/src/eassy/center"
	"game_framework/src/eassy/conf"
	"game_framework/src/eassy/game"
	"game_framework/src/eassy/gate"
	"game_framework/src/eassy/login"
)

const IsDev = true

func init() {
	//todo 加载配置
	conf.GetConf(IsDev)

}

func main() {
	// os.Args[0] == 执行文件的名字
	// os.Args[1] == 第一个参数
	//args := os.Args
	//if len(args)<3 {
	//	panic("参数小于2个！！！ 例如：xxx.exe +【端口】+【服务器类型】")
	//	return
	//}
	//port,err := strconv.Atoi(args[1])
	//if err != nil {
	//	panic("第一个参数要整数！！！port")
	//	return
	//}
	args := []string{"eassy", "5050", "game"}
	switch args[2] {
	case "gate":
		gate.Run(args[1])
	case "login":
		login.Run(args[1])
	case "game":
		game.Run(args[1])
	case "center":
		center.Run(args[1])
	default:
		panic("参数错误！！！服务类型为 gate/login/game/center")
		return
	}
}
