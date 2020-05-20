package dispatch

import (
	"context"
	"errors"
	pb "game_framework/src/eassy/proto"
	"log"
)

func gameHandle(cc pb.GameServiceClient, method string, msg interface{}) interface{} {
	//if !gameFlag {
	//	initGameFuncMap(cc)
	//}
	var resp interface{}
	var err error
	switch method {
	case "SendMail":
		resp, err = cc.SendMail(context.TODO(), msg.(*pb.MailRequest))
	default:
		err = errors.New("非法的方法")
	}
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return resp
}
