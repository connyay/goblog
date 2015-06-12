package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db, _ = sql.Open("mysql", "root@tcp(localhost:3306)/goblog")

type statementMap map[string]*sql.Stmt

var statementMaps = make(map[string]statementMap)

func GetDB() *sql.DB {
	return db
}

func PreparedMap(name string) statementMap {
	if statementMaps[name] == nil {
		statementMaps[name] = statementMap{}
	}
	return statementMaps[name]
}

func GetPrepared(ns string, route string, query string) *sql.Stmt {
	if statementMaps[ns] == nil {
		statementMaps[ns] = statementMap{}
	}
	st := statementMaps[ns][route]
	if st == nil {
		st, _ = db.Prepare(query)
		statementMaps[ns][route] = st
	}
	return st
}
