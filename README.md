# danmu-play
bilibili直播间弹幕互动机器人

# cmd
All available command

| Cmd               | Feature |
| -------:          | ------- |
| **#签到**          | `签到 积分+1` |
| **#积分**          | `弹幕返回用户当前积分` |
| **#运势**          | `弹幕返回用户今天运势` |
| **#能力**          | `弹幕返回用户随机能力和代价(目前有超过30弹幕发送失败bug)` |

# Usage

You need write cookies to etc/danmuplay.yaml

``` yml
cookie: 打开直播间f12获取
roomid: 直播间
redis:
  url: REDIS_URL
```