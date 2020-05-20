package landlord

func countValueNums(cards []int) map[int]int {
	m := make(map[int]int)
	for _, v := range cards {
		if _, ok := m[getCardValue(v)]; ok {
			m[getCardValue(v)]++
		} else {
			m[getCardValue(v)] = 1
		}
	}
	return m
}

func compareSingleCard(A, B int) bool {
	A = getCardRealValue(A)
	B = getCardRealValue(B)
	return A < B
}

func compareSameTypeCards(last, mine []int, t int) bool {
	sortCards(last)
	sortCards(mine)
	switch t {
	case TYPE_SINGLE, TYPE_COUPLE, TYPE_THREE, TYPE_BOOM:
		return compareSingleCard(last[0], mine[0])
	case TYPE_COUPLES, TYPE_SHUNZI, TYPE_PLANE:
		if len(last) != len(mine) {
			return false
		}
		return compareSingleCard(last[0], mine[0])
	case TYPE_THREE_WITH_SINGLE:
		return compareSingleCard(last[1], mine[1])
	case TYPE_THREE_WITH_COUPLE:
		return compareSingleCard(last[2], mine[2])
	case TYPE_PLANE_SINGLES, TYPE_PLANE_COUPLES:
		if len(last) != len(mine) {
			return false
		}
		m1 := make(map[int]int)
		var k1, k2 int
		for _, v := range last {
			if _, ok := m1[v*10+1]; ok {
				m1[v*10+1]++
			} else {
				m1[v*10+1] = 1
			}
		}
		for k, v := range m1 {
			if v == 3 {
				k1 = k
			}
		}
		m2 := make(map[int]int)
		for _, v := range last {
			if _, ok := m2[v*10+1]; ok {
				m2[v*10+1]++
			} else {
				m2[v*10+1] = 1
			}
		}
		for k, v := range m2 {
			if v == 3 {
				k2 = k
			}
		}
		return getCardRealValue(k1) < getCardRealValue(k2)
	case TYPE_FOUR_WITH_SINGLE, TYPE_FOUR_WITH_COUPLE:
		if len(last) != len(mine) {
			return false
		}
		m1 := make(map[int]int)
		var k1, k2 int
		for _, v := range last {
			if _, ok := m1[v*10+1]; ok {
				m1[v*10+1]++
			} else {
				m1[v*10+1] = 1
			}
		}
		for k, v := range m1 {
			if v == 4 {
				k1 = k
			}
		}
		m2 := make(map[int]int)
		for _, v := range last {
			if _, ok := m2[v*10+1]; ok {
				m2[v*10+1]++
			} else {
				m2[v*10+1] = 1
			}
		}
		for k, v := range m2 {
			if v == 4 {
				k2 = k
			}
		}
		return getCardRealValue(k1) < getCardRealValue(k2)
	}
	return false
}

//vals 已经排序
func checkContinueValue(vals []int) (T bool) {
	l := len(vals)
	if vals[0] == 1 {
		if vals[l-1] != 13 {
			return
		}
		T = checkContinueValue(vals[1:])
		return
	}
	for i := 0; i < l-1; i++ {
		if vals[i] == 2 || vals[i] == 14 {
			return
		}
		if vals[i]+1 != vals[i+1] {
			return
		}
	}
	T = true
	return
}

func sortCards(cards []int) {
	for i := 0; i < len(cards)-1; i++ {
		for j := len(cards) - 1; j > i; j-- {
			if cards[j] < cards[j-1] {
				cards[j], cards[j-1] = cards[j-1], cards[j]
			}
		}
	}
}

func getCardValue(card int) int {
	return card / 10
}

func getCardRealValue(card int) int {
	v := getCardValue(card)
	if v > 2 {
		if card == 141 {
			return 16
		}
		if card == 142 {
			return 17
		}
		return v
	}
	if v == 2 {
		return 15
	}
	return 14
}

func isPlane(cards []int) (T bool) {
	l := len(cards)
	if l%3 == 0 && l >= 6 {
		m := countValueNums(cards)
		tmp := make([]int, 0)
		for _, v := range m {
			if v == 3 {
				tmp = append(tmp, v)
			}
		}
		sortCards(tmp)
		T = checkContinueValue(tmp)
		if len(tmp)*3 != len(cards) {
			T = false
		}
		//sortCards(cards)
		//var tmp = make([]int,0)
		//for i := 0; i < l-3; {
		//	vTmp1 := getCardValue(cards[i])
		//	vTmp2 := getCardValue(cards[i+1])
		//	vTmp3 := getCardValue(cards[i+2])
		//	if vTmp1 != vTmp2 || vTmp1 != vTmp3 {
		//		return
		//	}
		//	tmp = append(tmp,vTmp1)
		//	i += 3
		//}
		//T = checkContinueValue(tmp)
	}
	return
}
func isPlaneWithSingles(cards []int) (T bool) {
	l := len(cards)
	if l >= 8 && l%4 == 0 {
		m := countValueNums(cards)
		tmp := make([]int, 0)
		for _, v := range m {
			if v == 3 {
				tmp = append(tmp, v)
			}
		}
		sortCards(tmp)
		T = checkContinueValue(tmp)
		if len(tmp)*4 != len(cards) {
			T = false
		}
	}
	return
}
func isPlaneWithCouples(cards []int) (T bool) {
	l := len(cards)
	if l >= 10 && l%5 == 0 {
		m := countValueNums(cards)
		tmp := make([]int, 0)
		tmpCouple := make([]int, 0)
		for _, v := range m {
			if v == 3 {
				tmp = append(tmp, v)
			}
			if v == 2 {
				tmpCouple = append(tmpCouple, v)
			}
		}
		sortCards(tmp)
		T = checkContinueValue(tmp)
		if len(tmp)*3+len(tmpCouple)*2 != len(cards) {
			T = false
		}
	}
	return
}
func isShunZi(cards []int) (T bool) {
	l := len(cards)
	if l < 5 || l > 12 {
		return
	}
	sortCards(cards)
	var tmp = make([]int, l)
	for k, v := range cards {
		tmp[k] = getCardValue(v)
	}
	T = checkContinueValue(tmp)
	return
}
func isCouples(cards []int) (T bool) {
	l := len(cards)
	if l%2 != 0 || l < 6 {
		return
	}
	m := countValueNums(cards)
	tmp := make([]int, 0)
	for k, v := range m {
		if v != 2 {
			return
		}
		tmp = append(tmp, k)
	}
	sortCards(tmp)
	T = checkContinueValue(tmp)
	return
}
func isFourWithSingle(cards []int) (T bool) {
	if len(cards) == 6 {
		m := countValueNums(cards)
		for _, v := range m {
			if v == 4 {
				T = true
			}
		}
		return
	}
	return
}
func isFourWithCouple(cards []int) (T bool) {
	//todo
	if len(cards) == 8 {
		m := countValueNums(cards)
		for _, v := range m {
			if v == 4 {
				T = true
			}
		}
	}
	return
}
func isBoom(cards []int) (T bool) {
	if len(cards) == 4 {
		tmp1 := getCardValue(cards[0])
		tmp2 := getCardValue(cards[1])
		tmp3 := getCardValue(cards[2])
		tmp4 := getCardValue(cards[3])
		if tmp1 == tmp2 && tmp2 == tmp3 && tmp3 == tmp4 {
			T = true
		}
	}
	return
}
func isThree(cards []int) bool {
	if len(cards) == 3 && getCardValue(cards[0]) == getCardValue(cards[1]) && getCardValue(cards[1]) == getCardValue(cards[2]) {
		return true
	}
	return false
}
func isThreeWithSingle(cards []int) bool {
	if len(cards) == 4 {
		sortCards(cards)
		tmp1 := getCardValue(cards[0])
		tmp2 := getCardValue(cards[1])
		tmp3 := getCardValue(cards[2])
		tmp4 := getCardValue(cards[3])
		if (tmp1 != tmp2 && tmp3 == tmp4 && tmp2 == tmp4) || (tmp1 == tmp2 && tmp2 == tmp3 && tmp3 != tmp4) {
			return true
		}
	}
	return false
}
func isThreeWithCouple(cards []int) bool {
	if len(cards) == 5 {
		sortCards(cards)
		tmp1 := getCardValue(cards[0])
		tmp2 := getCardValue(cards[1])
		tmp3 := getCardValue(cards[2])
		tmp4 := getCardValue(cards[3])
		tmp5 := getCardValue(cards[4])
		if tmp1 == tmp2 && tmp4 == tmp5 && tmp1 != tmp5 && (tmp3 == tmp4 || tmp2 == tmp3) {
			return true
		}
	}
	return false
}
