package main

import (
	"fmt"
	"gsc_rest/handler"
	"gsc_rest/repository"
	"net/http"
)

func main() {

	InitService()

	defer repository.DB.Close()

	// service.SendCodeVerification("rasyid08mufara@gmail.com")
	// service.SendCodeVerification("sibangkekunyuk@gmail.com")

	handler.RouteInit()

	fmt.Println("starting web server at http://localhost:5580/")
	http.ListenAndServe(":5580", nil)
}

// lsof -i:5580
// kill pid
