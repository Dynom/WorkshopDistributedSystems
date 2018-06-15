package main

import (
	"fmt"
	"net/http"
	"sync"
)

var (
	counter uint64     // Our counter
	lock    sync.Mutex // Our mutex
)

func main() {
	// Register a path for our handler
	http.HandleFunc("/counter", func(w http.ResponseWriter, _ *http.Request) {

		// Handling concurrency correctly, by locking our counter.
		// This way only one thread can ever only be incrementing our counter
		lock.Lock()
		defer lock.Unlock()

		// Incrementing the counter
		counter++

		// Returning the response to the client
		fmt.Fprintf(w, "Counter is: %d", counter)
	})

	// Listen on port 8080
	fmt.Println(http.ListenAndServe(":8080", nil))
}
