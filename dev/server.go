package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// 単純な構造体
type Like struct {
	Timestamp time.Time `bson:"time"`
}
type Likes struct {
	Likes []Like
}

// main
func main() {
	router := mux.NewRouter()
	http.Handle("/src/", http.StripPrefix("/src/", http.FileServer(http.Dir("src/"))))
	router.Handle("/", http.FileServer(http.Dir("static")))
	router.HandleFunc("/api/likes", getLikesHandler).Methods("GET")
	router.HandleFunc("/api/likes", postLikeHandler).Methods("POST")

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}

func getLikesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	session, _ := mgo.Dial("mongodb://localhost/test")
	defer session.Close()
	db := session.DB("test")

	//log.Print(db)

	var result []Like
	err := db.C("like").Find(bson.M{}).All(&result)
	if err != nil {
		panic(err)
	}
	likes := Likes{
		Likes: result,
	}
	j, _ := json.Marshal(likes)
	w.Write(j)
}

func postLikeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	session, _ := mgo.Dial("mongodb://localhost/test")
	defer session.Close()
	db := session.DB("test")

	like := &Like{
		Timestamp: time.Now(),
	}
	col := db.C("like")
	col.Insert(like)

	j, _ := json.Marshal(like)
	w.Write(j)
}
