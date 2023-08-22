package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

const (
	host     = "ec2-3-213-146-52.compute-1.amazonaws.com"
	port     = 5432
	user     = "srgdxuyweclewc"
	password = "611f2569e0778feff2c815d6a48f2b5baf4fcc6959c541bf13837db6a806a281"
	dbname   = "dblis5f8n48gel"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func main() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)
	fmt.Println("Connected!")

	merchants_id_query := `SELECT id FROM merchants`
	merchants_id_rows, err := db.Query(merchants_id_query)
	CheckError(err)

	private_key := 100000
	for merchants_id_rows.Next() {
		var id string

		err = merchants_id_rows.Scan(&id)
		CheckError(err)

		payident_query := fmt.Sprintf(`select pi2.merchant_id,pi2."type" from merchants m inner join payment_identities pi2 on m.id = pi2.merchant_id WHERE merchant_id = %s;`, id)
		payident_rows, err := db.Query(payident_query)

		var installed_payments []string
		for payident_rows.Next() {
			var merchant_id string
			var pay_type string

			err = payident_rows.Scan(&merchant_id, &pay_type)
			CheckError(err)

			installed_payments = append(installed_payments, pay_type)
		}

		fmt.Println(id, installed_payments)

		types_list := []string{"BANCOLOMBIA_COLLECT_SANDBOX", "CARD_SANDBOX", "BANCOLOMBIA_TRANSFER_SANDBOX", "NEQUI_SANDBOX", "PSE_SANDBOX"}

		for _, s := range types_list {
			if Contains(installed_payments, s) {
				fmt.Println(s, "-- FOUND")
			} else {
				fmt.Println(s, "-- NOT FOUND -- INSERTING...")
				insert_query := fmt.Sprintf(`INSERT INTO payment_identities VALUES('%s', '2020-11-26 20:58:05', '2020-11-26 20:58:05', '%s', '{}', '%s' , 'DEFAULT', true, true, '%s')`, strconv.Itoa(private_key), id, s, strings.Replace(s, "_SANDBOX", "", -1))
				fmt.Println(insert_query)
				_, err := db.Exec(insert_query)
				CheckError(err)

			}
			private_key = private_key + 1
		}

		fmt.Println("  ")

	}

}
