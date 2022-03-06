package query

import (
	"database/sql"
	"log"
	"strconv"
	"reflect"
	"time"

	"github.com/blockloop/scan"
)
// GET/:id
func ReadRecord(db *sql.DB, v interface{}, id string, table_name string) {
	// Read existing records.
	query_string := "SELECT * FROM " + table_name + " WHERE id = ?"
	rows, err := db.Query(query_string, id)
	if err != nil {
		log.Fatal(err)
	}
	err = scan.Row(v, rows)
	if err != nil {
		log.Fatal(err)
	}
	// id: String to Int
	id_int64, _ := strconv.ParseInt(id, 10, 64)
	val := reflect.ValueOf(v).Elem()
	val.Field(0).SetInt(id_int64)
}

// GET all
func ReadAllRecords(db *sql.DB, v interface{}, table_name string) {
	// Read existing records.
	query_string := "SELECT * FROM " + table_name
	rows, err := db.Query(query_string)
	if err != nil {
		log.Fatal(err)
	}
	err = scan.Rows(v, rows)
	if err != nil {
		log.Fatal("scan: ", err)
	}
	defer rows.Close()
}

func DeleteRecord(db *sql.DB, id string, table_name string) {
	// Update the DeletedAt column. time now.
	query_string := "UPDATE " + table_name + " SET deleted_at = ? WHERE id = ?"
	stmt, err := db.Prepare(query_string)
	if err != nil {
		log.Fatal(err)
	}
	
	id_int, _ := strconv.Atoi(id)
	_ , err = stmt.Exec(time.Now(), id_int)
	if err != nil {
		log.Fatal(err)
	}
}