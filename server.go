package main

import (
	"fmt"
	"net/http"
	"sms/services"
)

func main() {

	var connParams = services.NewConnParameters("config.properties")

	vanso := &services.VansoConnection{}

	var conn services.Connection

	conn = services.NewConnection(vanso, connParams.GetParameters())

	fmt.Println("Server started")
	http.HandleFunc("/", conn.ServeHTTP)
	http.ListenAndServe(":8080", nil)

}
