/*
 * Maintained by jemo from 2021.2.26 to now
 * Created by jemo on 2021.2.26 17:17:18
 * Main
 */

package main

import (
  "log"
  "net/http"
  "github.com/rs/cors"
  "github.com/jiangtaozy/sister-naughty-api/handle"
  "github.com/jiangtaozy/sister-naughty-api/database"
)

var port = ":8000"

func main() {
  database.InitDB()
  log.Println("listen at ", port)
  handle.InitWechat()
  mux := http.NewServeMux()
  mux.HandleFunc("/mainImagesList", handle.MainImagesList)
  mux.HandleFunc("/mainImageBrowseRecord", handle.MainImageBrowseRecord)
  mux.HandleFunc("/wechat", handle.Wechat)
  mux.HandleFunc("/oauth", handle.Oauth)
  mux.HandleFunc("/jwt", handle.Jwt)
  mux.Handle("/image/", http.StripPrefix("/image/", http.FileServer(http.Dir("../image"))))
  mux.Handle("/", http.FileServer(http.Dir("../sister-naughty/build")))
  handler := cors.Default().Handler(mux)
  log.Fatal(http.ListenAndServe(port, handler))
}
