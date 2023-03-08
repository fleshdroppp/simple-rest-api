package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"simple-rest-api/internal/config"
	"simple-rest-api/internal/user"
	"simple-rest-api/pkg/logging"
	"time"
)

func main() {
	l := logging.GetLogger()
	l.Info("creating router...")
	router := httprouter.New()

	cfg := config.GetConfig()

	//cfgMongoDB := cfg.MongoDB
	//client, err := mongodb.NewClient(context.Background(), cfgMongoDB.Host, cfgMongoDB.Port, cfgMongoDB.Username,
	//	cfgMongoDB.Password, cfgMongoDB.Database, cfgMongoDB.AuthDB)
	//if err != nil {
	//	panic(err)
	//}
	//storage := db.NewStorage(client, cfg.MongoDB.Collection, l)

	l.Info("registering user handler...")
	handler := user.NewHandler(l)
	handler.Register(router)

	start(router, cfg)
}

func start(router *httprouter.Router, cfg *config.Config) {
	l := logging.GetLogger()

	var listener net.Listener
	var listenErr error

	if cfg.Listen.Type == "sock" {
		l.Info("detecting app path...")

		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			l.Fatal(err)
		}

		l.Info("creating socket...")
		socketPath := path.Join(appDir, "app.sock")

		l.Info("creating unix socket...")
		listener, listenErr = net.Listen("unix", socketPath)
		l.Infof("server is up on unix socket: %s", socketPath)

	} else {
		l.Info("starting tcp server...")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
		l.Infof("server is up on %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)
	}

	if listenErr != nil {
		l.Fatal(listenErr)
	}

	srv := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	l.Fatal(srv.Serve(listener))
}
