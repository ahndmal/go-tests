package net

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var origins = []string{"http://127.0.0.1:8083", "http://localhost:9999"}
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	//CheckOrigin: func(r *http.Request) bool {
	//	return  true
	//},
	CheckOrigin: func(r *http.Request) bool {
		var origin = r.Header.Get("origin")
		for _, allowOrigin := range origins {
			if origin == allowOrigin {
				return true
			}
		}
		return false
	},
}

type MySocket struct {
}

func (ms MySocket) Cast() {
	http.HandleFunc("/echo", func(writer http.ResponseWriter, req *http.Request) {
		conn, err := upgrader.Upgrade(writer, req, nil)
		if err != nil {
			log.Panicln(err)
		}
		for {
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				log.Panicln(err)
			}
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))
			if err = conn.WriteMessage(msgType, msg); err != nil {
				log.Panicln(err)
			}
		}
	})
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	http.ServeFile(w, r, "websockets.html")
	//})

	err3 := http.ListenAndServe(":8083", nil)
	if err3 != nil {
		log.Panicln(err3)
	}
}
