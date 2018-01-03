package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

var (
	cfg    Config
	client *redis.Client
	logger Logger
)

func startHttpServer(port string) *http.Server {
	srv := &http.Server{Addr: ":" + port}

	r := mux.NewRouter()
	r.HandleFunc("/", GetHandler)
	r.HandleFunc("/create/{id:[0-9]+}", CreateHandler)
	r.HandleFunc("/update/{id:[0-9]+}", UpdateHandler)
	r.HandleFunc("/delete/{id:[0-9]+}", DeleteHandler)

	http.Handle("/", r)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Error("Httpserver: ListenAndServe() error: " + err.Error())
		}
	}()

	return srv
}

func main() {
	var config = cfg.Init("./config.json")
	redisHost := config.Get("redis", "host")
	redisPort := config.Get("redis", "port")
	loglevel := config.Get("log", "level")
	logger.Level = loglevel

	client = redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer client.Close()

	result, err := client.Ping().Result()
	if err != nil {
		logger.Error("Cannot connect to Redis: " + err.Error())
		os.Exit(-1)
	}

	logger.Info("Info: Response from server: " + result)

	srv := startHttpServer("9000")
	logger.Info("main: starting server on port " + srv.Addr)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	exit_chan := make(chan int)

	go func() {
		for {
			s := <-c
			switch s {
			case syscall.SIGHUP:
				exit_chan <- 0
			case syscall.SIGINT:
				exit_chan <- 0
			case syscall.SIGTERM:
				exit_chan <- 0
			case syscall.SIGQUIT:
				exit_chan <- 0
			default:
				exit_chan <- 1
			}
		}
	}()

	code := <-exit_chan

	if err := srv.Shutdown(nil); err != nil {
		panic(err)
	}
	logger.Info("main: server shutdown successfully")
	os.Exit(code)
}
