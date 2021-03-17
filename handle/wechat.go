/*
 * Maintained by jemo from 2021.3.15 to now
 * Created by jemo on 2021.3.15 16:18:40
 * Wechat
 */

package handle

import (
  "log"
  "net/http"
  wechat "github.com/silenceper/wechat/v2"
  "github.com/silenceper/wechat/v2/cache"
  offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
  "github.com/silenceper/wechat/v2/officialaccount/message"
  "github.com/silenceper/wechat/v2/officialaccount"
)

var OfficialAccount *officialaccount.OfficialAccount

func Wechat(w http.ResponseWriter, r *http.Request) {
  server := OfficialAccount.GetServer(r, w)
  server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
    log.Println("msg: ", msg)
    text := message.NewText(msg.Content)
    return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
  })
  err := server.Serve()
  if err != nil {
    log.Println("wechat.go-error: ", err)
    return
  }
  server.Send()
}

func InitWechat() {
  wc := wechat.NewWechat()
  redisOpts := &cache.RedisOpts{
    Host: Host,
    Database: Database,
    MaxActive: MaxActive,
    MaxIdle: MaxIdle,
    IdleTimeout: IdleTimeout, // second
  }
  redisCache := cache.NewRedis(redisOpts)
  cfg := &offConfig.Config{
    AppID: AppID,
    AppSecret: AppSecret,
    Token: AppToken,
    Cache: redisCache,
  }
  OfficialAccount = wc.GetOfficialAccount(cfg)
}
