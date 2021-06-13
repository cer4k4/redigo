package main

import (
	"strconv"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type Message struct {
	Sender     string `json:"sender"`
	Receiver   string `json:"receiver"`
	ChatRoomID int    `json:"chat_room_id"`
	Message    string `json:"message"`
	Type       string `json:"type"`
	File       string `json:"file"`
}


func main(){
	var r []string
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	for i:=1368;i<=1420;i++{
		res , _ := client.HGet(ctx,"message:"+strconv.Itoa(i),"message").Result()
r = append(r,res)
}
for i:=0; i<len(r);i++{
	fmt.Println(r[i])
}
}


