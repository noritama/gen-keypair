gen-keypair
==========

Automatically generates RSA key-pair.

# Download

- Release page: [https://github.com/noritama/gen-keypair/releases](https://github.com/noritama/gen-keypair/releases)
- Download binary file(version 0.0.1): [https://github.com/noritama/gen-keypair/releases/download/v0.0.1/gen-keypair](https://github.com/noritama/gen-keypair/releases/download/v0.0.1/gen-keypair)

# install

```sh
$ go get github.com/noritama/gen-keypair
```

# Use

```sh
$ gen-keypair -out /tmp/hoge_gen.go -pkgname main
Output generate file: /tmp/hoge_gen.go
	pacakge name: main

$ cat /tmp/hoge_gen.go
package main

func GetPrivateKey() (string) {
	return `-----BEGIN RSA PRIVATE KEY-----
...
-----END RSA PRIVATE KEY-----
`
}

func GetPublicKey() (string) {
	return `-----BEGIN RSA PUBLIC KEY-----
...
-----END RSA PUBLIC KEY-----
`
}
```

# Options

```
$ gen-keypair -h

Usage of gen-keypair:
   gen-keypair [OPTIONS] ARGS...

Options  -out="/Users/noritama/repository/github/gen-keypair/keypair_gen.go": output file path
  -pkgname="main": package name
  
```

# go generate

```go
//go:generate gen-keypair -pkgname gen -out ./gen/keypair_gen.go
```

> go1.4

# build

```sh
$ make # => ./gen-keypair
```
