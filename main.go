package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"time"
)

type FootballMatch struct {
	div      string //League Division
	date     string //Match Date (dd/mm/yy)
	time     string //Time of match kick off
	homeTeam string //Home Team
	awayTeam string //Away Team
	fthg     string //Full Time Home Team Goals
	ftag     string //Full Time Away Team Goals
}

func PrintFootballMatches(footballMatches []FootballMatch) {
	for _, match := range footballMatches {
		fmt.Printf("%+v\n", match)
	}
}

// Slices are passed by reference so this sorts the passed in slice
func SortAccendingDate(footballMatches []FootballMatch) error {
	var sortingError error
	sort.Slice(footballMatches, func(i, j int) bool {
		iDateTime, err := time.Parse("02/01/2006 15:04", fmt.Sprintf("%s %s", footballMatches[i].date, footballMatches[i].time))
		if err != nil {
			sortingError = fmt.Errorf("error parsing date: %v\n", err)
			fmt.Println(sortingError)
			return false
		}
		jDateTime, err := time.Parse("02/01/2006 15:04", fmt.Sprintf("%s %s", footballMatches[j].date, footballMatches[j].time))
		if err != nil {
			sortingError = fmt.Errorf("error parsing date: %v\n", err)
			fmt.Println(sortingError)
			return false
		}
		return iDateTime.After(jDateTime)
	})

	return sortingError
}

func FetchFootballMatches(url string) ([]FootballMatch, error) {

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	csvReader := csv.NewReader(resp.Body)
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return nil, err
	}

	var allData []FootballMatch
	for _, record := range records[1:] {

		footballData := FootballMatch{
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

	return allData, nil
}

func main() {
	log.Println("Running Program...")
	//default value
	url := "https://www.football-data.co.uk/mmz4281/1920/E0.csv"
	if len(os.Args) > 1 {
		url = os.Args[1]
	}

	fmt.Printf("Fetching match data from %s\n", url)
	matches, err := FetchFootballMatches(url)
	if err != nil {
		fmt.Printf("FetchFootballMatches error: %v\n", err)
		return
	}

	err = SortAccendingDate(matches)
	if err != nil {
		fmt.Printf("SortAccendingDate error: %v\n", err)
		return
	}

	PrintFootballMatches(matches)
}
