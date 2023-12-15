package main

import (
	"fmt"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

func main() {
	query := 0
	fmt.Println("Pick query number: ")
	fmt.Scan(&query)

	fmt.Println("Choose Driver (number):\n1. Postgres\n2. Sqlite3\n3. Dataframe (pandas)")
	pick := 0
	fmt.Scan(&pick)

	timeValue := time.Now()
	switch pick {
	case 1:
		postgresUtil(query)
	case 2:
		sqliteUtil(query)
	case 3:
		dataframeUtil(query)
	default:
		fmt.Println("No such driver number.")
	}
	duration := time.Since(timeValue)
	fmt.Println("Average elapsed time (ms): ", duration.Milliseconds()/10)
}
