package util

import (
	"math/rand"
	"reflect"
	"time"
)

func InitCards() []int {
	return []int{
		11, 12, 13, 14,
		21, 22, 23, 24,
		31, 32, 33, 34,
		41, 42, 43, 44,
		51, 52, 53, 54,
		61, 62, 63, 64,
		71, 72, 73, 74,
		81, 82, 83, 84,
		91, 92, 93, 94,
		101, 102, 10, 104,
		111, 112, 113, 114,
		121, 122, 123, 124,
		131, 132, 133, 134,
		141, 142,
	}
}

func ShuffleCards(cards interface{}) {
	rv := reflect.ValueOf(cards)
	if rv.Type().Kind() != reflect.Slice {
		return
	}
	length := rv.Len()
	if length < 2 {
		return
	}
	swap := reflect.Swapper(cards)
	rand.Seed(time.Now().Unix())
	for i := length - 1; i >= 0; i-- {
		j := rand.Intn(length)
		swap(i, j)
	}
	return
}
