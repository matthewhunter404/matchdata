package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type FootballData struct {
	div      string //League Division
	date     string //Match Date (dd/mm/yy)
	time     string //Time of match kick off
	homeTeam string //Home Team
	awayTeam string //Away Team
	fthg     string //Full Time Home Team Goals
	ftag     string //Full Time Away Team Goals
}

func FetchFile(url string) {

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	csvReader := csv.NewReader(resp.Body)
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	var allData []FootballData
	for _, record := range records {

		footballData := FootballData{
			div:      record[0],
			date:     record[1],
			time:     record[2],
			homeTeam: record[3],
			awayTeam: record[4],
			fthg:     record[5],
			ftag:     record[6],
		}
		allData = append(allData, footballData)
	}

	for _, match := range allData {
		fmt.Printf("%+v\n", match)
	}
}

func main() {

	url := os.Args[1]
	log.Println("Running Program...")
	log.Println(url)
	FetchFile(url)

}
