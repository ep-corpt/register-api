package config

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
	"log"
	"register-api/entity"
)

const(
	config = "config"
	dot = "."
	postgres = "postgres"
)

func InitConfig(){
	viper.SetConfigName(config)
	viper.AddConfigPath(dot)
	if err:=viper.ReadInConfig(); err!=nil{
		log.Panicln(err)
	}
}

func InitDB() *gorm.DB{
	log.Print(viper.Get("db"))
	db, err := gorm.Open(postgres, viper.Get("db"))
	if err!=nil{
		log.Panicln(err)
	}
	db.SingularTable(true)
	db.AutoMigrate(&entity.CredentialInfo{})
	return db
}

func InitProducer() *kafka.Producer{
	p, err :=kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": viper.GetString("kafka.host"),
	})
	if err!=nil{
		log.Panicln(err)
	}
	return p
}