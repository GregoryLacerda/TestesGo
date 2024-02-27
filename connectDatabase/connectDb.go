package connectdatabase

import (
	"Testes/constants"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {

	db, err := sql.Open("mysql", constants.USER+":"+constants.PASSWORD+"@tcp("+constants.URL+")/acordo_certo")
	if err != nil {
		fmt.Println("Falha ao conectar no banco:")
		panic(err)
	}
	return db

}

func ShowAll() {

	selectAllInstallmentAgreementID := "SELECT acordo_id FROM suspd_acordo_parcela sap " +
		"INNER JOIN suspd_acordo ON sap.acordo_id = suspd_acordo.id " +
		"WHERE suspd_acordo.data_acordo = '2023-12-27' " +
		"AND suspd_acordo.credor_id = 84 " +
		"AND sap.parcela = 2 " +
		"AND sap.data_vencimento = '2024-03-02' "

	/* selectAllDueDate := "SELECT data_vencimento FROM suspd_acordo_parcela sap" +
	"WHERE sap.acordo_id = %i" +
	"AND sap.data_vencimento > '2024-01-02'" */

	db := ConnectDB()

	allAgreemnetsId, err := db.Query(selectAllInstallmentAgreementID)
	if err != nil {
		fmt.Println("Falha ao conectar no banco:")
		panic(err)
	}

	var count int
	for allAgreemnetsId.Next() {
		var agreementID int
		allAgreemnetsId.Scan(&agreementID)
		fmt.Println(agreementID)
		count++
	}

	fmt.Println(count)
}
