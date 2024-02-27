package connectdatabase

import (
	"database/sql"
	"fmt"
	"log"
)

func executa() {

	db, err := sql.Open("mysql", "aca:!4GVHG!wCLG7u.TXe3R!NbtZ@tcp(35.247.224.117:3306)/acordo_certo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Conexão com o banco de dados aberta com sucesso!")

	rows, err := db.Query(`
	SELECT DISTINCT sap.acordo_id
	FROM acordo_certo.suspd_acordo_parcela sap
	JOIN acordo_certo.suspd_acordo sa ON sap.acordo_id = sa.id
	WHERE (sap.acordo_id, sap.data_vencimento) IN (
		SELECT acordo_id, data_vencimento
		FROM acordo_certo.suspd_acordo_parcela
		WHERE data_criacao >= '2023-10-01 00:00:00'
		GROUP BY acordo_id, data_vencimento
		HAVING COUNT(*) > 1
	)
	AND sap.data_criacao >= '2023-10-01 00:00:00'
	AND sa.credor_id = 95
	ORDER BY sap.acordo_id;
	`)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {

		var acordoID int
		rows.Scan(&acordoID)

		fmt.Println(acordoID)
		installments, err := db.Query(`
		SELECT id, parcela from suspd_acordo_parcela sap 
		WHERE acordo_id = ?
		ORDER by parcela ASC
		`, acordoID)
		if err != nil {
			log.Fatal(err)
		}
		defer installments.Close()

		parcelas := make(map[int]int)
		for installments.Next() {
			var id, parcela int
			err = installments.Scan(&id, &parcela)
			if err != nil {
				log.Fatal(err)
			}

			parcelas[id] = parcela
		}

		vistos := make(map[int]bool)
		parcelasVistas := make(map[int]int)

		for id, parcela := range parcelas {

			if vistos[parcela] {
				if parcelasVistas[parcela] > id {
					db.Query(`DELETE FROM suspd_acordo_parcela WHERE id = ?`, parcelasVistas[parcela])
				} else {
					db.Query(`DELETE FROM suspd_acordo_parcela WHERE id = ?`, id)
				}

			} else {
				vistos[parcela] = true
				parcelasVistas[parcela] = id
			}

		}

	}

}
