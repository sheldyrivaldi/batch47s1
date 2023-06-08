package utilities

import (
	"fmt"
	"strconv"
	"time"
)

func GetDuration(startDate string, endDate string) string {
	var duration string

	dateLayout := "2006-01-02"
	startDateParse, err := time.Parse(dateLayout, startDate )
	if err != nil {
		fmt.Println("Parsing Date Error", err)
	}
	endDateParse, err := time.Parse(dateLayout, endDate)
	if err != nil {
		fmt.Println("Parsing Date Error", err)
	}

	difference := endDateParse.Sub(startDateParse).Hours()


	day := int(difference / 24)
	week := day / 7
	month := week / 4
	year := month / 12

	if day >= 0 {
		duration = strconv.Itoa(day) + "hari"
	}
	if week > 0 {
		duration = strconv.Itoa(week) + " minggu"
	}
	if month > 0 {
		duration = strconv.Itoa(month) + " bulan"
	}
	if year > 0 {
		duration = strconv.Itoa(year) + " tahun"
	}

	return duration
}

func GetDurationFormat(date string) string {
	dateLayout := "2006-01-02"
	dateParse, err := time.Parse(dateLayout, date )
	if err != nil {
		fmt.Println("Parsing Date Error", err)
	}

	return dateParse.Format("02 Jan 2006")

}