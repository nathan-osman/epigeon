package main

import (
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
	case len(os.Args) == 0:
		config = LoadFromEnv()
	case len(os.Args) == 1 && os.Args[1] != "--help":
		config, err = LoadFromFile(os.Args[1])
	default:
		err = fmt.Errorf("Usage: %s [CONFIG]", os.Args[0])
	}
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	a, err := NewAdapter(config)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT)
	<-ch
	a.Close()
}
