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
		log.Printf("######\nRequest received from Pod %s at path %s with cookies %v\n", podName, r.URL.Path, r.Cookies())
		podSelector, err := r.Cookie("pod-selector")
		if err != nil {
			podSelector = &http.Cookie{Name: "pod-selector", Value: "N/A"}
		}
		// fmt.Println("Printing all cookies")
		// for _, c := range r.Cookies() {
		// 	fmt.Println(c)
		// }
		log.Printf("Response from Pod %s to path %s, with 'pod-selector' cookie: %v\n", podName, r.URL.Path, podSelector.Value)
		fmt.Fprintf(w, "Response from Pod %s to path %s, with 'pod-selector' cookie: %v\n", podName, r.URL.Path, podSelector.Value)
	})

	log.Println("Listening on localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
