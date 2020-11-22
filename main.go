package main

import (
	"fmt"
	"miniature-giggle/client"
	"miniature-giggle/server"
	"os"
	"strings"
)

func main() {
	argsWithoutProg := os.Args[1:]

	//Run either server or client based on arguments received
	if strings.EqualFold(argsWithoutProg[0], "client") {
		client.GetToServer(argsWithoutProg[1], argsWithoutProg[2])
	} else if strings.EqualFold(argsWithoutProg[0], "server") {
		server.Serve(argsWithoutProg[1])
	} else {
		fmt.Println("Choose either server or client while running application like ==> go run main.go client / go run main.go server")
	}
}
