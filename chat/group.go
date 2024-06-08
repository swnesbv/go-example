package chat

import (
	"time"
	"log"
	"fmt"
    //"os"
	"net/http"
	
	"go_authentication/options"
	"go_authentication/authtoken"
	"github.com/gorilla/websocket"
)


var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		fmt.Println("groups upgrader")
		return true
	},
}

var (
	clients    	  = make(map[*websocket.Conn]bool)
	register      = make(chan *websocket.Conn)
	unregister    = make(chan *websocket.Conn)
	broadcast     = make(chan Message)
)


// go run
func groupCh() {
	for {
		select {

		case client := <-register:
			fmt.Println("Add groups client")
			clients[client] = true

		case client := <-unregister:
			fmt.Println("Del groups client")
			delete(clients, client)
			client.Close()

		case message := <-broadcast:
			fmt.Println("msg..", message)

			for client := range clients {

				if err := client.WriteJSON(message); err != nil {

					fmt.Println("groups message err:", err)
					delete(clients, client)
					client.Close()
				}
			}
		}
	}
}


func GrMsg(w http.ResponseWriter, r *http.Request) {

    cls := authtoken.WhoisWho(w,r)
    
    id,err := options.IdUrl(w,r)
    if err != nil { return }

	conn,err := upgrade.Upgrade(w,r, nil)
	fmt.Println("Group upgrade..", cls.Email)
	fmt.Println("Group id..", id)
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		fmt.Println("msg groups user close..", cls.Email)
		unregister <- conn
		conn.Close()
	}()

	register <- conn

	str := "INSERT INTO msggroups (coming, owner, to_group, completed, created_at) VALUES ($1,$2,$3,$4,$5)"

	for {
		var message Message
		err := conn.ReadJSON(&message)
		if err != nil {
			fmt.Println("err msg groups message", err)
			break
		}
		fmt.Println("msg groups message..", message.Message)

    	owner := cls.User_id

    	_,err = db.Exec(str, message.Message,owner,id,true,time.Now())
		if err != nil {
			fmt.Println("err msg groups Exec..", err)
			break
		}

		broadcast <- message
	}
}
