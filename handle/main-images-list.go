/*
 * Maintained by jemo from 2021.2.26 to now
 * Created by jemo on 2021.2.26 17:40:11
 * Main Images List
 * 主图列表
 */

package handle

import (
  "log"
  "math/rand"
  "net/http"
  "database/sql"
  "encoding/json"
  "github.com/jiangtaozy/sister-naughty-api/database"
)

func MainImagesList(w http.ResponseWriter, r *http.Request) {
  db := database.DB
  var count int
  err := db.QueryRow(`
    SELECT
      COUNT(*)
    FROM
      womenItemMainImage
    WHERE
      isLongImage = 1
  `).Scan(&count)
  if err != nil {
    log.Println("image-images-list.go-count-error: ", err)
    http.Error(w, err.Error(), 500)
    return
  }
  var list []interface{}
  for i := 0; i < 5; {
    image, err := GetOneImage(count)
    if err != nil {
      log.Println("main-images-list.go-GetOneImage-error: ", err)
      http.Error(w, err.Error(), 500)
      return
    }
    hasInList := false
    for j := 0; j < len(list); j++ {
      if image["id"] == list[j].(map[string]interface{})["id"] {
        hasInList = true
        break
      }
    }
    if !hasInList {
      list = append(list, image)
      i++
    }
  }
  json.NewEncoder(w).Encode(list)
}

func GetOneImage(count int) (map[string]interface{}, error) {
  db := database.DB
  randNum := rand.Intn(count)
  var (
    id int64
    searchId int64
    imgPath sql.NullString
  )
  err := db.QueryRow(`
    SELECT
      id,
      searchId,
      imgPath
    FROM
      womenItemMainImage
    WHERE
      isLongImage = 1
    LIMIT ?, 1
  `, randNum).Scan(
    &id,
    &searchId,
    &imgPath,
  )
  if err != nil {
    log.Println("main-images-list.go-GetOneImage-query-error: ", err)
    return nil, err
  }
  image := map[string]interface{}{
    "id": id,
    "searchId": searchId,
    "imgPath": imgPath.String,
  }
  return image, nil
}
