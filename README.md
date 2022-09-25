# Football Match Data Parser
This program gets football match data, sorts the matches by date (latest to earliest) and then prints it.

## Running
To run this program:

```bash
go run main.go
```

This by default fetches the data from https://www.football-data.co.uk/mmz4281/1920/E0.csv

Optionally you can specify a URL to get the data from by passing it in as the first command line arg

```bash
go run main.go https://www.football-data.co.uk/mmz4281/1920/E1.csv
```

## Building
To build the program into a single executable:

```bash
go build main.go
```

And then to use:

```bash
./main https://www.football-data.co.uk/mmz4281/1920/E1.csv
```

