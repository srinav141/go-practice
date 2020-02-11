package main

import (
	"./lissagous/liss"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/h2", handler2)
	http.HandleFunc("/liss", func(w http.ResponseWriter, r *http.Request) {
		liss.Liss(w)
	})
	http.HandleFunc("/lissw", handlerliss)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path=%q\n", r.URL.Path)

}

func handlerliss(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["cycles"]
	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'cycles' is missing")
		return

	}
	// Query()["key"] will return an array of items, // we only want the single item.
	key := keys[0]
	log.Println("Url Param 'cycles' is: " + string(key))
	v, err := strconv.ParseFloat(key, 64)
	if err == nil {
		liss.LissWithCycles(w, v)
	}
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
