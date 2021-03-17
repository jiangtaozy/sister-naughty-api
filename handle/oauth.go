/*
 * Maintained by jemo from 2021.3.16 to now
 * Created by jemo on 2021.3.16 16:36:59
 * Oauth
 * 微信登录
 */

package handle

import (
  "io"
  "log"
  "net/http"
)

func Oauth(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query()
  scope := "snsapi_base"
  if len(query["scope"]) != 0 {
    scope = query["scope"][0]
  }
  oauth := OfficialAccount.GetOauth()
  redirectURI := r.Referer()
  state := ""
  err := oauth.Redirect(w, r, redirectURI, scope, state)
  if err != nil {
    log.Println("oauth.go-error: ", err)
    io.WriteString(w, "出错了：" + err.Error())
  }
}
