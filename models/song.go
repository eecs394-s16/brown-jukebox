package models

import (
  "github.com/jinzhu/gorm"
)

type Song struct {
  gorm.Model
  Title  string `json:"title"`
  Artist string `json:"artist"`
  Album  string `json:"album"`

  Votes  int    `json:"score"`
}
