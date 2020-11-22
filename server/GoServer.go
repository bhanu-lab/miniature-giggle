package server

import (
	"fmt"
	"net/http"
)

//Serve opens and awaits for connections from client
func Serve(){
	fmt.Println("Server listening on port 3000")
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":3000", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, " ==> %s!", r.URL.Path[1:])
	fmt.Fprintf(w, " ==> %s!", r.Header)
	fmt.Println(r.Header)
}
