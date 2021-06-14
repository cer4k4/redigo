package main

import (
	"time"
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
	var data []Message
	t := time.Tick(1 * time.Minute)
	for next := range t {
	for i:= 1368;i<=1420;i++{
	data = append(data,GetMessage(i))
}
for d:=0;d<len(data);d++{
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	_ = client.HSet(ctx,"message:"+strconv.FormatUint(uint64(data[d].ID),10),
	"sender",data[d].Sender,
	"receiver",data[d].Receiver,
	"chatroom",data[d].ChatRoomID,
	"message",data[d].Message,
	"type",data[d].Type,
	"file","",
	"create_at",
	data[d].CreatedAt,"update_at",
	data[d].UpdatedAt,"delete_at","")
	/*	if err != nil{
		log.Println(err)
		}
	*/
//	fmt.Println("affected : ",data[d].ID)
	}
	fmt.Println(next)
}
}


func GetMessage(i int) Message{
	Connection()
	defer DB.Close()
	var message Message
	DB.Find(&message,i)
	return message
}
func GetAllMessage() interface{} {
	Connection()
	messages := []Message{}
	DB.Find(&messages)
	return messages
}
