package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"register-api/config"
	"register-api/service"
)

const(
	createPath = "/create"
)

func main() {
	config.InitConfig()

	//initDb
	db := config.InitDB()
	defer db.Close()

	//init producer
	p := config.InitProducer()
	defer p.Close()

	//init handler
	h := service.NewHandler(db, p)

	r := initRouter(h)
	panic(http.ListenAndServe(":8080", r))
}

func initRouter(h *service.Handler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc(createPath, h.Create).Methods(http.MethodPost)
	return r
}