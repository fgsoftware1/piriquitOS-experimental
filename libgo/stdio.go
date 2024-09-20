package stdio

import (
	"fmt"
	"log"
	"os"

	"github.com/fgsoftware1/piriquitOS-experimental/utils"
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
	coloredFormat := fmt.Sprintf("%s%s%s%s", bgColor, fgColor, format, color.Reset)
	n, err := fmt.Fprintf(os.Stdout, coloredFormat, args...)
	if err != nil {
		log.Println("Error printing to standard output:", err)
		return n, err
	}
	return n, nil
}