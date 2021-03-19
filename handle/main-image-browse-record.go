/*
 * Maintained by jemo from 2021.3.19 to now
 * Created by jemo on 2021.3.19 14:42:15
 * Main image browse record
 * 主图浏览记录
 */

package handle

import (
  "io"
  "log"
  "time"
  "net/http"
  "encoding/json"
  "github.com/jiangtaozy/sister-naughty-api/database"
)

func MainImageBrowseRecord(w http.ResponseWriter, r *http.Request) {
  var body map[string]interface{}
  err := json.NewDecoder(r.Body).Decode(&body)
  if err != nil {
    log.Println("main-image-browse-record-decode-error: ", err)
    http.Error(w, err.Error(), 500)
    return
  }
  jwtString := body["jwt"]
  userId, err := getJwtUserId(jwtString.(string))
  if err != nil {
    log.Println("main-image-browse-record.go-getJwtUserId-error: ", err)
    http.Error(w, err.Error(), 500)
    return
  }
  imageId := body["imageId"]
  start := time.Unix(0, int64(body["start"].(float64)) * int64(time.Millisecond))
  end := time.Unix(0, int64(body["end"].(float64)) * int64(time.Millisecond))
  duration := body["duration"]
  db := database.DB
  stmtInsert, err := db.Prepare(`
    INSERT INTO imageBrowseRecord (
      userId,
      imageId,
      start,
      end,
      duration
    )
    VALUES (
      ?, ?, ?, ?, ?
    )
  `)
  if err != nil {
    log.Println("main-image-browse-record.go-insert-prepare-error: ", err)
    http.Error(w, err.Error(), 500)
    return
  }
  defer stmtInsert.Close()
  _, err = stmtInsert.Exec(
    userId,
    imageId,
    start,
    end,
    duration,
  )
  if err != nil {
    log.Println("main-image-browse-record.go-insert-exec-error: ", err)
    http.Error(w, err.Error(), 500)
    return
  }
  io.WriteString(w, "ok")
}
