package main

import (
	"log"
	"net/http"

	"github.com/Lokeshranjan8/user-svc/handlers"
	"github.com/Lokeshranjan8/user-svc/utils"
)

func main(){
	utils.InitDB()

	http.HandleFunc("/createuser",handlers.Create_user)
	http.HandleFunc("/getuser",handlers.Get_user)
	http.HandleFunc("/getusers",handlers.Getall_user)


	log.Println("user service is running ")
	log.Fatal(http.ListenAndServe(":8081", nil))

}

