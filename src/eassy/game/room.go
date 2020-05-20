package game

import (
	"game_framework/src/eassy/game/landlord"
	"math/rand"
	"sync"
	"time"
)

type Room struct {
	Locker    sync.Mutex
	Type      string
	Players   map[int]*RoomPlayer
	Turn      int   //谁的回合 pos
	CallTime  int   //叫地主次数
	Boss      int   //地主pos
	State     int   //当前状态
	LastCards []int //上家出牌
	BigTurns  int   //连续要不起的次数
}

type RoomPlayer struct {
	Uid    int64
	Nick   string
	Icon   string
	Coin   int64
	Status int
	Pos    int
	Cards  []int
}

func CreateRoom(roomType string) *Room {
	return &Room{
		Type:    roomType,
		Players: make(map[int]*RoomPlayer),
		State:   landlord.STATE_END,
	}
}

type IRoom interface {
	JoinRoom(playerId int64)
	Ready(pos int)
	StartGame()
	CallBoss(pos int, flag int)
	PutCards(pos int, cards []int)
	endGame()
	ticker(duration time.Duration, f func()) //房间内定时器，定时执行func f()
}

func (p *Room) ticker(duration time.Duration, f func()) {
	time.AfterFunc(duration, f)
}

func (p *Room) JoinRoom(playerId int64) {
	p.Locker.Lock()
	defer p.Locker.Unlock()
	newer := &RoomPlayer{
		Uid:  playerId,
		Nick: "",
		Icon: "",

		Coin:   100,
		Status: 0,
		Pos:    len(p.Players),
	}
	p.Players[newer.Pos] = newer
}

func (p *Room) Ready(pos int) {
	p.Locker.Lock()
	defer p.Locker.Unlock()
	p.Players[pos].Status = 1
	if checkAllReady(p.Players) {
		p.StartGame()
	}
}

func (p *Room) StartGame() {
	p.State = landlord.STATE_PUT_CARDS
	//发牌 && 随机确定第一个叫地主玩家
	var tmp [3][]int
	var L []int
	tmp[0], tmp[1], tmp[2], L = landlord.GetBeginCards()
	rand.Seed(time.Now().Unix())
	p.Turn = rand.Intn(2) //确定起始位置
	for k, _ := range p.Players {
		p.Players[k].Cards = tmp[k]
		//todo 给每位玩家发牌 todo广播(自己的牌和L和是否自己回合)
		_ = L
	}
	//utils.TimeOut(p.C.allBoss(),landlord.TIME_PUT_CARDS)
}

func (p *Room) CallBoss(pos int, flag int) {
	if p.Turn != pos {
		//todo 回消息 不是你的回合
		return
	}
	if p.CallTime == 2 && flag == 0 && p.Boss == getNextTurnPos(pos) {
		//叫地主环节结束
		return
	}
	if p.CallTime == 3 {
		p.State = landlord.STATE_GAME
		p.Turn = p.Boss
	}
	if p.CallTime >= 4 {
		//todo reply: wrong data
		return
	}
	p.CallTime++
	if flag == 1 { //要地主
		p.Boss = pos
	}
	return
}

func (p *Room) PutCards(pos int, cards []int) {
	tmpType := landlord.GetCardsType(cards)
	if len(cards) == 0 && len(p.LastCards) > 0 {
		p.BigTurns++
		goto checkNum
	}

	if tmpType == landlord.TYPE_ERROR {
		//todo reply: illegal request
		return
	}
	//都要不起则清空上次牌 means 继续出可以任意打
	if p.BigTurns == 2 {
		p.LastCards = cards
		p.BigTurns = 0
		goto checkNum
	}
	if !landlord.CompareCards(p.LastCards, cards) {
		//todo illegal
		return
	}

checkNum:
	p.Locker.Lock()
	defer p.Locker.Unlock()
	var t bool
	p.Players[pos].Cards, t = landlord.PutCards(p.Players[pos].Cards, cards)
	if !t {
		//todo illegal
		return
	}
	if len(p.Players[pos].Cards) == 0 {
		p.endGame()
	}
}

func (p *Room) endGame() {
	p.State = landlord.STATE_END
}

func checkAllReady(Players map[int]*RoomPlayer) bool {
	for _, v := range Players {
		if v.Status != 1 {
			return false
		}
	}
	return true
}

func getNextTurnPos(nowPos int) int {
	if nowPos+1 == 3 {
		return 0
	}
	return nowPos + 1
}
