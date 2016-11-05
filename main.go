package main

import (
	"github.com/ChimeraCoder/anaconda"

	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var (
		config *Config
		err    error
	)
	switch {
	case len(os.Args) == 1:
		log.Print("loading configuration from environment")
		config = LoadFromEnv()
	case len(os.Args) == 2 && os.Args[1] != "--help":
		log.Printf("loading configuration from %s", os.Args[1])
		config, err = LoadFromFile(os.Args[1])
	default:
		err = fmt.Errorf("Usage: %s [CONFIG]", os.Args[0])
	}
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	anaconda.SetConsumerKey(config.TwitterConsumerKey)
	anaconda.SetConsumerSecret(config.TwitterConsumerSecret)
	a, err := NewAdapter(config)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	log.Print("server started")
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT)
	<-ch
	log.Print("server shutting down...")
	a.Close()
}
