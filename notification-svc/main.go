package main

import (
	"log"
	"net/http"

	"github.com/Lokeshranjan8/notification-svc/handlers"
)

func main(){

	http.HandleFunc("/notify",handlers.SendNotification)



	log.Println("user service is running ")
	log.Fatal(http.ListenAndServe(":8081", nil))

}

