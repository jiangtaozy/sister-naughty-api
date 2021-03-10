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
  mux := http.NewServeMux()
  mux.HandleFunc("/mainImagesList", handle.MainImagesList)
  mux.Handle("/", http.FileServer(http.Dir("../image")))
  handler := cors.Default().Handler(mux)
  log.Fatal(http.ListenAndServe(port, handler))
}
