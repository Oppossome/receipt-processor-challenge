package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	netHTTP "net/http"
	"os"

	"receipt-processor-challenge/internal/delivery/http"
	"receipt-processor-challenge/internal/delivery/oapi"
	"receipt-processor-challenge/internal/domain/usecases"
)

func main() {
	port := flag.String("port", "8080", "Port for the server to listen on.")
	flag.Parse()

	// Provide our HTTP handler to our Chi router.
	apiHandler := http.HTTPRepo{UsecasesRepo: usecases.NewUsecases()}
	r, err := oapi.NewChiRouter(&apiHandler)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec:\n%s", err)
		os.Exit(1)
	}

	s := &netHTTP.Server{
		Handler: r,
		Addr:    net.JoinHostPort("0.0.0.0", *port),
	}

	fmt.Printf("Running Receipt Processor on http://%s\n", s.Addr)

	// Serve our newly established http server.
	log.Fatal(s.ListenAndServe())
}
