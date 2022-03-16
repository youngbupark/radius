package bindings

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

func MicrosoftSqlBinding(envParams map[string]string) BindingStatus {
	// From https://docs.microsoft.com/en-us/azure/azure-sql/database/connect-query-go
	connString := envParams["CONNECTIONSTRING"]
	if connString == "" {
		log.Println("CONNECTIONSTRING is required")
		return BindingStatus{false, "CONNECTIONSTRING is required"}
	}
	db, err := sql.Open("sqlserver", connString)
	defer db.Close()
	if err != nil {
		log.Println("Error creating connection pool - ", err.Error())
		return BindingStatus{false, "Connection to mssql db failed"}
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Println("failed to ping mssqldb - ", err.Error())
		return BindingStatus{false, "failed to ping mssqldb"}
	}
	stmt, err := db.Prepare("select 1 as number")
	if err != nil {
		log.Println("mssql access failed - ", err.Error())
		return BindingStatus{true, "database access failed"}
	}
	defer stmt.Close()
	row := stmt.QueryRow()
	var newID int64
	err = row.Scan(&newID)
	if err != nil {
		log.Println("mssql access failed - ", err.Error())
		return BindingStatus{true, "mssql database access failed"}
	}
	return BindingStatus{true, "mssql database accessed"}
}