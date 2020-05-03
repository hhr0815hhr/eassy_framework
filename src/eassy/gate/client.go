package gate

import (
	"game_framework/src/eassy/network/dispatch"
	"game_framework/src/eassy/network/msg"
	"golang.org/x/net/websocket"
	"sync"
)

type CliMgr struct {
	Clis map[uint]*Cli
	Lock sync.RWMutex
}

type ICliMgr interface {
	Connect(ws *websocket.Conn)
	DisConnect(ws *websocket.Conn)
	GetCliIdByWs(ws *websocket.Conn) uint
}

var nextCliId uint

func (p *CliMgr)Connect(ws *websocket.Conn)  {
	p.Lock.Lock()
	defer p.Lock.Unlock()
	p.Clis[nextCliId] = &Cli{
		Id:   nextCliId,
		Conn: ws,
	}
	nextCliId++
}

func (p *CliMgr)DisConnect(ws *websocket.Conn)  {
	p.Lock.Lock()
	defer p.Lock.Unlock()
	delete(p.Clis,p.GetCliIdByWs(ws))
}

func (p *CliMgr)GetCliIdByWs(ws *websocket.Conn) uint {
	p.Lock.RLock()
	defer p.Lock.RUnlock()
	for k,v :=range p.Clis {
		if v.Conn == ws {
			return k
		}
	}
	return 0
}

type Cli struct {
	Id   uint
	Conn *websocket.Conn
	ServerId uint
}

type ICli interface {
	RecvData()
	handleData(content []byte)
	Send(protoId int,buffer []byte)
}

func (p *Cli)RecvData()  {
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

func (p *Cli)handleData(content []byte)  {
	pkgType,body := msg.PkgDecode(content)
	switch pkgType {
	case msg.TYPE_HEARTBEAT:
		var msg []byte
		p.Send(0,msg)
	case msg.TYPE_DATA:
		//dispatch
		protoId,buffer := msg.MsgUnpack(body)
		dispatch.Dispatch(protoId,buffer)
	case msg.TYPE_HANDSHAKE,msg.TYPE_HANDSHAKE_ACK,msg.TYPE_KICK:
		//skip
	default:
		//错误的数据类型
	}
}

func (p *Cli)Send(protoId int,buffer []byte)  {
	bytes := msg.PkgEncode(msg.TYPE_DATA,msg.MsgPack(protoId,buffer))
	websocket.Message.Send(p.Conn,bytes)
	//p.Conn.Write(bytes)
}
