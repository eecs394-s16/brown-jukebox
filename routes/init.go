package routes

import (
  "net/http"

  "github.com/gorilla/mux"

  "github.com/codegangsta/negroni"
)

func GetRouter() *negroni.Negroni {
  n  := negroni.New()
  n.Use(negroni.HandlerFunc(configureResponseMiddleware))

  r  := mux.NewRouter()
  addSongRoutes(r)

  n.UseHandler(r)
  return n
}

func configureResponseMiddleware(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
  w.Header().Set("Content-Type", "application/json")
  next(w, req)
}
