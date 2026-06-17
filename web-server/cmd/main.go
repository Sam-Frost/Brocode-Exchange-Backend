package main

import (
	"sync"

	"github.com/Sam-Frost/web-server/internal"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(2)

	go internal.StartServer()
	go internal.CreateGrpcClient()

	wg.Wait()
}
