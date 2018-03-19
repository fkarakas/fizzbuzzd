package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/fkarakas/fizzbuzzd/router"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	version = "undefined"
)

// Endpoint for the API server
func main() {
	port := flag.Int("port", 8080, "API server port number")
	addr := fmt.Sprintf(":%d", *port) // :8080

	flag.Parse()

	router := router.NewRouter(gin.ReleaseMode, version)

	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	go func() {
		log.Printf("fizzbuzzd listening on port %v ...", *port)

		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Printf("Shutdown Server ...")

	// wait max 10 seconds before killing
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Printf("Server shutdown")
}
