/*
 * Maintained by jemo from 2021.3.16 to now
 * Created by jemo on 2021.3.16 21:06:30
 * jwt
 * 获取jwt
 */

package handle

import (
  "fmt"
  "log"
  "time"
  "net/http"
  "encoding/json"
  "database/sql"
  "github.com/silenceper/wechat/v2/officialaccount/oauth"
  "github.com/jiangtaozy/sister-naughty-api/database"
  "github.com/dgrijalva/jwt-go"
)

func Jwt(w http.ResponseWriter, r *http.Request) {
  /*
  query := r.URL.Query()
  code := query["code"][0]
  officialAccountOauth := OfficialAccount.GetOauth()
  userAccessToken, err := officialAccountOauth.GetUserAccessToken(code)
  if err != nil {
    log.Println("jwt.go-get-user-access-token-error: ", err)
    http.Error(w, err.Error(), 500)
    return
  }
  */
  // test
  userAccessToken := oauth.ResAccessToken{
    AccessToken: "ACCESS_TOKEN",
    ExpiresIn: 7200,
    RefreshToken: "REFRESH_TOKEN",
    OpenID: "oIPhb0UTUDakjefaZPx0A_-SC_D9",
    //Scope: "snsapi_base",
    Scope: "snsapi_userinfo",
  }
  openId := userAccessToken.OpenID
  db := database.DB
  var id int
  err := db.QueryRow(`
    SELECT
      id
    FROM
      user
    WHERE
      openid = ?
  `, openId).Scan(&id)
  if err != nil {
    if err == sql.ErrNoRows {
      scope := userAccessToken.Scope
      if scope == "snsapi_userinfo" {
        /*
        accessToken := userAccessToken.AccessToken
        lang := "zh_CN"
        officialAccountOauth := OfficialAccount.GetOauth()
        userInfo, err := officialAccountOauth.GetUserInfo(accessToken, openId, lang)
        if err != nil {
          log.Println("jwt.go-get-user-info-error: ", err)
          http.Error(w, err.Error(), 500)
          return
        }
        */
        // test
        userInfo := oauth.UserInfo{
          OpenID: "oIPhb0UTUDakjefaZPx0A_-SC_D9",
          Nickname: "江涛zy6",
          Sex: 1,
          Province: "Hebei",
          City: "Baoding",
          Country: "China",
          HeadImgURL: "http://wx.qlogo.cn/mmopen/vi_32/xWnqaSB8R7VIBugWZjWoge03K7oAl2XYsybhfdwlMgrOricvUKaWh3xkSDm4s9UCRB7V7eGMfx2DSypbYMbpEfg/132",
          Privilege: []string{},
        }
        stmtInsert, err := db.Prepare(`
          INSERT INTO user (
            openid,
            nickname,
            sex,
            city,
            province,
            country,
            headimgurl,
            privilege
          ) VALUES (
            ?, ?, ?, ?, ?, ?, ?, ?
          )
        `)
        if err != nil {
          log.Println("jwt.go-insert-user-prepare-error: ", err)
          http.Error(w, err.Error(), 500)
          return
        }
        defer stmtInsert.Close()
        _, err = stmtInsert.Exec(
          userInfo.OpenID,
          userInfo.Nickname,
          userInfo.Sex,
          userInfo.City,
          userInfo.Province,
          userInfo.Country,
          userInfo.HeadImgURL,
          fmt.Sprint(userInfo.Privilege),
        )
        if err != nil {
          log.Println("jwt.g-insert-user-exec-error: ", err)
          http.Error(w, err.Error(), 500)
          return
        }
        err = db.QueryRow(`
          SELECT
            id
          FROM
            user
          WHERE
            openid = ?
        `, userInfo.OpenID).Scan(&id)
      } else {
        json.NewEncoder(w).Encode(map[string]interface{}{
          "scope": "snsapi_userinfo",
        })
        return
      }
    } else {
      log.Println("jwt.go-count-openid-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
  }
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "id": id,
    "exp": time.Now().Add(time.Hour * 72).Unix(),
  })
  tokenString, err := token.SignedString([]byte(jwtKey))
  if err != nil {
    log.Println("jwt.go-sign-token-error: ", err)
    http.Error(w, err.Error(), 500)
    return
  }
  json.NewEncoder(w).Encode(map[string]interface{}{
    "jwt": tokenString,
  })
}
