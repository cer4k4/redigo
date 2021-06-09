package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

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
	data := GetMessage()
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	bytemessage := []byte(data.Message)
	fmt.Println(bytemessage)
	for i:=range bytemessage {
		fmt.Printf("%T",bytemessage[i])
}
	result, err := client.HSet(ctx,"message:"+strconv.Itoa(data.ChatRoomID),"sender",data.Sender,"receiver",data.Receiver,
	"chatroom",data.ChatRoomID,"message",bytemessage,"type",data.Type,"file","","create_at",data.CreatedAt,"update_at",
	data.UpdatedAt,"delete_at","").Result()
	if err != nil{
		log.Println(err)
	}
	fmt.Println(result)
	fmt.Println(data.Receiver)

//	data := GetAllMessage()
//	fmt.Println(data)
}


func GetMessage() Message{
	Connection()
	var message Message
	DB.First(&message,1370)
	return message
}
func GetAllMessage() interface{} {
	Connection()
	messages := []Message{}
	DB.Find(&messages)
	return messages
}

