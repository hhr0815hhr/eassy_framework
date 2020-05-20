package game

import "sync"

type RoomList struct {
	RoomList map[int]*Room
	Locker   sync.RWMutex
}

type IRoomManager interface {
	CreateRoom(roomType string) int
	DestroyRoom(roomId int)
	Join(playerId int64, roomId int)
	GetRooms() (rooms map[int]*roomBrief)
}

var RoomMgr *RoomList
var nextRoomId int

func init() {
	RoomMgr = &RoomList{
		RoomList: make(map[int]*Room),
		Locker:   sync.RWMutex{},
	}
	nextRoomId = 1
}

func (p *RoomList) CreateRoom(roomType string) int {
	p.Locker.Lock()
	defer p.Locker.Unlock()
	p.RoomList[nextRoomId] = CreateRoom(roomType)
	nextRoomId++
	return nextRoomId - 1
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

type roomBrief struct {
	RoomId     int
	PlayerNums int
}

func (p *RoomList) GetRooms() (rooms map[int]*roomBrief) {
	p.Locker.RLock()
	defer p.Locker.RUnlock()

	for k, _ := range p.RoomList {
		rooms[k].RoomId = k
		rooms[k].PlayerNums = len(p.RoomList[k].Players)
	}
	return
}
