package db_handle

import (
    "fmt"
    "database/sql"
    
)

func CloseDB(db *sql.DB){
    db.Close()
    fmt.Println("database closed")
}
