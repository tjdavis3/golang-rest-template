package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	log "github.com/magna5/go-logger"

	"../../api"
)

var Version = "dev"
var Service = "RESTAPI"

func main() {
	// log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Info("Starting API version", Version)
	// configurtion with viper
	cfg := api.Configure(os.Args)

	if issuer := cfg.JWTIssuer; issuer != "" {
		// fetch JWT key set
		err := api.FetchJWTKeySet(cfg)
		if err != nil {
			log.Fatal("failed to fetch JWT Key Set ", err)
		}
		api.RefreshJWTKS(cfg)
	}
	// create server instance
	s, err := api.NewServer(context.Background(), cfg)
	if err != nil {
		log.Fatal(err)
	}

	api.InfoMetric.WithLabelValues(Service, Version).Inc()
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: s,
	}

	// graceful shutdown
	idleConnsClosed := make(chan struct{})
	go func() {
		shutdown := make(chan os.Signal, 1)
		signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
		for i := 0; true; i++ {
			<-shutdown

			if i > 0 {
				// force quit if received twice
				log.Fatal("force quit")
			}

			log.Println("shutting down")
			err := server.Shutdown(context.Background())
			if err != nil {
				log.Fatal("failed to shutdown gracefully")
			}

			close(idleConnsClosed)
		}
	}()

	log.Printf("starting %s", server.Addr)
	log.Print(server.ListenAndServe())
	<-idleConnsClosed
}
