package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const PI string = "3.141592653589793238462643383279502884197169399375105"
const PORT string = ":8080"

type myRequest struct {
	*http.Request
}

func (r *myRequest) String() string {
	return r.Method + " " + r.URL.RequestURI()
}

func hello(w http.ResponseWriter, req *http.Request) {
	r := &myRequest{req}
	fmt.Fprintf(w, "%v", r)
	log.Printf("%v", r)
}

func pi(w http.ResponseWriter, req *http.Request) {
	index, ok := req.URL.Query()["index"]

	if !ok || len(index) < 1 {
		fmt.Fprintf(w, PI)
		return
	}

	i, err := strconv.Atoi(index[0])
	if err != nil {
		fmt.Fprintf(w, "Invalid index, must be an integer")
		return
	}

	if i < 0 {
		fmt.Fprintf(w, "Invalid index, must be a positive integer")
		return
	}

	fmt.Fprintf(w, "%c", PI[i])
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Printf("here")
	log.Printf(r.URL.RequestURI())
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/pi", pi)
	http.HandleFunc("/pi/index/*", index)

	log.Printf("Running on port %v", PORT)
	http.ListenAndServe(PORT, nil)
}
