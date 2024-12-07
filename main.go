package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"
)

func main() {
	// Define command line flags
	targetPort := flag.Int("port", 0, "Target port of the local API (required)")
	proxyPort := flag.Int("proxy", 8443, "Port for the HTTPS proxy (default: 8443)")
	flag.Parse()

	// Validate required port argument
	if *targetPort == 0 {
		fmt.Println("Error: Target port is required")
		fmt.Println("Usage: go run main.go -port <target_port> [-proxy <proxy_port>]")
		flag.PrintDefaults()
		os.Exit(1)
	}

	targetURL := fmt.Sprintf("http://localhost:%d/", *targetPort)
	proxyURL := fmt.Sprintf("https://localhost:%d/", *proxyPort)
	fmt.Printf("Starting HTTPS proxy on port %d forwarding to %s\n", *proxyPort, targetURL)
	fmt.Printf("Visit %s\n", proxyURL)
	
	ReverseHttpsProxy(*proxyPort, targetURL, "my.crt", "my.key")
}

func ReverseHttpsProxy(port int, dst string, crt string, key string) {
	u, e := url.Parse(dst)
	if e != nil {
		log.Fatal("Bad destination.")
	}
	h := httputil.NewSingleHostReverseProxy(u)
	//if your certificate signed by yourself,you need use this bypass secure verify
	var InsecureTransport http.RoundTripper = &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		TLSHandshakeTimeout: 10 * time.Second,
	}
	h.Transport = InsecureTransport
	err := http.ListenAndServeTLS(fmt.Sprintf(":%d", port), crt, key, h)
	if err != nil {
		log.Println("Error:", err)
	}
}

func ReverseHttpProxy(port int, dst string) {
	u, e := url.Parse(dst)
	if e != nil {
		log.Fatal("Bad http destination.")
	}
	h := httputil.NewSingleHostReverseProxy(u)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), h)
	if err != nil {
		log.Println("Error:", err)
	}
}
