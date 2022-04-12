package db_handle

import(
    //"fmt"
    //"time"
    "database/sql"
    //_"github.com/mattn/go-sqlite3"
    er "github.com/Antzed/anthony-cli/error_handle"
   
)

func SelectForeignKey(db *sql.DB, idName string, tbName string, condition string)(id int){
    var queryForeignKey = "SELECT "+ idName +" FROM "+ tbName +" WHERE " + condition
    rows, err := db.Query(queryForeignKey)
    er.CheckErr(err)

    for rows.Next() {
        err = rows.Scan(&id)
        er.CheckErr(err)
    }
    return id
}
