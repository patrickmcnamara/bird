// Package bird provides Bird protocol client and server implementations.
package bird

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"net/url"
	"strconv"
	"strings"

	"github.com/patrickmcnamara/bird/seed"
)

// Fetch fetches the requested Seed document using Bird.
func Fetch(rawurl string) (sr *seed.Reader, err error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		return
	}
	if u.Scheme != "bird" {
		err = errors.New("incorrect URL scheme")
		return
	}
	if u.Port() == "" {
		u.Host += ":" + strconv.Itoa(int(DefaultPort))
	}
	conn, err := net.Dial("tcp", u.Host)
	if err != nil {
		return
	}
	if _, err = fmt.Fprintln(conn, rawurl); err != nil {
		return
	}
	sr = seed.NewReader(conn)
	return
}

// Serve starts a Bird server using h to handle requests.
func Serve(address string, h Handler) (err error) {
	lst, err := net.Listen("tcp", address)
	if err != nil {
		return
	}
	for {
		conn, err := lst.Accept()
		if err != nil {
			continue
		}
		br := bufio.NewReader(conn)
		rawurl, err := br.ReadString('\n')
		rawurl = strings.TrimSuffix(rawurl, "\n")
		if err != nil && err != io.EOF {
			continue
		}
		_, err = url.Parse(rawurl)
		if err != nil {
			continue
		}
		go func() {
			h.ServeBird(rawurl, seed.NewWriter(conn))
			if err := conn.Close(); err != nil {
			}
		}()
	}
}
