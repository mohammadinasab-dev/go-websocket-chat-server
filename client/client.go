package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"golang.org/x/net/websocket"
)

type Message struct {
	EventName    string      `json:"eventName"`
	EventPayload interface{} `json:"eventPayload"`
}

type Payload struct {
	// UserID  string `json:"userID"`
	UserName string `json:"username"`
	Message  string `json:"message"`
}

func main() {
	origin := "http://localhost/"
	conn, err := websocket.Dial("ws://localhost:8000/chat/"+CreateDemoIp(), "", origin)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer conn.Close()

	go receive(conn)

	send(conn)

}

func CreateDemoIp() string {
	//192.168.1.1
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(256)
	return fmt.Sprintf("user%d", id)
}

func receive(conn *websocket.Conn) {
	for {
		var m Message
		err := websocket.JSON.Receive(conn, &m)
		if err != nil {
			log.Fatalln("Error in Recieve Data :", err)
			continue
		}
		fmt.Println("Message from Server account: ", conn.RemoteAddr().String())
		fmt.Println("Message from Server: ", m.EventPayload)
	}
}

func send(conn *websocket.Conn) {

	var payload Payload
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()
		decoder := json.NewDecoder(bytes.NewReader([]byte(text)))
		decoderErr := decoder.Decode(&payload)
		if decoderErr != nil {
			fmt.Println("Error in Send Data, ", decoderErr)
			continue
		}
		log.Println(text)
		mp := make(map[string]interface{})
		// mp["userID"] = payload.UserID
		mp["username"] = payload.UserName
		mp["message"] = payload.Message
		m := Message{
			EventName:    "message",
			EventPayload: mp,
		}
		err := websocket.JSON.Send(conn, m)
		if err != nil {
			fmt.Println("Error in Send Data, ", err)
			continue
		}
	}

}
