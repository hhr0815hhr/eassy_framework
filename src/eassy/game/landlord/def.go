package landlord

//各状态时间
const (
	TIME_PUT_CARDS = 5
	TIME_CALL_BOSS = 5
	TIME_TURN      = 15
)

//状态
const (
	STATE_END       = iota
	STATE_PUT_CARDS //发牌
	STATE_CALL_BOSS //叫地主
	STATE_GAME
)

//牌型
const (
	TYPE_ERROR = iota
	TYPE_SINGLE
	TYPE_COUPLE
	TYPE_COUPLES
	TYPE_THREE
	TYPE_THREE_WITH_SINGLE
	TYPE_THREE_WITH_COUPLE
	TYPE_SHUNZI
	TYPE_PLANE
	TYPE_PLANE_SINGLES
	TYPE_PLANE_COUPLES
	TYPE_FOUR_WITH_SINGLE
	TYPE_FOUR_WITH_COUPLE
	TYPE_BOOM
	TYPE_JOKER_BOOM
)

//牌值大小
const (
	VAL_3 = iota
	VAL_4
	VAL_5
	VAL_6
	VAL_7
	VAL_8
	VAL_9
	VAL_10
	VAL_11
	VAL_12
	VAL_13
	VAL_1
	VAL_2
	VAL_14
)
