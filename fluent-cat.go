package main

import (
	"fmt"
	"github.com/fluent/fluent-logger-golang/fluent"
	"sync"
	"time"
)

func fluentCat(host string, port int, requests int, concurrency int, tag string, message string) {
	requestsPerRoutine := requests / concurrency
	logger, err := fluent.New(fluent.Config{FluentHost: host, FluentPort: port, SubSecondPrecision: true})
	if err != nil {
		fmt.Println(err)
	}
	defer logger.Close()

	wg := new(sync.WaitGroup)
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go post(wg, logger, tag, message, requestsPerRoutine)
	}
	wg.Wait()
}

func post(wg *sync.WaitGroup, logger *fluent.Fluent, tag string, message string, requestsPerRoutine int) {

	for i := 0; i < requestsPerRoutine; i++ {
		now := time.Now()
		var data = map[string]string{
			"log": now.Format("2006-01-02T15:04:05.999999999") + message,
		}
		err := logger.PostWithTime(tag, time.Now(), data)
		if err != nil {
			panic(err)
		}
	}
	wg.Done()
}
