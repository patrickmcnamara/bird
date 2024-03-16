// Package bird provides Bird protocol client and server implementations.
package bird

import (
	"io"
	"net"
	"strconv"
	"time"

	"github.com/patrickmcnamara/bird/bl"
	"github.com/patrickmcnamara/bird/seed"
)

// Fetch fetches the requested Bird page using Bird.
//
// rawBL is a Bird link. sr is the Bird page as a Seed document reader, close is
// a function that closes the underlying connection and err is any error in the
// connection to the Bird server.
func Fetch(rawBL string) (sr *seed.Reader, close func() (err error), err error) {
	// parse and validate BL
	b, err := bl.Parse(rawBL)
	if err != nil {
		return
	}
	// find Bird server
	target := b.Host
	port := DefaultPort
	_, srvs, err := net.LookupSRV("bird", "tcp", b.Host)
	if err == nil {
		srv := srvs[0]
		target = srv.Target
		port = srv.Port
	}
	// connect to Bird server
	conn, err := net.DialTimeout("tcp", target+":"+strconv.Itoa(int(port)), 5*time.Second)
	close = func() error { return conn.Close() }
	if err != nil {
		close()
		return
	}
	// send request to server
	if _, err = conn.Write([]byte(b.String() + "\n")); err != nil {
		close()
		return
	}
	// seed.Reader
	sr = seed.NewReader(conn)
	return
}

// Serve starts a Bird server using h to handle requests. It creates a new
// goroutine to handle each request. It automatically closes the connection
// after h returns. Any invalid requests are immediately skipped.
func Serve(address string, h Handler) (err error) {
	// start listening for connections
	lst, err := net.Listen("tcp", address)
	if err != nil {
		return
	}
	// accept connections in loop
	for {
		// accept connection
		conn, err := lst.Accept()
		if err != nil {
			continue
		}
		// handle request in new goroutine
		go serve(conn, h)
	}
}

// serve parses, validates and handles a Bird request. It closes the connection
// after it is done.
func serve(rwc io.ReadWriteCloser, h Handler) {
	// close connection when done
	defer rwc.Close()
	// parse BL from request
	request := make([]byte, bl.MaxSize+1) // +1 for newline character
	n, err := rwc.Read(request)
	if err != nil || request[n-1] != '\n' {
		return
	}
	rawBL := string(request[:n-1])
	b, err := bl.Parse(rawBL)
	if err != nil {
		return
	}
	// handle request
	h.ServeBird(b, seed.NewWriter(rwc))
}
