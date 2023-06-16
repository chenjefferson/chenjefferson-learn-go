# Front Matter and Setup
This covers the front matter (Learn Go with Tests), Install Go, and "Hello, World" sections.
## Why

The blue book w / exercises requires high commitment. Katas are scoped.

What worked was [Go by Example](https://gobyexample.com/).

## Additional Material

- Go by Example
- LeetCode
- blue book exercises

## Environment

For builds, `GOPATH` is not recommended and has been replaced by `Modules`.

Modules attempt to solve dependency management, versions, reproducible builds, and running outside of `GOPATH`.

`go mod init` initializes a module, creating a `go.mod` file with the module path, version, and its dependencies. It attempts to guess the modulepath via the directory structure, but the modulepath can be provided as well.

Refactoring is super important. You should be able to:

- extract/inline variable
- extract method/function
- rename
- go fmt
- run tests

And in-editor, you should be able to see:

- fn signatures
- fn definitions
- usages of a symbol

These will help you reduce context switching.

## Hello, World

See the Anki cards in the CSV. I'm missing some as I had some Anki cards already.

## Integers

See above. you can document your code at `pkg.go.dev` if you share your code!