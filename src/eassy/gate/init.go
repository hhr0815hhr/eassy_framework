package gate

import "game_framework/src/eassy/service/codecService"

var CliManager *CliMgr

func init() {
	CliManager = &CliMgr{
		Clis: make(map[int64]*Cli),
	}
	codecService.SetCodecType(codecService.TYPE_CODEC_PROTOBUF)
}
