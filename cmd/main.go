package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
)

type Config struct {
	KeyFile  string
	CertFile string
	Port     int
}

func main() {
	cfg := &Config{}
	flag.StringVar(&cfg.KeyFile, "key", "", "tls server key file path.")
	flag.StringVar(&cfg.CertFile, "cert", "", "tls server key file path.")
	flag.IntVar(&cfg.Port, "port", 443, "server port.")
	flag.Parse()

	run(cfg)

}

func run(cfg *Config) {

	pair, err := tls.LoadX509KeyPair(cfg.CertFile, cfg.KeyFile)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	server := &http.Server{
		Addr:      fmt.Sprintf(":%d", cfg.Port),
		TLSConfig: &tls.Config{Certificates: []tls.Certificate{pair}},
	}

	server.Handler = mux

	err = server.ListenAndServeTLS("", "")
	if err != nil {
		log.Fatal(err)
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi")
}
