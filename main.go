package main // import "dummy-server"

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	portArg := flag.String("port", "4416", "Port number - 4416, 5525, 6636")
	flag.Parse()

	port := *portArg

	role := map[string]string{
		"4416": "school",
		"5525": "class",
		"6636": "teacher",
		"7749": "student",
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, req *http.Request) {
		log.Println("Received health check request")
		w.Write([]byte("It works!\n"))
	})
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		log.Println("Received hello request")
		w.Write([]byte("Hello this is " + port + " for " + role[port] + "\n"))
	})

	http.ListenAndServe("127.0.0.1:"+port, nil)
}
