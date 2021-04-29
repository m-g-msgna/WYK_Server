package infrastructures

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"wyk_server.src/interfaces"
)

type MySQLHandler struct {
	Conn *sql.DB
}

func (handler *MySQLHandler) Connection() *sql.DB {
	return handler.Conn
}

func (handler *MySQLHandler) Execute(statement string) (sql.Result, error) {
	res, err := handler.Conn.Exec(statement)
	if err != nil {
		panic(err)
	}

	return res, err
}

//For updating multiple tables at the same time. Such as when creating a record across multiple tables.
func (handler *MySQLHandler) ExecuteMultipleStatements(statements []string) error {
	tx, err := handler.Conn.Begin()
	if err != nil {
		panic(err)
		//return err
	}

	for i := range statements {
		_, err = handler.Conn.Exec(statements[i])
		if err != nil {
			tx.Rollback()
			panic(err)
			//return err
		}
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
		//return err
	}

	return nil
}

func (handler *MySQLHandler) Query(statement string) (interfaces.IRow, error) {
	rows, err := handler.Conn.Query(statement)

	if err != nil {
		//panic(err)
		return new(MySQLRow), err
	}
	row := new(MySQLRow)
	row.Rows = rows

	return row, nil
}

type MySQLRow struct {
	Rows *sql.Rows
}

func (r MySQLRow) Scan(dest ...interface{}) error {
	err := r.Rows.Scan(dest...)
	if err != nil {
		return err
	}

	return nil
}

func (r MySQLRow) Next() bool {
	return r.Rows.Next()
}

func (r MySQLRow) Close() {
	r.Rows.Close()
}
