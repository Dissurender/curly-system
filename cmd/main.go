package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"sync"
)

type Backend struct {
	URL         *url.URL
	Alive       bool
	mux         sync.RWMutex
	ReverseProx *httputil.ReverseProxy
}

type ServerPool struct {
	backends []*Backend
	current  uint64
}
u, _ := url.Parse("http://localhost:8080")
rp := httputil.NewSingleHostReverseProxy(u)

func main() {

	http.HandlerFunc(rp.ServeHTTP)
	
	fmt.Println("Test")

	os.Exit(0)
}
