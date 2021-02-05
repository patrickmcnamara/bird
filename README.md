# Bird

Bird is an alternative web protocol.
It is used to share Seed documents which can contain links.
Seed is a simple markup language.
The birdspace is entirely plaintext, no binary at all (as God intended).

This repository specifically is a Go library for using Bird protocol and the Seed markup language.
It is the refence implementation of both.
It also provides `eagle` which is a mix between `curl` and `cat` for documents in the birdspace.
It also also provides the library and binary `owl` which allows you to easily host a Bird protocol file server.

## Quickstart

### Bird protocol

Bird is a request-response networking protocol for the birdspace.
It is incredibly simple and has no error handling whatsoever.
It is transported on TCP.

A request is simply sending a Bird protocol-relative URL without a trailing slash followed by a newline character to a bird server.
A response is simply a valid Seed document in reply to the request.

### Seed document

Seed is a simple markup language for documents, similar to Markdown.
It is completely plaintext and it is line based.
It does not support embedding, only linking.
Every line has to end in just a newline character, not CRLF.

## Other Documentation

- [Bird Specification](#bird-protocol) :)
- [Bird Library](https://pkg.go.dev/github.com/patrickmcnamara/bird)
- [Seed Specification](./seed/README.md)
- [Seed Library](https://pkg.go.dev/github.com/patrickmcnamara/bird/seed)

## License

This project is licensed under the MIT license.
