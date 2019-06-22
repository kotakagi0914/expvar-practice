package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/sheva0914/expvar-practice/metrics"
)

func main() {
	// Init metrics
	// metrics.InitMetrics()

	// Run http server
	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
	}()

	// Init metrics
	metrics.InitMetrics()

	for i := 0; i < 20; i++ {
		countFunc(i)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)
	select {
	case sig := <-sigChan:
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM:
			log.Println("end")
		}
	}
}

func countFunc(num int) {
	log.Println("num: ", num)

	metrics.Count1.Add(1)
	if num%3 == 0 {
		metrics.Count2.Add(1)
	}
}
