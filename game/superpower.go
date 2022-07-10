/*
 * @Author: ww
 * @Date: 2022-07-10 03:54:32
 * @Description:
 * @FilePath: /danmuplay/game/superpower.go
 */
package game

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var buff = []string{"暂停时间", "复制任何东西", "治疗任何疾病", "控制自己的梦境", "飞行", "现实生活存读档", "发射镭射光线", "读取别人内心", "瞬间移动", "不会被雨淋到", "使用易容术", "隐身", "给手机充满电", "使用分身术", "瞬间让自己睡着", "自爆", "时间旅行", "永远不死", "每天得到一桶泡面", "免疫疾病", "玩任何游戏都能赢", "不需要吃饭喝水", "不需要睡觉", "牛牛增加10厘米", "永远不会排泄", "存款翻倍", "精通一门外语", "倒立吃饭", "预知未来", "和幽灵对话", "不会变老", "消除存在感", "吸取别人寿命"}
var debuff = []string{"需要支付10w元", "自己会似", "无法穿衣服", "寿命减半", "身高减少20厘米", "双目失明", "牛牛缩减3厘米", "会被全国通缉", "寿命-1s", "永远找不到女朋友", "少一个肾", "什么都不会发生", "体重增加30kg", "会得癌症", "考试永远不及格", "失去所有米", "成为孤儿", "永远不能上网", "永远不能瑟瑟", "全世界的人都会", "变丑", "变成智障", "性别反转", "只能在没人的时候用", "每天早上都宿醉", "每天失眠", "导致世界末日", "双目失明", "出门就肚子疼", "散发体臭", "牛牛会发光", "只能倒立"}

func GetSuperpower(uid, name string) string {
	now:= time.Now().Unix()
	
	u, _ := strconv.ParseInt(uid, 10, 64)
	
	rand.Seed(now+u+int64(len(name)))
	x, y := rand.Intn(len(buff)), rand.Intn(len(debuff))
	return fmt.Sprintf("[%s]的超能力是:%s 但是:%s", name, buff[x],debuff[y])
}