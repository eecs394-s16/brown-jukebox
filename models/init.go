package models

import (
  "os"

  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func init() {
  var err interface{}
  DB, err = gorm.Open("postgres", "user=" + os.Getenv("PG_USER") + " dbname=jukebox password=" + os.Getenv("PG_PASSWORD") + " sslmode=disable")
  if err != nil {
    panic(err)
  }
}
