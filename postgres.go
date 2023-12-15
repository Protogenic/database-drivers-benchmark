package main

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "DBLab03"
)

func postgresUtil(query int) {
	db, _ := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	var rows *sql.Rows
	var err error
	switch query {
	case 1:
		for i := 0; i < 10; i++ {
			rows, err = db.Query("SELECT VendorID, count(*) FROM trips GROUP BY 1")

			if err != nil {
				log.Fatal(err)
			}

			for rows.Next() {
				cabType := ""
				count := 0
				rows.Scan(&cabType, &count)
				//fmt.Printf("%s type: %d\n", cabType, count)
			}
			//fmt.Println()
		}
	case 2:
		for i := 0; i < 10; i++ {
			rows, err = db.Query("SELECT passenger_count, avg(total_amount) FROM trips GROUP BY 1")
			if err != nil {
				log.Fatal(err)
			}

			for rows.Next() {
				passengerCount := 0
				avg := 0.0
				rows.Scan(&passengerCount, &avg)
				//fmt.Printf("%d passengers: %0.2f dollars\n", passengerCount, avg)
			}
			//fmt.Println()
		}
	case 3:
		for i := 0; i < 10; i++ {
			rows, err = db.Query("SELECT passenger_count, extract (year from pickup_datetime), count(*) FROM trips GROUP BY 1, 2")

			if err != nil {
				log.Fatal(err)
			}

			for rows.Next() {
				passengerCount := ""
				date := 0
				count := 0
				rows.Scan(&passengerCount, &date, &count)
				//fmt.Printf("%s passengers, year %d: %d\n", passengerCount, date, count)
			}
			//fmt.Println()
		}
	case 4:
		for i := 0; i < 10; i++ {
			rows, err = db.Query("SELECT passenger_count, extract (year from pickup_datetime),  round(trip_distance), count(*) FROM trips GROUP BY 1, 2, 3 ORDER BY 2, 4 desc")
			if err != nil {
				log.Fatal(err)
			}

			for rows.Next() {
				passengerCount := ""
				date := 0
				distance := 0
				count := 0
				rows.Scan(&passengerCount, &date, &distance, &count)
				//fmt.Printf("%s passengers, %d, %d miles: %d\n", passengerCount, date, distance, count)
			}
			//fmt.Println()
		}
	default:
		fmt.Println("Wrong query number.")
	}
}
