package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var Tasks TaskModel

func init() {
	Tasks = NewTaskModel()
}

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	sm := mux.NewRouter()

	getRouter := sm.Methods("GET").Subrouter()
	getRouter.HandleFunc("/api/v1/items", getItems)

	putRouter := sm.Methods("PUT").Subrouter()
	putRouter.HandleFunc("/api/v1/{id:[0-9]+}", deleteItem)

	postRouter := sm.Methods("POST").Subrouter()
	postRouter.HandleFunc("/api/v1/", addItem)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Terminating Server ", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}

func getItems(rw http.ResponseWriter, rq *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(Tasks.ToJSON())
}

func addItem(rw http.ResponseWriter, rq *http.Request) {
	bytes, _ := ioutil.ReadAll(rq.Body)
	Tasks.AddTask(string(bytes))
}

func deleteItem(rw http.ResponseWriter, rq *http.Request) {
	vars := mux.Vars(rq)
	id, _ := strconv.Atoi(vars["id"])
	Tasks.DeleteTask(id)
}
