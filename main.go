package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/noritama/gen-keypair/keypair"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	///
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `
Usage of %s:
   %s [OPTIONS] ARGS...
Options`, os.Args[0], os.Args[0])

		flag.PrintDefaults()
	}

	pkgname := flag.String("pkgname", "main", "package name")
	out := flag.String("out", cwd+"/keypair_gen.go", "output file path")

	flag.Parse()


	keypair.Generate(pkgname, out)
}
