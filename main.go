package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
)

type User struct{
	Name string `redis:"name"`
	Age int     `redis:"age"`
}



func main() {

	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
/*	for i:=0;i<=3;i++{
		val, err := client.HGetAll(ctx,"message:"+strconv.Itoa(i)).Result()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(val)
		if val["chatrooms"] == "1"{
			chatroom, err := client.HGetAll(ctx,"chatrooms:1").Result()
			if err != nil {
				fmt.Println(err)
			}
		fmt.Println("chatroom : ",chatroom)
		}
	}
*/

////////////////////////////////////////////////////////////////////////////////////
	chatroom, err := client.HGetAll(ctx,"chatrooms:1").Result()
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(chatroom)
	for i:=1;i<=3;i++{
		val, err := client.HGetAll(ctx,"message:"+strconv.Itoa(i)).Result()
		if err != nil{
			fmt.Println(err)
		}
		if i > 1{
			fmt.Println("message:"+strconv.Itoa(i),"sender:",val["sender"],"resiver:",val["resiver"],"\n",val["message"])
		}
	}
}


