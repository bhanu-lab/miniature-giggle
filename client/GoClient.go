package client

import (
	"bufio"
	"fmt"
	"net/http"
	"strconv"
)

//GetToServer sends http request to server
func GetToServer() {
	PORT := 3000
	URL := "http://localhost"
	fmt.Println("Sending request to server on URL:",URL," on PORT: ",PORT)
	resp, err := http.Get(URL+":"+strconv.Itoa(PORT))
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
