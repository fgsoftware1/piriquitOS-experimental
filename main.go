package main

import (
	"github.com/fgsoftware1/piriquitOS-experimental/libgo"
	"github.com/fgsoftware1/piriquitOS-experimental/utils"
)

func main() {
	stdio.Printf(color.Blue, "", "Hello, in blue!\n")
	stdio.Printf("", color.GreenBG, "Hello, in green background!\n")
	stdio.Printf(color.Magenta, color.YellowBG, "Foreground and background test!\n")
	format := "Hello, %s! You are %d years old."
	args := []interface{}{"Alice", 30}

	stdio.Vprintf(color.Green, color.BlackBG, format, args)

	// For a more dynamic approach:
	printColored := func(fg, bg, format string, args ...interface{}) {
		stdio.Vprintf(fg, bg, format, args)
	}

	printColored(color.Blue, color.WhiteBG, "The answer is %d", 42)
}
