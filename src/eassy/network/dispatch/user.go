package dispatch

import (
	"context"
	pb "game_framework/src/eassy/proto"
)

var c pb.GameServiceClient

func init() {
	funcMap["SendMail"] = c.SendMail
}

func gameHandle(cc pb.GameServiceClient, method string, msg interface{}) interface{} {
	c = cc
	resp, err := funcMap[method].(func(context.Context, interface{}) (interface{}, error))(context.TODO(), msg)
	if err != nil {
		return nil
	}
	return resp
}
