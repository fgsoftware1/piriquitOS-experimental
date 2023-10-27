package libgo

import (
	"fmt"
	"os"
)

func printf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, format, args...)
}
