package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type box struct {
	Size int     `json:"size"`
	X    int     `json:"x"`
	Y    int     `json:"y"`
	Z    int     `json:"z"`
	R    float64 `json:"r"`
	G    float64 `json:"g"`
	B    float64 `json:"b"`
}

func randomBox() box {
	var b box
	b.Size = rand.Intn(5) + 1
	b.X, b.Y, b.Z = -rand.Intn(100)+1, -rand.Intn(100)+1, -rand.Intn(100)+1
	b.R, b.G, b.B = float64(rand.Intn(254)+1)/255, float64(rand.Intn(254)+1)/255, float64(rand.Intn(254)+1)/255
	return b
}

func main() {
	log.Println("starting web server on port 8080")

	// handler for default website
	http.Handle("/", http.FileServer(http.Dir("../client")))

	// handler for websocket
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil) // upgrade http connection to websocket connection
		if err != nil {
			log.Println(err)
		}

		for {
			message := randomBox()
			if err := conn.WriteJSON(message); err != nil {
				log.Println("could not send message to client")
			}
			log.Println("sent message")
			time.Sleep(time.Millisecond)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
