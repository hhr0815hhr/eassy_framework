package dispatch

import (
	"context"
	pb "game_framework/src/eassy/proto"
)

var funcMap map[string]interface{}
var c pb.UserServiceClient

func init() {
	funcMap["login"] = c.Login
}

func userHandle(cc pb.UserServiceClient, method string, msg interface{}) interface{} {
	c = cc
	resp, err := funcMap[method].(func(context.Context, interface{}) (interface{}, error))(context.TODO(), msg)
	if err != nil {
		return nil
	}
	return resp
}
