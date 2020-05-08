package login

import (
	"game_framework/src/eassy/login/accounts"
	"github.com/gorilla/mux"
	"net/http"
)

func Run(port string) {
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/checkToken", checkHandler)
	http.Handle("/", r)
	accounts.InitDbConn()
	http.ListenAndServe(":"+port, nil)
}
