package routes

import (
  "encoding/json"
  "net/http"
  "fmt"
  "io/ioutil"

  "github.com/gorilla/mux"

  "github.com/YaminLi/jukebox/models"

)

func addSongRoutes(r *mux.Router) {
  // r.PathPrefix("/web/").Handler(http.StripPrefix("/web/", http.FileServer(http.Dir("web/"))))
  r.HandleFunc("/", homeHandler)
  r.HandleFunc("/songs", getSongsHandler).Methods("GET")
  r.HandleFunc("/songs", createSongHandler).Methods("POST")
  r.HandleFunc("/songs/{song_id}", updateSongHandler).Methods("POST")
  r.HandleFunc("/songs/{song_id}/upvote", upvoteSongHandler).Methods("POST") // TODO
  r.HandleFunc("/songs/{song_id}", deleteSongHandler).Methods("DELETE") // TODO
}

func updateSongHandler(w http.ResponseWriter, req *http.Request){
  //decode the request
  decoder := json.NewDecoder(req.Body)
  var new_song models.Song
  err := decoder.Decode(&new_song)
  if err != nil {
    panic(err)
    return
  }
  //get the song by song_id
  song_id := mux.Vars(req)["song_id"]
  var song models.Song
  models.DB.Where("ID = ?", song_id).First(&song)

  //update song info
  song.Title = new_song.Title
  song.Artist = new_song.Artist
  song.Album = new_song.Album

  //save the song
  models.DB.Save(&song)

  //return
  json.NewEncoder(w).Encode(&song)
}

func homeHandler(w http.ResponseWriter, req *http.Request){
  fmt.Println("Welcome to the home page")
  data, err := ioutil.ReadFile("web/song.html")
  if err == nil{
    w.Header().Set("Content-Type", "HTML")
    w.Write(data)
  } else {
    w.WriteHeader(404)
    w.Write([]byte("404 terrible"))
  }

}

// TODO
func deleteSongHandler(w http.ResponseWriter, req *http.Request) {
  song_id := mux.Vars(req)["song_id"]
  fmt.Println(song_id)

  // Get song by id
  var song models.Song
  models.DB.Where("ID = ?", song_id).First(&song)

  // Delete song
  models.DB.Delete(&song)
  //
  // Return response
  json.NewEncoder(w).Encode(&song)
}

// TODO
func upvoteSongHandler(w http.ResponseWriter, req *http.Request) {
  song_id := mux.Vars(req)["song_id"]
  fmt.Println(song_id)

  // Get song by id specified
  var song models.Song
  models.DB.Where("ID = ?", song_id).First(&song)

  // +1 to score
  song.Votes += 1
  //
  // Save song
  models.DB.Save(&song)

  // Return song in response
  json.NewEncoder(w).Encode(&song)

}

func getSongsHandler(w http.ResponseWriter, req *http.Request) {
  // Get songs
  var songs []models.Song
  models.DB.Order("votes desc").Find(&songs)

  // Create response
  // data := make(map[string]interface{})
  // data["songs"] = songs
  // json.NewEncoder(w).Encode(&data)
  bytes, e := json.Marshal(&songs)
  if e != nil {
    http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
    return
  }
  w.Header().Set("Content-Type", "application/json")
  w.Write(bytes)
}

func createSongHandler(w http.ResponseWriter, req *http.Request) {

  // Get JSON request data
  decoder := json.NewDecoder(req.Body)
  var song models.Song
  err := decoder.Decode(&song)
  if err != nil {
    panic(err)
    return
  }

  // Initialize votes value
  song.Votes = 1

  // Save song
  models.DB.Create(&song)

  // Create response
  // json.NewEncoder(w).Encode(&song)
  bytes, e := json.Marshal(&song)
	if e != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}
  w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}
