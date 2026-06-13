package main

import (
	"fmt"
	"sync"
)

func Withdraw(user int, amount int, balance *int, wg *sync.WaitGroup, mu *sync.Mutex) {

	defer wg.Done()
	mu.Lock()
	defer mu.Unlock()

	*balance -= amount

	fmt.Printf("User %d withdraws %d\n", user, amount)
}

func Deposit(user int, amount int, balance *int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	mu.Lock()
	defer mu.Unlock()
	*balance += amount

	fmt.Printf("User %d Deposit %d\n", user, amount)
}

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup

	balance := 1000

	wg.Add(1)
	go Deposit(1, 200, &balance, &wg, &mu)

	wg.Add(1)
	go Withdraw(2, 100, &balance, &wg, &mu)

	wg.Add(1)
	go Deposit(3, 50, &balance, &wg, &mu)

	wg.Add(1)
	go Withdraw(4, 300, &balance, &wg, &mu)

	wg.Wait()
	fmt.Println("Bal:", balance)

}
