package main

import (
//	"time"
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
var client *redis.Client

//Connetions
func ConnectionDB() {
	DB, DBErr = gorm.Open("mysql", "golang:golang123@(localhost)/golang_test?charset=utf8&parseTime=True&loc=Local")
	if DBErr != nil {
		log.Println(DBErr)
	}

}
func ConnectionRedis() {
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
}

func main() {
	ctx := context.Background()
	ConnectionDB()
	ConnectionRedis()
	defer DB.Close()
	//var datards []Message
//	var data []Message
//	t := time.Tick( 30 * time.Second)
//	for next := range t {
var a []string
		//get from redis
		for i:= 1368;i<=1390;i++{
				//data = append(data,GetMessageFromDB(i,DB))
				res, _ := client.HMGet(ctx,"message:"+strconv.Itoa(i),"message","id","file").Result()
				if res != nil && res[0] != nil || res[2] != nil{
					fmt.Println(res[1])
					a = res{"id"}
				}
			}
			fmt.Println(a)


		//append to redis
		/*for d:=0;d<len(data);d++{
			if data[d].File == ""{
				data[d].File = ""
			}else{
				data[d].File = data[d].File
			}
			if data[d].Message == ""{
				data[d].Message = ""
			}else{
				data[d].Message = data[d].Message
			}
			_ = client.HSet(ctx,"message:"+strconv.FormatUint(uint64(data[d].ID),10),
			"sender",data[d].Sender,
			"receiver",data[d].Receiver,
			"chatroom",data[d].ChatRoomID,
			"message",data[d].Message,
			"type",data[d].Type,
			"file",data[d].File,
			"id",int(data[d].ID),
			"create_at",data[d].CreatedAt,
			"update_at",data[d].UpdatedAt,
			"delete_at","")
		}*/
	//	fmt.Println(next)
//	}
}






























// Query's
func GetMessageFromDB(i int, db *gorm.DB) Message{
	var message Message
	db.Find(&message,i)
	return message
}
/*
func GetMessageFromRedis(i int,client *redis.Client) Message{
	var r []Message
	ctx := context.Background()
	for i:=1368;i<=1420;i++{
		res, _ := client.HGetAll(ctx,"message:"+strconv.Itoa(i)).Result()
		r = append(r,r[i].Message)
		if i == 1420{
			fmt.Println(i)
		}
		for i:=1;i<len(r);i++{
			fmt.Println(r[i])
		}
	}
	return r
}*/

