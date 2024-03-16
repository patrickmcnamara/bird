# Bird

Bird is an alternative web protocol.
It is used to share Bird pages.
Bird pages are Seed documents which can contain links.
Seed is a simple markup language.
Seed, and thus the birdspace, is entirely plaintext. No binary at all - as God intended.

This repository is a Go library for using Bird protocol and the Seed markup language.
It is the reference implementation of both.

It also provides `bl`, a library for handling and parsing Bird links.
It also also provides `eagle`, a binary for `curl`ing and `cat`ing documents in the birdspace.
It also also also provides `owl`, a library and binary which allows you to easily host a Bird protocol file server.

## Quickstart

### Bird protocol

Bird is a request-response networking protocol for the birdspace, similar to HTTP.
It is incredibly simple and has no error handling whatsoever.
It is transported via TCP.

A request is simply sending a Bird link followed by a newline character to a Bird server.
A response is simply a Bird page (i.e. a valid Seed document) in reply to the request.

### Bird links

Bird links are the hyperlinks used in the birdspace.
They have the format `//host/[path]`.

### Seed document

Seed is a simple markup language for documents, similar to Markdown.
It is completely plaintext and line-based.
It does not support embedding any content, only linking.
Every line has to end in just a newline character, not CRLF.

## Other Documentation

- [Bird Specification](#bird-protocol) :)
- [Bird Library](https://pkg.go.dev/github.com/patrickmcnamara/bird)
- [Seed Specification](./seed/README.md)
- [Seed Library](https://pkg.go.dev/github.com/patrickmcnamara/bird/seed)

## License

This project is licensed under the MIT license.
