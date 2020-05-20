package landlord

import "game_framework/src/eassy/util"

func GetBeginCards() (A, B, C, L []int) {
	var cards = util.InitCards()
	util.ShuffleCards(cards)
	A = cards[:17]
	B = cards[17:35]
	C = cards[34:52]
	L = cards[52:]
	return
}

func PutCards(cards, put []int) ([]int, bool) {
	m := make(map[int]int)
	for k, v := range cards {
		m[v] = k
	}
	for _, v := range put {
		if k, ok := m[v]; ok {
			cards = append(cards[:k], cards[k+1:]...)
		} else {
			return cards, false
		}
	}
	return cards, true
}

func CompareCards(last, mine []int) bool {
	lastType := GetCardsType(last)
	mineType := GetCardsType(mine)
	if mineType == TYPE_JOKER_BOOM {
		return true
	}
	if lastType == TYPE_JOKER_BOOM {
		return false
	}
	if mineType == lastType {
		return compareSameTypeCards(last, mine, lastType)
	}

	if mineType == TYPE_BOOM {
		return true
	}
	return false
}

func GetCardsType(cards []int) (cardType int) {
	l := len(cards)
	switch {
	case l == 1:
		cardType = TYPE_SINGLE
	case l == 2:
		if getCardValue(cards[0]) != getCardValue(cards[1]) {
			cardType = TYPE_ERROR
			goto GOTO
		}
		if getCardValue(cards[0]) == 14 {
			cardType = TYPE_JOKER_BOOM
		} else {
			cardType = TYPE_COUPLE
		}
	case l == 3:
		if isThree(cards) {
			cardType = TYPE_THREE
		} else {
			cardType = TYPE_ERROR
		}
	case l == 4:
		if isBoom(cards) {
			cardType = TYPE_BOOM
			goto GOTO
		}
		if isThreeWithSingle(cards) {
			cardType = TYPE_THREE_WITH_SINGLE
			goto GOTO
		}
		cardType = TYPE_ERROR
	case l >= 5:
		if isShunZi(cards) {
			cardType = TYPE_SHUNZI
			goto GOTO
		}
		if isThreeWithCouple(cards) {
			cardType = TYPE_THREE_WITH_COUPLE
			goto GOTO
		}
		if isCouples(cards) {
			cardType = TYPE_COUPLES
			goto GOTO
		}
		if isPlane(cards) {
			cardType = TYPE_PLANE
			goto GOTO
		}
		if isPlaneWithSingles(cards) {
			cardType = TYPE_PLANE_SINGLES
			goto GOTO
		}
		if isPlaneWithCouples(cards) {
			cardType = TYPE_PLANE_COUPLES
			goto GOTO
		}
		if isFourWithCouple(cards) {
			cardType = TYPE_FOUR_WITH_COUPLE
			goto GOTO
		}
		if isFourWithSingle(cards) {
			cardType = TYPE_FOUR_WITH_SINGLE
			goto GOTO
		}
	default:
		cardType = TYPE_ERROR
	}
GOTO:
	return
}
