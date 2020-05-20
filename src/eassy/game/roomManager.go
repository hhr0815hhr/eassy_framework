package game

import "sync"

type RoomList struct {
	RoomList map[int]*Room
	Locker   sync.Mutex
}

type IRoomManager interface {
	CreateRoom(roomType string)
	DestroyRoom(roomId int)
	Join(playerId int64, roomId int)
}

var RoomMgr *RoomList
var nextRoomId int

func init() {
	RoomMgr = &RoomList{
		RoomList: make(map[int]*Room),
		Locker:   sync.Mutex{},
	}
	nextRoomId = 1
}

func (p *RoomList) CreateRoom(roomType string) {
	p.Locker.Lock()
	defer p.Locker.Unlock()
	p.RoomList[nextRoomId] = CreateRoom(roomType)
	nextRoomId++
}

func (p *RoomList) Join(playerId int64, roomId int) {
	p.Locker.Lock()
	defer p.Locker.Unlock()
	if len(p.RoomList[roomId].Players) >= 3 {
		return
	}
	p.RoomList[roomId].JoinRoom(playerId)
}

func (p *RoomList) DestroyRoom(roomId int) {
	p.Locker.Lock()
	defer p.Locker.Unlock()
	if _, ok := p.RoomList[roomId]; ok {
		delete(p.RoomList, roomId)
	}
}
