package dates

import (
	"fmt"
	"strings"
	"time"
)

func ParseDate() time.Time {
	locBrSp := time.FixedZone("America/Sao_Paulo", -3*60*60)
	layout := "2006-01-02T15:04:05-07:00"
	parsedDate, err := time.ParseInLocation(layout, removeTimeFromDate("2024-01-27T00:00:00-03:00"), locBrSp)
	if err != nil {
		fmt.Println("erro ao converter")
		return time.Time{}
	}

	return parsedDate
}

func removeTimeFromDate(dateAsStr string) (result string) {
	if dateAsStr == "" {
		return ""
	}
	result = strings.Split(dateAsStr, "T")[0]
	if result != "" {
		return fmt.Sprintf("%sT00:00:00-03:00", result)
	}
	return dateAsStr
}

func ParsedDateTime() {
	dataAprovacao, err := time.Parse("2006-01-02T15:04:05.000000", "2024-02-15T10:38:39.930000")
	if err != nil {
		fmt.Println("Erro ao converter a string para time:", err)
	}
	fmt.Println(dataAprovacao)
}

func NegotiationDate() {

	nagtiationDate, err := time.Parse(time.RFC3339, "2024-01-27T00:00:00-03:00")
	if err != nil {
		fmt.Println("Erro ao converter a string para time:", err)
	}

	fmt.Printf("NegotiationDate: %s", nagtiationDate.Format("2006-01-02"))

	location, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		fmt.Println("Erro ao converter a string para time:", err)
	}
	nagtiationDate = nagtiationDate.In(location)

	fmt.Printf("NegotiationDate: %s", nagtiationDate)
}

func GetMinutesBeteewnDate(start time.Time, end time.Time) (retVal int64, err error) {
	if start.After(end) {
		return retVal, fmt.Errorf("start date is after end date")
	}

	minutes := int64(0)
	for start.Before(end) {
		start = start.Add(1 * time.Minute)
		minutes++
	}

	return minutes, nil
}
