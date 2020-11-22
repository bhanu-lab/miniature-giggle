package client

import (
	"bufio"
	"fmt"
	"net/http"
)

//GetToServer sends http request to server
func GetToServer() {

	resp, err := http.Get("http://localhost")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}