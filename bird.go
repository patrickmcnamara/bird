// Package bird provides Bird protocol client and server implementations.
package bird

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/patrickmcnamara/bird/seed"
)

// Fetch fetches the requested Seed document using Bird.
//
// rawurl is a URL that is protocol-relative URL without a trailing slashes.
// bird://hello/world/ is simply //hello/world.
//
// sr is the Seed document reader, close is a function that closes the
// underlying connection and err is any error in creating a connection to the
// Bird server.
func Fetch(rawurl string) (sr *seed.Reader, close func() (err error), err error) {
	// parse and validate url
	u, err := parseURL(rawurl)
	if err != nil {
		return
	}
	// add default port if none defined
	if u.Port() == "" {
		u.Host += ":" + strconv.Itoa(int(DefaultPort))
	}
	// connect to Bird server
	conn, err := net.DialTimeout("tcp", u.Host, 5*time.Second)
	if err != nil {
		return
	}
	// send request to server
	if _, err = fmt.Fprintln(conn, u.String()); err != nil {
		return
	}
	// create seed.Reader and function to close connection
	sr = seed.NewReader(conn)
	close = func() error { return conn.Close() }
	return
}

// Serve starts a Bird server using h to handle requests. It opens a new
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

// serve parses, validates and handles a Bird request and closes the connection
// after it is done.
func serve(rwc io.ReadWriteCloser, h Handler) {
	// close connection when done
	defer rwc.Close()
	// parse URL from request and validate
	br := bufio.NewReaderSize(rwc, 256)
	rawurl, err := br.ReadString('\n')
	if err != nil && err != io.EOF {
		return
	}
	rawurl = strings.TrimSuffix(rawurl, "\n")
	u, err := parseURL(rawurl)
	if err != nil {
		return
	}
	// handle request
	h.ServeBird(u, seed.NewWriter(rwc))
}
