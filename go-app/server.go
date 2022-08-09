package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		podName := "NA"
		podName, isSet := os.LookupEnv("POD_NAME")
		if !isSet {
			podName = "NA"
		}
		podSelector, err := r.Cookie("pod-selector")
		if err != nil {
			podSelector = &http.Cookie{Name: "pod-selector", Value: "N/A"}
			log.Println("No 'pod-selector' cookie received")
		} else {
			log.Printf("Received 'pod-selector' cookie %s\n", podSelector.Value)
		}
		fmt.Println("Printing all cookies")
		for _, c := range r.Cookies() {
			fmt.Println(c)
		}
		fmt.Fprintf(w, "Response from Pod %s to path %s, with 'pod-selector' cookie: %v", podName, r.URL.Path, podSelector.Value)
	})

	log.Println("Listening on localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
