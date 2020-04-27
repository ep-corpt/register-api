package service

import (
	"encoding/json"
	"errors"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"register-api/common"
	"register-api/entity"
	"register-api/model"
	"register-api/rule"
)

const (
	RegisterSuccessfully = "Register Successfully."
	InternalServerError = "Internal server error."
)

type CreateTask struct{
	db *gorm.DB
	req *model.MessageModel
	reqJ []byte
	p *kafka.Producer
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request){
	var t CreateTask
	t.initTask(h)

	//init request
    if err:=t.initReq(r); err!=nil{
		common.RespErr(w, InvalidRequest)
		return
	}

	//validate request message
	if err:=t.valid(); err!=nil{
		common.RespErr(w, err.Error())
		return
	}

	//validate duplicate username/email
	if err:=t.duplicate(); err!=nil{
		common.RespErr(w, err.Error())
		return
	}

	//product message to kafka
	if err:=t.produce(); err!=nil{
		common.RespErr(w, InternalServerError)
		return
	}

	common.RespSuccess(w, RegisterSuccessfully)
}

func (t *CreateTask) initTask(h *Handler){
	t.db = h.db
	t.p = h.p
}

func (t *CreateTask) initReq(r *http.Request) error {
	if err := json.NewDecoder(r.Body).Decode(&t.req); err!= nil {
		return err
	}
	reqJ, err := json.Marshal(&t.req)
	if err!=nil {
		return err
	}
	t.reqJ = reqJ
	return nil
}

func (t *CreateTask) valid() error {
	if err:= validation.ValidateStruct(&t.req.UserDetail,
		//User detail
		validation.Field(&t.req.UserDetail.FirstName, rule.FirstName...),
		validation.Field(&t.req.UserDetail.LastName, rule.LastName...),
		validation.Field(&t.req.UserDetail.Email, rule.Email...)); err!=nil{
		return err
	}

	//Company detail
	if err:= validation.ValidateStruct(&t.req.CompanyDetail,
		validation.Field(&t.req.CompanyDetail.CompanyName, rule.CompanyName...));err!=nil{
		return err
	}

	//Credential detail
	if err:=validation.ValidateStruct(&t.req.CredentialDetail,
		validation.Field(&t.req.CredentialDetail.Username, rule.UserName...),
		validation.Field(&t.req.CredentialDetail.Password, rule.Password...));err!=nil{
		return err
	}
	return nil
}

func (t *CreateTask) duplicate() error {
	r := t.db.Where("username = ? OR email = ?", t.req.CredentialDetail.Username, t.req.UserDetail.Email).Find(&entity.CredentialInfo{})
	log.Printf("credentialInfo %v\n", r.Value)
	if r.Value == nil{
		return errors.New(usernameOrEmailIsDuplicate)
	}
	return nil
}

func (t *CreateTask) produce() error{
	topic := viper.GetString("kafka.topic")
	return t.p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value: t.reqJ,
	}, nil)
}