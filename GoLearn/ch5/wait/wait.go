package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	url := "https://www.google.com"
	//if err := WaitForServer(url); err != nil {
	//	fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
	//	os.Exit(1)
	//}
	log.SetPrefix("wait: ")
	log.SetFlags(0)
	if err := WaitForServer(url); err != nil {
		log.Fatalf("Site is down: %v\n", err)
	}
	fmt.Println("Site is up")
}

// WaitForServer attempts to contact the server of a URL.
// It tries for one minute using exponential back-off.
// It reports an error if all attempts fail.
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // success
		}
		log.Printf("server not responding (%s);retrying…", err)
		time.Sleep(time.Second << uint(tries)) // exponential back-off
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
