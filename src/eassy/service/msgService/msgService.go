package msgService

import (
	"log"
	"reflect"
	"sync"
)

var msgMgrSingleton *msgMgr
var once sync.Once

func GetMsgService() *msgMgr {
	once.Do(func() {
		msgMgrSingleton = &msgMgr{
			msgMap: map[int]*MsgInfo{},
		}
	})
	return msgMgrSingleton
}

type msgMgr struct {
	msgMap map[int]*MsgInfo
}

//methodMap := map[string]interface{}{
//"SetName": p.SetName,
//"GetName": p.GetName,
//"SetAge":  p.SetAge,
//"GetAge":  p.GetAge,

type MsgInfo struct {
	Route       string
	MsgReqType  reflect.Type
	MsgRespType reflect.Type
}

func (m *msgMgr) Register(route int, fun interface{}, method string, msgReq interface{}, msgResp interface{}) {
	msgType := reflect.TypeOf(msgReq)
	if msgType == nil || msgType.Kind() != reflect.Ptr {
		log.Fatal("message request pointer required")
		return
	}

	if _, ok := m.msgMap[route]; ok {
		log.Fatal("route %s is already registered", route)
		return
	}

	msgRespType := reflect.TypeOf(msgResp)
	if msgRespType == nil || msgRespType.Kind() != reflect.Ptr {
		log.Fatal("message response pointer required")
	}

	i := new(MsgInfo)
	i.Route = method
	i.MsgReqType = msgType
	i.MsgRespType = msgRespType
	m.msgMap[route] = i

}

func (m *msgMgr) RegisterPush(route int, method string) {
	if _, ok := m.msgMap[route]; ok {
		log.Fatal("route %s is already registered", route)
		return
	}
	i := new(MsgInfo)
	i.Route = method
	m.msgMap[route] = i
}

func (m *msgMgr) GetMsgByRouteId(route int) (info *MsgInfo, ok bool) {
	info, ok = m.msgMap[route]
	return
}

func (m *msgMgr) GetMethodByRouteId(route int) (method string) {
	info, ok := m.msgMap[route]
	if !ok {
		return ""
	}
	return info.Route
}

func (m *msgMgr) GetMsgMap() map[int]*MsgInfo {
	return m.msgMap
}
