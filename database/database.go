/*
 * Maintained by jemo from 2021.2.26 to now
 * Created by jemo on 2020.2.26 17:21:47
 * database
 */

package database

import (
  "database/sql"
  "log"
  _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
  ConnectDB()
  //execSQL(createUser)
  //execSQL(createImageBrowseRecord)
}

func ConnectDB() {
  var err error
  DB, err = sql.Open("mysql", DbUrl)
  if err != nil {
    log.Println("db-initdb-open-error: ", err)
  }
}

func execSQL(sqlStmt string) {
  log.Println("db exec sql: ", sqlStmt)
  stmt , err := DB.Prepare(sqlStmt)
  if err != nil {
    log.Println("database.go-execSQL-prepare-error: ", err)
  }
  defer stmt.Close()
  _, err = stmt.Exec()
  if err != nil {
    log.Println("database.go-execSQL-exec-error: ", err)
  }
}
