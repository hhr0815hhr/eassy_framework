package dispatch

import (
	"context"
	pb "game_framework/src/eassy/proto"
	"google.golang.org/grpc"
)

func userHandle(cc pb.UserServiceClient, method string, msg interface{}) interface{} {
	if !userFlag {
		initUserFuncMap(cc)
	}
	//todo 有问题
	resp, err := funcMap[method].(func(context.Context, interface{}, ...grpc.CallOption) (interface{}, error))(context.TODO(), msg)
	if err != nil {
		return nil
	}
	return resp
}
