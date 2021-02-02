# Bird

Bird is an alternative web protocol.
It is used to share hyperlinked Seed documents.
Seed is a simple markup language.
The birdspace is entirely plaintext, no binary at all.

This repo specifically is a Go library for using Bird protocol and the Seed markup language.
It also provides `eagle` which is a mix between `curl` and `cat` for documents in the birdspace.

## Quickstart

### Bird

Bird is a networking protocol.
It is incredibly simple and has no error handling whatsoever.
It is transported on TCP.

A request is simply sending a full URL followed by a newline character to a bird server.
A response is simply a valid Seed document in reply to the request.

### Seed document

Seed is a simple markup language for documents, similar to Markdown.
It is completely plaintext and it is line based.
It does not support embedding, only linking.
Every line has to end in just a newline character, not CRLF.

- Text line is just text (`$TEXT`).
- Header line one to six hashes (indicating header level), followed by the header text (`#{1,6} $TEXT`).
- Link line is of the form `=> $TEXT ||| $URL`.
- Quote marker is `>>>`.
- Code marker is ` ``` `.
- Break is just the newline.

Markers, like the quote marker and code marker, indicate the start of a block of that type.
The block ends when the marker appears again.
Anything inside the block is parsed as a text line.

## Documentation

- [Bird](https://pkg.go.dev/github.com/patrickmcnamara/bird)
- [Seed](https://pkg.go.dev/github.com/patrickmcnamara/bird/seed)

## License

This project is licensed under the MIT license.
