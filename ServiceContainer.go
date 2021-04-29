package main

import (
	"sync"

	"database/sql"

	"wyk_server.src/controllers"
	"wyk_server.src/infrastructures"
	"wyk_server.src/repositories"
	"wyk_server.src/services"
)

type IServiceContainer interface {
	InjectWYKController(sqlConn *sql.DB) controllers.WYKController
}

type kernel struct{}

func DbConnectionInjector() *sql.DB {
	sqlConn, _ := sql.Open("mysql", "root:meh1hfgi!@tcp(localhost:3306)/wyk_service_db")
	return sqlConn
}

func (k *kernel) InjectWYKController(sqlConn *sql.DB) controllers.WYKController {
	mysqlHandler := &infrastructures.MySQLHandler{}
	mysqlHandler.Conn = sqlConn

	wykRepository := &repositories.WYKRepository{mysqlHandler} //
	wykService := &services.WYKService{&repositories.WYKRepository{wykRepository}}
	wykController := controllers.WYKController{wykService}

	return wykController
}

var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}
