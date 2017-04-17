package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
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

	///
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	priASN1 := x509.MarshalPKCS1PrivateKey(privateKey)
	pubASN1, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		panic(err)
	}

	priKeyStr := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: priASN1,
	})
	pubKeyStr := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubASN1,
	})

	///
	tmpl := "package {{.pkgname}}\n\n" +
		"func GetPrivateKey() (string) {\n" +
		"	return `{{.priKeyStr}}`\n" +
		"}\n\n" +
		"func GetPublicKey() (string) {\n" +
		"	return `{{.pubKeyStr}}`\n" +
		"}\n"

	t := template.New("t")
	template.Must(t.Parse(tmpl))

	dest := new(bytes.Buffer)
	t.Execute(dest, map[string]string{
		"pkgname":   *pkgname,
		"priKeyStr": fmt.Sprintf("%s", priKeyStr),
		"pubKeyStr": fmt.Sprintf("%s", pubKeyStr),
	})

	ioutil.WriteFile(*out, dest.Bytes(), 0644)
	fmt.Println("Output generate file:", out)
	fmt.Println("\tpacakge name:", *pkgname)
}
