# danmu-play
bilibili直播间弹幕互动机器人
一个低门槛的弹幕互动机器人,不需要域名,不需要服务器,更不需要数据库(redis可改为文本存储).只需要本地启动即可运行
抽奖的结果通过github page展示

# cmd
All available command

| Cmd               | Feature |
| -------:          | ------- |
| **#签到**          | `签到 积分+1` |
| **#积分**          | `弹幕返回用户当前积分` |
| **#运势**          | `弹幕返回用户今天运势` |
| **#能力**          | `弹幕返回用户随机能力和代价(目前有超过30弹幕发送失败bug)` |
| **#rollstart 持续时间 奖品**         | `开启抽奖 只有配置的管理uid有效` |
| **#roll**         | `参与抽奖` |
| **#rollring**         | `结束抽奖并开奖 只有配置的管理uid有效` |
| **#{A-D} 积分**    | `押注的积分` |


# Usage

You need write cookies to etc/danmuplay.yaml

``` yml
cookie: 打开直播间f12获取
roomid: 直播间
userid: 管理员的uid, 只有该用户的弹幕可以开启活动
redis:
  url: REDIS_URL
```