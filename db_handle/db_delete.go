package db_handle

import (
    //"fmt"
    "database/sql"
    er "github.com/Antzed/anthony-cli/error_handle"
)

func DeleteJob(db *sql.DB, jobName string, dueDate string) {
    instruction := "DELETE FROM JOB WHERE JobName = '" + jobName + "' AND DueDate = '" + dueDate + "'"
    
    stmt, err := db.Prepare(instruction)
    er.CheckErr(err)
    _ ,err = stmt.Exec()
    er.CheckErr(err)
    
}

