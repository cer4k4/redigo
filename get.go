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
	var r []int
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	for i:=1368;i<=1420;i++{
		res , _ := client.HGet(ctx,"message:"+strconv.Itoa(i),"id").Result()
		ints , _ := strconv.Atoi(res)
		r = append(r,ints)
		if i == 1420{
			fmt.Println(i)
		}
	}
	for i:=1; i<len(r);i++{
		fmt.Println(r[i])
		fmt.Printf("%T",client)
	}
}


