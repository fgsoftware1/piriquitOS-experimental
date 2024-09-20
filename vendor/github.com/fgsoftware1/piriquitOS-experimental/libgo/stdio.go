package stdio

import (
	"fmt"
	"log"
	"os"

	c "github.com/fgsoftware1/piriquitOS-experimental/utils"
)

// Printf formats and prints the specified string to standard output.
//
// The format parameter specifies the format of the output string. It supports
// verbs that are similar to the ones used in the fmt package.
//
// The args parameter is a variadic parameter that accepts a variable number of
// arguments. They will be used to replace the verbs in the format string.
//
// This function does not return anything.
func Printf(fgColor, bgColor, format string, args ...interface{}) {
    if _, err := fmt.Fprintf(os.Stdout, "%s%s%s%s\n", bgColor, fgColor, fmt.Sprintf(format, args...), c.Reset); err != nil {
        log.Println("Error printing to standard output:", err)
    }
}
