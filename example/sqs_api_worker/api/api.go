package main

import (
	"context"
	"flag"
	"log"

	"github.com/lilien1010/taskq/example/api_worker/sqs_api_worker"
)

func main() {
	flag.Parse()

	go sqs_api_worker.LogStats()

	go func() {
		for {
			err := sqs_api_worker.MainQueue.Add(sqs_api_worker.CountTask.WithArgs(context.Background()))
			if err != nil {
				log.Fatal(err)
			}
			sqs_api_worker.IncrLocalCounter()
		}
	}()

	sig := sqs_api_worker.WaitSignal()
	log.Println(sig.String())
}
