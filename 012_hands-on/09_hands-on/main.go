package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"
)

type financialStats struct {
	Income, Expensis, Balance float64
	Date                      time.Time
}

func main() {
	parseCsv("table.csv")
}

func parseCsv(fileName string) []financialStats {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	records := make([]financialStats, 0, len(rows))

	for i, row := range rows {
		if i == 0 {
			continue
		}

		date, _ := time.Parse("02-01-2006", row[0])
		income, _ := strconv.ParseFloat(row[1], 64)
		expensis, _ := strconv.ParseFloat(row[1], 64)
		balance, _ := strconv.ParseFloat(row[1], 64)

		records = append(records, financialStats{
			Income:   income,
			Expensis: expensis,
			Balance:  balance,
			Date:     date,
		})
	}

	return records
}
