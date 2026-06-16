package main

import (
	"fmt"
	"sync"
	"unicode"
)

func hasNumber(s string) bool {
	for _, ch := range s {
		if unicode.IsDigit(ch) {
			return true
		}
	}
	return false
}

func Validator(password *string, wg *sync.WaitGroup) {
	defer wg.Done()
	var status string

	if len(*password) < 6 || !hasNumber(*password) {
		status = "invalid"
	} else {
		status = "Valid"
	}

	fmt.Println(*password, status)

}

func main() {
	var wg sync.WaitGroup

	passwords := []string{"abc", "hello123", "GOlang!", "hi", "secure99"}

	for _, password := range passwords {
		wg.Add(1)
		go Validator(&password, &wg)
	}
	wg.Wait()

}
