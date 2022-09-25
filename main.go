package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func FetchFile(url string) {

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	out, err := os.Create("output.csv")
	defer out.Close()

	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

}

func main() {

	url := os.Args[1]
	log.Println("Running Program...")
	log.Println(url)
	FetchFile(url)

}
