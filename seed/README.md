# Seed

Seed is the markup language for Bird protocol, inspired heavily by Markdown.
It can also be used outside of that context as it is very simple.
It is line-based and is easy to write/read in both plaintext and in a rendered format.
Embedding is not supported, only linking.

## Line types

### Text

"`$TEXT`"

Normal text.
A group of one or more text lines forms a paragraph unless they are within a code block.
Paragraphs are separated using blank text lines.
Any line that doesn't match another line type is a text line.

### Header

"`#{1,6} $TEXT`"

A header.
The header level can be one up to six with one being the highest level.

### Link

"`=> $TEXT ||| $URL`"

A link.
This links content from another URL.
The text is used to give context to the link.
Sometimes the URL won't be rendered at all.

### Quote marker

"`>>>`"

A quote marker.
This marks both the start and end of a quote block.
Only text lines are valid in a quote block.

```
>>>
Ninety percent of everything is crap.
~ Sturgeon's Law
>>>
```

### Code marker

"` ``` `"

A code marker.
This marks both the start and end of a code block.
Only text lines are valid in a code block.
Text lines do not form into paragraphs in a code block.

````
```
package main

import "fmt"

func main() {
    fmt.Println("Hello, world!")
}
```
````

## Other

- The MIME type is `text/seed`.
- UTF-8 must be used.
- Each line must end in a newline `\n`, not CRLF.
- Styling of the document is completely up to the client.

## Best practices

- Write one sentence per text line.
- Always close your quote and code blocks.
- Do not try to control formatting using ASCII or anything.
- Use protocol-relative URLs to link to other Seed documents in the birdspace.

## Example

```
# GNU/Linux

I'd just like to interject for a moment.
What you're referring to as Linux, is in fact, GNU/Linux, or as I've recently taken to calling it, GNU plus Linux.
Linux is not an operating system unto itself, but rather another free component of a fully functioning GNU system made useful by the GNU corelibs, shell utilities and vital system components comprising a full OS as defined by POSIX.

>>>
No, Richard, it's 'Linux', not 'GNU/Linux'.
~ Linus Torvalds.
>>>

=> GNU/Linux copypasta ||| https://wiki.installgentoo.com/wiki/Interjection
```
