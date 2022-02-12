package main

import (
	"fmt"
	"os"
)

func main() {
	var argsWithProg string = os.Args[1]

	switch arg := argsWithProg; arg {
	case "checkin":
		fmt.Println("Good Morning Martin, you started working at FOO !")
		//TODO REST call to the server
	case "checkout":
		fmt.Println("Good Morning Martin, you stopped working at FOO !")
	case "help":
		fmt.Print(helpStringHandler())

	default:
		//TODO REST call to the server
		err := fmt.Errorf("Wrong arguments! Please call `trkctl help` to get help!")
		fmt.Println(err.Error())
	}
	//fmt.Printf(argsWithProg[0])
}

func checkInHandler() {

}

func checkOutHandler() {

}

// Generating help string
func helpStringHandler() string {
	multiline := `- Using trkctl - 
Trkctl is a CLI tool to communicate with the trkbox server
__

trkctl [command]
	
	checkin 	Check in
	checkout	Check out
	help		shows this help message`

	return multiline
}
