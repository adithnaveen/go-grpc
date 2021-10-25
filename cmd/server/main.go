package main

import (
	"sync"

	"github.com/jyotishp/go-orders/pkg/http"
)

func main() {
	go http.StartGRPC()
	go http.StartHTTP()

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
