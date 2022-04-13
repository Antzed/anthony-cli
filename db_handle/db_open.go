package db_handle

import (
    "fmt"

    "database/sql"
    _"github.com/mattn/go-sqlite3"
    
    er "github.com/Antzed/anthony-cli/error_handle"

)

func OpenDB(dbName string, dbPath string)(db *sql.DB){
    dbFullPath := dbPath + dbName +".db"
    fmt.Println("dbFullPath: ", dbFullPath)
    db, err := sql.Open("sqlite3", dbFullPath)
    er.CheckErr(err)
    fmt.Println("opened database")
    return db
}
