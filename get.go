package main

import (
	"context"
	"fmt"
	"log"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Message struct {
	Sender     string `json:"sender"`
	Receiver   string `json:"receiver"`
	ChatRoomID int    `json:"chat_room_id"`
	Message    string `json:"message"`
	Type       string `json:"type"`
	File       string `json:"file"`
	gorm.Model
}

var DB *gorm.DB
var DBErr error

func Connection() {
	DB, DBErr = gorm.Open("mysql", "golang:golang123@(localhost)/golang_test?charset=utf8&parseTime=True&loc=Local")
	if DBErr != nil {
		log.Println(DBErr)
	}
}
func main(){
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	r , err := client.HGet(ctx,"message:152","message").Result()
	if err != nil{
		log.Println(err)
	}
	fmt.Println(r)
//	data := GetAllMessage()
//	fmt.Println(data)
}


