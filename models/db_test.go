package models

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/go-gorp/gorp"
)

func TestDb(t *testing.T) {
	dataSource := fmt.Sprintf("root:root@tcp(127.0.0.1:3306)/stakepool?charset=utf8mb4")
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		fmt.Println("sql.Open failed: ", err)
		return
	}

	// sql.Open just validates its arguments without creating a connection, so
	// verify that the data source name is valid with Ping.
	if err = db.Ping(); err != nil {
		fmt.Println("Unable to establish connection to database: ", err)
		return
	}

	// Construct a gorp DbMap.
	dbMap := &gorp.DbMap{
		Db:              db,
		Dialect:         gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8MB4"},
		ExpandSliceArgs: true,
	}

	// Add a table, setting the table name and specifying that the Id property
	// is an auto incrementing primary key
	dbMap.AddTableWithName(EmailChange{}, "EmailChange").SetKeys(true, "Id")
	dbMap.AddTableWithName(LowFeeTicket{}, "LowFeeTicket").SetKeys(true, "Id")
	dbMap.AddTableWithName(PasswordReset{}, "PasswordReset").SetKeys(true, "Id")
	dbMap.AddTableWithName(Session{}, "Session").SetKeys(true, "Id")
	usersTableName := "Users"
	dbMap.AddTableWithName(User{}, usersTableName).SetKeys(true, "Id")

	// Create the table.
	err = dbMap.CreateTablesIfNotExists()
	if err != nil {
		fmt.Println("Create tables failed: ", err)
		// There is no point proceeding, so return with nil.
		return
	}
	fmt.Println("suc")
}
