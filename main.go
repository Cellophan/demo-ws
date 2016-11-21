package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request, counterCh chan int) {
	hostname, _ := os.Hostname()
	requestId := <-counterCh
	fmt.Fprintf(w, "Hi from %s!", hostname)
	log.Printf("Host: %s RemoteAddr: %s RequestId: %d", hostname, r.RemoteAddr, requestId)
}

func main() {
	counterCh := make(chan int)
	addrs, _ := net.InterfaceAddrs()

	go func() {
		for x := 0; ; x++ {
			counterCh <- x
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, counterCh)
	})
	fmt.Printf("Server listening on port 8080 and all the interfaces: %s\n", addrs)
	http.ListenAndServe(":8080", nil)
}
