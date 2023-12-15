package main

import (
	"github.com/go-gota/gota/dataframe"
	"log"
	"os"
)

func dataframeUtil(query int) {
	file, err := os.Open("nyc_yellow_tiny.csv")

	if err != nil {
		log.Fatal(err)
	}

	df := dataframe.ReadCSV(file)

	switch query {
	case 1:
		for i := 0; i < 10; i++ {
			df.Select("SELECT VendorID, count(*) FROM trips GROUP BY 1")
		}
	case 2:
		for i := 0; i < 10; i++ {
			df.Select("SELECT passenger_count, avg(total_amount) FROM trips GROUP BY 1")
		}
	case 3:
		for i := 0; i < 10; i++ {
			df.Select("SELECT passenger_count, extract (year from pickup_datetime), count(*) FROM trips GROUP BY 1, 2")
		}
	case 4:
		for i := 0; i < 10; i++ {
			df.Select("SELECT passenger_count, extract (year from pickup_datetime),  round(trip_distance), count(*) FROM trips GROUP BY 1, 2, 3 ORDER BY 2, 4 desc")
		}
	}
}
