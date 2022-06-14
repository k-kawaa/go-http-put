package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/fatih/color"
)

func setup(){
	color.Green("Start Server")
	server := &http.Server{
		Addr:              ":80",
		Handler:           http.HandlerFunc(handler),
	}
	go StartServer(server)
	quit := make(chan os.Signal)
	signal.Notify(quit,os.Interrupt)
	<- quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	server.Shutdown(ctx)
	color.Yellow("[Server]Server is stop")
}

func handler(w http.ResponseWriter, r *http.Request) {
	
    switch r.Method {
	    case "PUT":
	      length := r.ContentLength
	      fmt.Println("length:",length)
	      body := make([]byte, length)
	      r.Body.Read(body)
	      fmt.Println(string(body))
	    default:
	      return
    }

}