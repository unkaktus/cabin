// generate.go - generate pool.go from ca bundle.
//
// To the extent possible under law, Ivan Markin waived all copyright
// and related or neighboring rights to this module of cabin, using the creative
// commons "CC0" public domain dedication. See LICENSE or
// <http://creativecommons.org/publicdomain/zero/1.0/> for full details.

// +build ignore

package main

import (
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

func main() {
	certs, err := ioutil.ReadFile("/etc/ssl/certs/ca-certificates.crt")
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create("pool.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	err = cabinTemplate.Execute(f, string(certs))
	if err != nil {
		log.Fatal(err)
	}
}

var cabinTemplate = template.Must(template.New("cabin").Parse(`// pool.go - ca certs bundle as a package.
//
// To the extent possible under law, Ivan Markin waived all copyright
// and related or neighboring rights to this module of cabin, using the creative
// commons "CC0" public domain dedication. See LICENSE or
// <http://creativecommons.org/publicdomain/zero/1.0/> for full details.

//go:generate go run generate.go
package pool

import (
	"crypto/x509"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Magic() {
	err := InitSystemPool()
	if err != nil {
		panic(err)
	}
}

func InitSystemPool() error {
	tmpf, err := ioutil.TempFile("/", "cabin")
	if err != nil {
		return err
	}
	filename := filepath.Join("/", tmpf.Name())
	defer os.Remove(filename)
	err = ioutil.WriteFile(filename, cacerts, 0400)
	if err != nil {
		return err
	}
	err = os.Setenv("SSL_CERT_FILE", filename)
	if err != nil {
		return err
	}
	defer os.Unsetenv("SSL_CERT_FILE")
	_, err = x509.SystemCertPool()
	if err != nil {
		return err
	}
	return nil
}

func NewCertPool() *x509.CertPool {
	pool := x509.NewCertPool()
	ok := pool.AppendCertsFromPEM(cacerts)
	if !ok {
		panic("cabin: unable to read certificates")
	}
	return pool
}

var cacerts = []byte(` + "`{{ . }}`" + `)
`))
