package interfaces

import "database/sql"

type IDbHandler interface {
	Connection() *sql.DB
	Execute(statement string) (sql.Result, error)
	ExecuteMultipleStatements(statements []string) error
	Query(statement string) (IRow, error)
}

type IRow interface {
	Scan(dest ...interface{}) error
	Next() bool
	Close()
	//RowsAffected() int
}
