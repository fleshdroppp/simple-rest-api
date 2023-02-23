package main

import (
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"simple-rest-api/internal/user"
	"simple-rest-api/pkg/logging"
	"time"
)

func main() {
	l := logging.GetLogger()
	l.Info("creating router...")
	router := httprouter.New()

	l.Info("registering user handler...")
	handler := user.NewHandler(l)
	handler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	l := logging.GetLogger()
	l.Info("starting server")
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}

	srv := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	l.Fatal(srv.Serve(listener))
}
