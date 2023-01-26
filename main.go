package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello Go!")
	})

	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		d, err := ioutil.ReadAll(r.Body)
		// log.Printf("Data is %s", d) For logging Request Body In Terminal
		fmt.Fprintf(w, "%s", d)
		if err != nil {
			http.Error(w, "Something Wrong happend", http.StatusBadRequest)
		}
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye")
	})
	http.ListenAndServe(":9090", nil)
}
