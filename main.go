package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	url := "https://camo.githubusercontent.com/7f22a4ed26a359284aedc749565c435f40cee30aa64bf5d7f7759d62c29cd6c4/68747470733a2f2f6b6f6d617265762e636f6d2f67687076632f3f757365726e616d653d7661636172616d696e266c6162656c3d50726f66696c65253230766965777326636f6c6f723d306537356236267374796c653d666c6174"
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go makeRequest(url, &wg, i)
	}

	wg.Wait()
	fmt.Println("All requests completed.")
}

func makeRequest(URL string, wg *sync.WaitGroup, i int) {
	defer wg.Done()

	_, err := http.Get(URL)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}
	fmt.Println(i, " complete")
}
