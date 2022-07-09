package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/stretchr/testify"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()
	router.HandleFunc("/name/{PARAM}", getParam).Methods(http.MethodGet)
	router.HandleFunc("/bad", getBad).Methods(http.MethodGet)
	router.HandleFunc("/data", postData).Methods(http.MethodPost)
	router.HandleFunc("/headers", postHeaders).Methods(http.MethodPost)
	http.Handle("/", router)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}

func getParam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "Hello, %v!", vars["PARAM"])
}

func getBad(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
}

func postData(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("content-type", "text/plain")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "I got message:\n%v", string(b))
}

func postHeaders(w http.ResponseWriter, r *http.Request) {
	var ans int
	var headers = []string{"a", "b"}

	for _, header := range headers {
		stringValue := r.Header.Get(header)
		intValue, _ := strconv.Atoi(stringValue)
		ans += intValue
	}

	w.Header().Set("a+b", fmt.Sprint(ans))
	w.WriteHeader(http.StatusOK)
}
