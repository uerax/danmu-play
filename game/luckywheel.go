/*
 * @Author: ww
 * @Date: 2022-08-22 16:13:45
 * @Description:
 * @FilePath: /danmuplay/game/luckywheel.go
 */
package game

import (
	"math/rand"
	"time"
)

type LuckyWheel struct {
	gift []*Gift
	rate []int // 二分法查询 第一个>=的数 i-1的gift
	probability int
}

type Gift struct {
	Title string
	Total int64
	id 	  int64
	Rate  float64
}

func NewLuckyWheel(gift []*Gift) *LuckyWheel {
	tmp := new(LuckyWheel)
	tmp.gift = gift
	tmp.Generate()
	return tmp
}

func (t *LuckyWheel) Generate() {
	min := 100.0
	for _, g := range t.gift {
		// id := method
		// g.id = id
		if min > g.Rate {
			min = g.Rate
		}
	}
	max := 100 / min
	giftRange := make([]int, 0)
	giftRange = append(giftRange, 0)
	
	pre := 0
	for _, g := range t.gift {
		giftRange = append(giftRange, pre + int(g.Rate * max / 100))
		pre += int(g.Rate * max / 100)
	}
	giftRange = append(giftRange, int(max))
	t.rate = giftRange
	t.probability = int(max)
}

func (t *LuckyWheel) binarySearch(num, l, r int) int {

	if num == 0 || r < l ||  r >= len(t.rate) || l < 0 {
		return 0
	}
	idx := (l+r)/2
	if t.rate[idx] == num {
		return idx + 1
	}
	if t.rate[idx] > num {
		if t.rate[idx - 1] <= num {
			return idx
		}
		return t.binarySearch(num, l, idx - 1)
	}
	return t.binarySearch(num, idx + 1, r)

}

// 0为没抽中
func (t *LuckyWheel) Lottery() int64 {
	rand.Seed(time.Now().Unix())
	i := t.binarySearch(rand.Intn(t.probability) + 1, 0, len(t.rate)-1)
	if i == 0 {
		return 0
	}
	if t.gift[i-1].Total > 0 {
		t.gift[i-1].Total--
		return t.gift[i-1].id
	}
	return 0
}
