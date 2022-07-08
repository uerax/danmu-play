/*
 * @Author: ww
 * @Date: 2022-07-09 04:07:55
 * @Description:
 * @FilePath: /danmuplay/fate/fate.go
 */
package fate

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var lucky = []string{"大吉","中吉","小吉","吉","末吉","凶","大凶"}
var star =  []string{"★★★★","★★★☆","★★☆☆","★☆☆☆","☆☆☆","☆☆","☆"}

func GetFate(uid, name string) string {
	now, _ := time.Parse("20060102", time.Now().Format("20060102"))
	
	u, _ := strconv.ParseInt(uid, 10, 64)
	
	rand.Seed(now.Unix()+u+int64(len(name)))
	idx := rand.Intn(7)
	return fmt.Sprintf("[%s] 今日吉凶指数:%s%s", name, star[idx],lucky[idx])
}

