package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"simple-rest-api/internal/user"
	"time"
)

func main() {
	log.Println("creating router...")
	router := httprouter.New()

	log.Println("registering user handler...")
	handler := user.NewHandler()
	handler.Register(router)

	log.Println("starting server")
	start(router)
}

func start(router *httprouter.Router) {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}

	srv := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatalln(srv.Serve(listener))
}
