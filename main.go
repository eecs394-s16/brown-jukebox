package main

import (
  "github.com/YaminLi/jukebox/routes"
)

func main() {
  n  := routes.GetRouter()
  n.Run(":3000")
}
