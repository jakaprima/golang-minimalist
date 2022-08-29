package main

import (
	service "belajar/11-service"
	"fmt"
)

func main() {
	// init
	var db []*service.User
	userServiceInstance := service.NewUserService(db)
	fmt.Println(userServiceInstance)
	// register
	response := userServiceInstance.Register(&service.User{Nama: "Jaka Prima Maulana"})
	fmt.Println(response)
	response2 := userServiceInstance.Register(&service.User{Nama: "budi"})
	fmt.Println(response2)

	// get user
	responseGetUser := userServiceInstance.GetUser()
	fmt.Println("---------------- HASIL GET USER")
	for _, v := range responseGetUser {
		fmt.Println(v.Nama)
	}
}
