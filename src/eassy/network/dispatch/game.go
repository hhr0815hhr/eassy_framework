package dispatch

import (
	"context"
	pb "game_framework/src/eassy/proto"
)

func gameHandle(cc pb.GameServiceClient, method string, msg interface{}) interface{} {
	if !gameFlag {
		initGameFuncMap(cc)
	}
	resp, err := funcMap[method].(func(context.Context, interface{}) (interface{}, error))(context.TODO(), msg)
	if err != nil {
		return nil
	}
	return resp
}
