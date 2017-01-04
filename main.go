package main

import (
	"fmt"
	"gopkg.in/redis.v5"
	"log"
	"net/http"
	"os"
	"strconv"
)

func counterHandler(w http.ResponseWriter, r *http.Request) {
	counter := r.URL.Path[1:]
	if r.Method == "GET" {
		count, err := client.Get(counter).Int64()
		if err != nil {
			count = 0
		}
		fmt.Fprintf(w, strconv.FormatInt(count, 10))
	} else if r.Method == "POST" {
		incCounter(counter)
		fmt.Fprintf(w, "%s increased", counter)
	} else {
		http.Error(w, "Invalid request method.", 405)
	}
}

func incCounter(counter string) {
	if err := client.Incr(counter).Err(); err != nil {
		panic(err)
	}

}

var (
	client *redis.Client
)

func main() {
	fmt.Println("Counters - Redis")
	redisAddr := os.Getenv("REDIS_URL")
	client = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	http.HandleFunc("/", counterHandler)     // set router
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
