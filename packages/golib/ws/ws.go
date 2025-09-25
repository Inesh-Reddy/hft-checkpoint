package ws

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type ws struct{
	websocket *websocket.Conn
}

func ConnnectToWs(WsUrl string) *ws{
	conn,_,err:=websocket.DefaultDialer.Dial(WsUrl, nil)
	if(err != nil){
		fmt.Println("Error while dialing to websocket....")
	}
	return &ws{websocket: conn}
}

