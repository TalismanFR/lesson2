package main

import (
	"encoding/json"
	"fmt"
	user2 "lesson2/user"
)

func main() {
	client := user2.User{"1@yr.ru", 32, true}
	client1 := user2.User{Age: 32}

	fmt.Printf("client %T %v\n", client, client)
	fmt.Printf("client1 %T %v\n", client1, client1)
	fmt.Println("email=", client.Email)

	fio := user2.NewFio("first", "last", "other")

	fmt.Println(fio, fio.FirstName())

	fio.SetFistName("newNAme")
	fmt.Println(fio.FirstName())

	g := user2.NewGeo(3, 2)

	fmt.Println(g, *g.Lat)

	gg := user2.NewGeo(4, 3)

	fmt.Println((*gg).Lat)

	fmt.Println("===================")

	bytes, _ := json.Marshal(client)

	fmt.Println(string(bytes))

	client2 := &user2.User{}
	json.Unmarshal(bytes, client2)
	fmt.Println(client2)

	bytes, _ = json.Marshal(fio)
	fmt.Println(string(bytes))
}
