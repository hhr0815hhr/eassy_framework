package gate

import (
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
)

func Run(port string)  {
	http.Handle("/eassy",websocket.Handler(handler))
	err := http.ListenAndServe("0.0.0.0:"+port,nil)
	if err != nil {
		panic(err)
	}
}

func handler(ws *websocket.Conn)  {
	data := ws.Request().URL.Query().Get("data")
	fmt.Println("data:", data)

}


