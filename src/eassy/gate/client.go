package gate

import (
	"game_framework/src/eassy/network/dispatch"
	"game_framework/src/eassy/network/msg"
	"game_framework/src/eassy/service/codecService"
	"game_framework/src/eassy/service/idService"
	"golang.org/x/net/websocket"
	"sync"
)

type CliMgr struct {
	Clis map[int64]*Cli
	Lock sync.RWMutex
}

type ICliMgr interface {
	Connect(ws *websocket.Conn) *Cli
	DisConnect(ws *websocket.Conn)
	GetCliIdByWs(ws *websocket.Conn) int64
}

func (p *CliMgr) Connect(ws *websocket.Conn) *Cli {
	p.Lock.Lock()
	defer p.Lock.Unlock()
	cliId := idService.GenerateID().Int64()
	p.Clis[cliId] = &Cli{
		Id:   cliId,
		Conn: ws,
	}
	return p.Clis[cliId]
}

func (p *CliMgr) DisConnect(ws *websocket.Conn) {
	p.Lock.Lock()
	defer p.Lock.Unlock()
	delete(p.Clis, p.GetCliIdByWs(ws))
}

func (p *CliMgr) GetCliIdByWs(ws *websocket.Conn) int64 {
	p.Lock.RLock()
	defer p.Lock.RUnlock()
	for k, v := range p.Clis {
		if v.Conn == ws {
			return k
		}
	}
	return 0
}

type Cli struct {
	Id       int64
	Conn     *websocket.Conn
	ServerId uint
}

type ICli interface {
	RecvData()
	handleData(content []byte)
	Send(protoId int, buffer []byte)
}

func (p *Cli) RecvData() {
	for {
		var content []byte
		if err := websocket.Message.Receive(p.Conn, &content); err != nil {
			break
		}
		if len(content) == 0 || len(content) >= 4096 {
			break
		}
		go p.handleData(content)
	}

}

func (p *Cli) handleData(content []byte) {
	pkgType, body := msg.PkgDecode(content)
	switch pkgType {
	case msg.TYPE_HEARTBEAT:
		var msg []byte
		p.Send(0, msg)
	case msg.TYPE_DATA:
		//dispatch
		protoId, buffer := msg.MsgUnpack(body)
		codecService.GetCodecService().Unmarshal(buffer, protoMsg)
		if err != nil {
			return
		}
		dispatch.Dispatch(protoId, protoMsg)
	case msg.TYPE_HANDSHAKE, msg.TYPE_HANDSHAKE_ACK, msg.TYPE_KICK:
		//skip
	default:
		//错误的数据类型
	}
}

func (p *Cli) Send(protoId int, buffer []byte) {
	bytes := msg.PkgEncode(msg.TYPE_DATA, msg.MsgPack(protoId, buffer))
	websocket.Message.Send(p.Conn, bytes)
	//p.Conn.Write(bytes)
}
