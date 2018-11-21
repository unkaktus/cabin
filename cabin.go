// cabin.go - ca certs bundle as a package.
//
// To the extent possible under law, Ivan Markin waived all copyright
// and related or neighboring rights to this module of cabin, using the creative
// commons "CC0" public domain dedication. See LICENSE or
// <http://creativecommons.org/publicdomain/zero/1.0/> for full details.

package cabin

import (
	"crypto/x509"

	"github.com/nogoegst/cabin/internal/pool"
)

func NewCertPool() *x509.CertPool {
	return pool.NewCertPool()
}
