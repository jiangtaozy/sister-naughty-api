/*
 * Maintained by jemo from 2021.3.19 to now
 * Created by jemo on 2021.3.19 16:26:03
 * Get JWT User Id
 */

package handle

import (
  "log"
  "errors"
  "github.com/dgrijalva/jwt-go"
)

func getJwtUserId(jwtString string) (interface{}, error) {
  token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
    return []byte(jwtKey), nil
  })
  if err != nil {
    log.Println("main-image-browse-record.go-jwt-parse-error: ", err)
    return nil, err
  }
  if !token.Valid {
    log.Println("main-image-browse-record.go-jwt-invalid")
    return nil, errors.New("jwt-invalid")
  }
  claims, ok := token.Claims.(jwt.MapClaims)
  if !ok {
    log.Println("main-image-browse-record.go-jwt-claims-error")
    return nil, errors.New("jwt-claims-error")
  }
  userId := claims["id"]
  return userId, nil
}
