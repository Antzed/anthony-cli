package db_handle

import (
    "fmt"
    "database/sql"
     er "github.com/Antzed/anthony-cli/error_handle"
)

func UpdateJobColumn(db *sql.DB, jobID int, columnValue string, toValue string) { 
    if columnValue == "JobTypeName" {
        condition := columnValue + " = '" + toValue +"'"
        tempID := SelectForeignKey(db, "JobTypeID", "JOB_TYPE", condition)
        stmt, err := db.Prepare("UPDATE JOB SET JobTypeID = ? WHERE JobID = ?")
        er.CheckErr(err)
        _, err1 := stmt.Exec(tempID, jobID)
        er.CheckErr(err1)
    } else {
        stmt, err := db.Prepare("UPDATE JOB SET "+ columnValue +" = ? WHERE JobID = ?")
        er.CheckErr(err)
        _, err1 := stmt.Exec(toValue, jobID)
        er.CheckErr(err1)
    }
    
    fmt.Println("sucessfule updated job")
}

