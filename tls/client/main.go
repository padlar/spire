package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
)

const (
	serverAddress  = "server.server.svc.cluster.local:33333"
	serverSpiffeID = "spiffe://cluster.example.com/ns/server/sa/default"
	serverID       = "spiffe://cluster.example.com/ns/server/sa/default"
)

func main() {
	ticker := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-ticker.C:
			run()
		}
	}
}

func run() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Allowed SPIFFE ID
	serverID := spiffeid.RequireFromString(serverID)

	// workload socket read from the env SPIFFE_ENDPOINT_SOCKET
	conn, err := spiffetls.Dial(ctx, "tcp", serverAddress, tlsconfig.AuthorizeID(serverID))
	if err != nil {
		log.Fatalf("Unable to create TLS connection: %v", err)
	}
	defer conn.Close()

	// Send a message to the server using the TLS connection
	fmt.Fprintf(conn, "Hello server\n")

	// Read server response
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil && err != io.EOF {
		log.Fatalf("Unable to read server response: %v", err)
	}
	log.Printf("Server says: %q", status)
}
