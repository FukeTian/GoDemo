package main

import (
	"UserTest/controller"
	"UserTest/framework"
	"fmt"
	"net/http"
	"time"
)

func main() {

	framework.InitDB()
	framework.CreateTable()

	server := &http.Server{
		Addr:        ":8089",
		Handler:     framework.Router,
		ReadTimeout: 5 * time.Second,
	}
	RegiterRouter(framework.Router)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("start server error")
	}
	fmt.Println("start server success")
}

// RegiterRouter 注册路由
func RegiterRouter(handler *framework.RouterHandler) {
	new(controller.UserConterller).Router(handler)
}
