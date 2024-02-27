package testes

import (
	"fmt"
	"time"
)

func TestTime() {

	location, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		print("erro ao carregar fuso", err)
	}

	now := time.Date(2024, 01, 02, 18, 02, 30, 20, location)

	println("Data e Hora:", now.Format("02/01/2006 15:04:05"))
	i := int64(2)
	now = now.AddDate(0, int(i-1), 0)

	println("Data e Hora:", now.Format("02/01/2006 15:04:05"))
}

func TestDateInstallment() {

	const FIRST_SCHEDULE_DUE_DATE = "2023-01-05"
	const FIRST_INSTALLMENT_DUE_DATE = "2023-02-01"

	parcelas := []int{
		1,
	}

	result := make(map[int]time.Time)

	for i := 0; i < len(parcelas); i++ {

		var dueDate time.Time
		var err error

		if len(parcelas) == 1 || i == 0 {
			dueDate, err = time.Parse("2006-01-02", FIRST_SCHEDULE_DUE_DATE)
			if err != nil {
				fmt.Println(err)
			}
			result[parcelas[i]] = dueDate
		} else {
			dueDate, err = time.Parse("2006-01-02", FIRST_INSTALLMENT_DUE_DATE)
			if err != nil {
				fmt.Println(err)
			}
			dueDate = dueDate.AddDate(0, i-1, 0)
			result[parcelas[i]] = dueDate
		}
		fmt.Println(result[parcelas[i]])
	}

}
