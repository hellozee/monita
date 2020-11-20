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
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var Tasks TaskModel

var (
	requestsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "monita",
			Subsystem: "service",
			Name:      "processed_total",
			Help:      "Total number of requests processed by the service",
		},
		[]string{"type"},
	)
)

func init() {
	Tasks = NewTaskModel()
	prometheus.MustRegister(requestsCounter)
}

func main() {
	sm := mux.NewRouter()

	getRouter := sm.Methods("GET").Subrouter()
	getRouter.HandleFunc("/api/v1/items", getItems)
	getRouter.Handle("/metrics", promhttp.Handler())

	putRouter := sm.Methods("PUT").Subrouter()
	putRouter.HandleFunc("/api/v1/{id:[0-9]+}", deleteItem)

	postRouter := sm.Methods("POST").Subrouter()
	postRouter.HandleFunc("/api/v1/", addItem)

	s := &http.Server{
		Addr:         ":8080",
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

	<-sigChan
	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
	cancel()
}

func getItems(rw http.ResponseWriter, rq *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(Tasks.ToJSON())
	requestsCounter.WithLabelValues("get_items").Inc()
}

func addItem(rw http.ResponseWriter, rq *http.Request) {
	bytes, _ := ioutil.ReadAll(rq.Body)
	Tasks.AddTask(string(bytes))
	rw.WriteHeader(http.StatusOK)
	requestsCounter.WithLabelValues("add_item").Inc()
}

func deleteItem(rw http.ResponseWriter, rq *http.Request) {
	vars := mux.Vars(rq)
	id, _ := strconv.Atoi(vars["id"])
	Tasks.DeleteTask(id)
	rw.WriteHeader(http.StatusOK)
	requestsCounter.WithLabelValues("delete_item").Inc()
}
