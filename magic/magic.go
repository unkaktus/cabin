// magic.go - run package init (for _-imports).
//
// To the extent possible under law, Ivan Markin waived all copyright
// and related or neighboring rights to this module of cabin, using the creative
// commons "CC0" public domain dedication. See LICENSE or
// <http://creativecommons.org/publicdomain/zero/1.0/> for full details.

package magic

import (
	"github.com/unkaktus/cabin/internal/pool"
)

func init() {
	pool.Magic()
}
