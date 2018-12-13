package g

import (
	"fmt"

	"github.com/jhonghee/f"
)

// Version returns the tagged version of the module
func Version() string {
	return fmt.Sprint("G v1.1.0", "->", f.Version())
}
