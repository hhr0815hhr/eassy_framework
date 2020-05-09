package gate

import (
	"golang.org/x/net/websocket"
	"net/http"
)

/** gate节点采用websocket
 *
 */
func Run(port string) {
	http.Handle("/eassy", websocket.Handler(handler))
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

func handler(ws *websocket.Conn) {
	cli := CliManager.Connect(ws)
	cli.RecvData()

	//data := ws.Request().URL.Query().Get("data")
	//fmt.Println("data:", data)
	//for {
	//	//var data []byte
	//	//
	//	//if err = websocket.Message.Receive(ws, data); err != nil {
	//	//	fmt.Println(err)
	//	//	continue
	//	//}
	//
	//}
}
