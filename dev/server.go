package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/net/websocket"

	//"github.com/gorilla/websocket"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//
const DB_DNS = "mongodb://localhost/test"
const DB_NAME = "test"

// 単純な構造体
type Like struct {
	Timestamp time.Time `bson:"time"`
}
type Likes struct {
	Likes []Like
}

type Comment struct {
	Content   string    `bson:"content"`
	Timestamp time.Time `bson:"time"`
}

type Comments struct {
	Comments []Comment
}

// main
func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/likes", getLikesHandler).Methods("GET")
	router.HandleFunc("/api/likes", postLikeHandler).Methods("POST")
	router.HandleFunc("/api/comments", getCommentsHandler).Methods("GET")
	router.HandleFunc("/api/comments", postCommentsHandler).Methods("POST")
	http.Handle("/echo", websocket.Handler(wsEchoHandler))
	router.PathPrefix("/src/").Handler(http.StripPrefix("/src/", http.FileServer(http.Dir("src/"))))
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("static"))))

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}

func getLikesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	session, _ := mgo.Dial(DB_DNS)
	defer session.Close()
	db := session.DB(DB_NAME)

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

	session, _ := mgo.Dial(DB_DNS)
	defer session.Close()
	db := session.DB(DB_NAME)

	like := &Like{
		Timestamp: time.Now(),
	}
	col := db.C("like")
	col.Insert(like)

	j, _ := json.Marshal(like)
	w.Write(j)
}

func getCommentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	session, _ := mgo.Dial(DB_DNS)
	defer session.Close()
	db := session.DB(DB_NAME)

	var result []Comment
	err := db.C("comment").Find(bson.M{}).All(&result)
	if err != nil {
		panic(err)
	}

	fmt.Printf("", result)
	print(result)

	comments := Comments{
		Comments: result,
	}
	j, _ := json.Marshal(comments)
	w.Write(j)
}

func postCommentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	session, _ := mgo.Dial(DB_DNS)
	defer session.Close()
	db := session.DB(DB_NAME)

	content := r.FormValue("content")
	comment := &Comment{
		Content:   content,
		Timestamp: time.Now(),
	}
	col := db.C("comment")
	col.Insert(comment)

	j, _ := json.Marshal(comment)
	w.Write(j)
}

func wsEchoHandler(ws *websocket.Conn) {
	log.Println("connect")
	io.Copy(ws, ws)
}
