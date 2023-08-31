package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Result defines the data restructure of the result
type Result struct {
	TimeTaken      time.Duration `json:"TimeTaken"`
	NumberOfCoffee int           `json:"NumberOfCoffee"`
}

func main() {
	portNo := ":8080"
	http.HandleFunc("/serve-coffee/", ServeCoffee)

	log.Printf("Coffee server starting on: %s", portNo)
	http.ListenAndServe(portNo, nil)
}

// ServeCoffee is setup of the coffee shop API
func ServeCoffee(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// pass to the URL the number of customers
	noCustomers, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/serve-coffee/"))

	// Q: is there another way to initialise the noCustomers
	if err != nil || noCustomers == 0 {
		noCustomers = 1
	}

	// handler to wait for multiple goroutines to finish https://gobyexample.com/waitgroups
	wg := sync.WaitGroup{}
	// counter to keep track of customers
	count := 0
	for i := 0; i < noCustomers; i++ {
		wg.Add(1)
		go MakeCoffee(&wg)
		count++
	}
	wg.Wait()

	timeTaken := time.Since(start)
	response := Result{
		TimeTaken:      timeTaken,
		NumberOfCoffee: count,
	}
	json.NewEncoder(w).Encode(response)
	log.Printf("Took %s to serve coffee, customer no: %v", timeTaken, count)
}

// MakeCoffee is responsible to the process of serving a coffee
func MakeCoffee(wg *sync.WaitGroup) {
	defer wg.Done()

	// add more goroutines for each action in particular
	newWg := sync.WaitGroup{}
	newWg.Add(3)

	go TakePayment(&newWg)
	go MakeEspresso(&newWg)
	go FoamMilk(&newWg)
	newWg.Wait()
}

// MakeEspresso is responsible to make the espresso
func MakeEspresso(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	log.Printf("Coffee: Espresso is done.")
}

// FoamMilk is responsible to foam the milk
func FoamMilk(wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(2 * time.Second)
	log.Printf("Milk: Milk is ready.")

}

// TakePayment is responsible to take the payment from the customer
func TakePayment(wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(2 * time.Second)
	log.Printf("Money: Payment has been taken")
}
