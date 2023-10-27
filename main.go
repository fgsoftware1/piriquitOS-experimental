package main

import (
	"github.com/fgsoftware1/piriquitOS-experimental/libgo"
	"github.com/fgsoftware1/piriquitOS-experimental/utils"
)

func main() {
	stdio.Printf(color.Blue, "", "Hello, in blue!\n")
	stdio.Printf("", color.GreenBG, "Hello, in green background!\n")
	stdio.Printf(color.Magenta, color.YellowBG, "Foreground and background test!\n")
}
