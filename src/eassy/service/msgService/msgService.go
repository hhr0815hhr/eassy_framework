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
			msgMap:   map[int]string{},
			msgArray: make([]*MsgInfo, 0),
		}
	})
	return msgMgrSingleton
}

type msgMgr struct {
	msgMap   map[int]string
	msgArray []*MsgInfo
}

type MsgInfo struct {
	Route       string
	MsgReqType  reflect.Type
	MsgRespType reflect.Type
}

func (m *msgMgr) Register(route int, msgReq interface{}, msgResp interface{}) {
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
	i.Route = route
	i.MsgReqType = msgType
	i.MsgRespType = msgRespType

	m.msgMap[route] = uint16(len(m.msgArray) + 1)
	m.msgArray = append(m.msgArray, i)
}

func (m *msgMgr) RegisterPush(route string) {
	if _, ok := m.msgMap[route]; ok {
		log.Fatal("route %s is already registered", route)
		return
	}

	i := new(MsgInfo)
	i.Route = route

	m.msgMap[route] = uint16(len(m.msgArray) + 1)
	m.msgArray = append(m.msgArray, i)
}

func (m *msgMgr) GetMsgByRouteId(route int) *MsgInfo {
	if route > len(m.msgArray) {
		log.Fatal("routeId is out of range")
		return nil
	}
	return m.msgArray[route-1]
}

func (m *msgMgr) GetMsgByRoute(route string) *MsgInfo {
	routeId := m.msgMap[route]
	return m.msgArray[routeId-1]
}

func (m *msgMgr) GetRouteId(route string) uint16 {
	return m.msgMap[route]
}

func (m *msgMgr) GetMsgMap() map[string]uint16 {
	return m.msgMap
}
