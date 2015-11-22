package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "Hi from %s!", hostname)
}

func main() {
	addrs, _ := net.InterfaceAddrs()

	http.HandleFunc("/", handler)
	fmt.Printf("Server listening on  port 8080 and one of the ip from: %s\n", addrs)
	http.ListenAndServe(":8080", nil)
}
