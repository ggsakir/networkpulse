package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	fmt.Println("NetworkPulse started")
	servers := []string{
		"https://google.com",
		"https://github.com",
		"https://example.com",
		"https://example.invalid",
		"https://example.com",
		"https://httpbin.org/status/404",
	}
	for {
		var wg sync.WaitGroup
		wg.Add(len(servers))
		for _, server := range servers {
			go checkAndPrint(server, &wg)
		}
		wg.Wait()
		time.Sleep(30 * time.Second)
	}
}

func checkServer(server string) (int, error) {
	client := http.Client{Timeout: 5 * time.Second}
	response, err := client.Get(server)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()
	return response.StatusCode, nil
}

func checkAndPrint(server string, wg *sync.WaitGroup) {
	status, err := checkServer(server)
	if err != nil {
		fmt.Println(server, "-", status, err)
	} else if status >= 200 && status < 300 {
		fmt.Println(server, "-", "Успешный HTTP-статус", status)
	} else {
		fmt.Println(server, "-", "Неуспешный HTTP-статус", status)
	}
	wg.Done()
}
