package db

import (
	"database/sql"
	"log"
	"time"

	//for postgres support
	_ "github.com/lib/pq"
)

//DB connection variable
var DB *sql.DB
var err error

//Connect function performs the connection operation to the postgres database
func Connect() error {
	DB, err = sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/tourist")
	return err
}

//Rollback is used to rollback a transaction incase of error
func Rollback() {
	if _, err := DB.Query(`rollback`); err != nil {
		log.Printf("Error rolling back")
	}
}

//NewNullString converts the empty string to null
//before inserting into the database
func NewNullString(val string) sql.NullString {
	if val == "" {
		return sql.NullString{}
	}
	return sql.NullString{
		String: val,
		Valid:  true,
	}
}

//NewNullID converts the empty reference to null
//before inserting into the database
func NewNullID(val int64) sql.NullInt64 {
	if val == 0 {
		return sql.NullInt64{}
	}
	return sql.NullInt64{
		Int64: val,
		Valid: true,
	}
}

//NewNullDate converts the empty date to null
//before inserting into the database
func NewNullDate(val string) sql.NullTime {
	if val == "" {
		return sql.NullTime{}
	}
	date, err := time.Parse("2006-01-02", val)
	if err != nil {
		return sql.NullTime{}
	}
	return sql.NullTime{
		Time:  date,
		Valid: true,
	}
}
