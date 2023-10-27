package stdio

import (
	"fmt"
	"os"
)

func Printf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, format, args...)
}
