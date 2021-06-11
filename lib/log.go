package lib

import (
	"fmt"
	"os"
)

// TODO: improve this help mesage
func HelpAndExit(message... string)  {
	fmt.Println(message)
	os.Exit(1)
}