package main

import (
	"fmt"
	"proxy/proxy"
	"proxy/service"
)

func main() {
	var server service.Server

	server = proxy.NewServiceProxy()

	fmt.Println(server.HandleRequest(service.APP_STATUS_URL, service.GET))
	fmt.Println(server.HandleRequest(service.APP_STATUS_URL, service.GET))
	fmt.Println(server.HandleRequest(service.APP_STATUS_URL, service.GET))

	fmt.Println(server.HandleRequest(service.CREATE_USER_URL, service.POST))
	fmt.Println(server.HandleRequest(service.CREATE_USER_URL, service.GET))
}
