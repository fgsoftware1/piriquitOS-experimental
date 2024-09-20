package stdio

import (
	"fmt"
	"log"
	"os"
)

func Printf(format string, args ...interface{}) (int, error) {
	n, err := fmt.Fprintf(os.Stdout, format, args...)
	if err != nil {
		log.Println("Error printing to standard output:", err)
		return n, err
	}
	return n, nil
}

func Vprintf(fgColor, bgColor, format string, args []interface{}) (int, error) {
	coloredFormat := fmt.Sprintf("%s%s%s%s", bgColor, fgColor, format, c.Reset)
	n, err := fmt.Fprintf(os.Stdout, coloredFormat, args...)
	if err != nil {
		log.Println("Error printing to standard output:", err)
		return n, err
	}
	return n, nil
}