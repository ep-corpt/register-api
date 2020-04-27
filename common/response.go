package common

import (
	"encoding/json"
	"log"
	"net/http"
	"register-api/model"
)

const (
	StatusSuccess = "00000"
	StatusFailed = "00001"
)

func RespErr(w http.ResponseWriter, m string){
	w.Header().Set("Content-type","application/json")
	log.Printf("Err %v\n",m)
	if err := json.NewEncoder(w).Encode(&model.Response{Status: StatusFailed, Desc: m}); err!=nil{
		log.Panicln(err)
	}
}

func RespSuccess(w http.ResponseWriter, m string){
	w.Header().Set("Content-type","application/json")
	if err := json.NewEncoder(w).Encode(&model.Response{Status: StatusSuccess, Desc: m}); err!=nil{
		log.Panicln(err)
	}
}

