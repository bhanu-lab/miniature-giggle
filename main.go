package main

import (
	"SampleServerClient/client"
	"SampleServerClient/server"
	"fmt"
	"os"
	"strings"
)

func main() {
	argsWithoutProg := os.Args[1:]

	//Run either server or client based on arguments received
	if strings.EqualFold(argsWithoutProg[0], "client"){
		client.GetToServer()
	} else if strings.EqualFold(argsWithoutProg[0], "server") {
		server.Serve()
	} else {
		fmt.Println("Choose either server or client while running application like ==> go run main.go client / go run main.go server")
	}
}
