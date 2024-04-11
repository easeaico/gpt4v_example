package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func main() {
	robotgo.ActiveName("Arc")
	title := robotgo.GetTitle()
	fmt.Println("title@@@ ", title)
}
