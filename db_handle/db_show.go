package db_handle

import (
    "fmt"
    "database/sql"
    "time"

    er "github.com/Antzed/anthony-cli/error_handle"
)

func ShowJob(db *sql.DB){
    rows, err := db.Query("SELECT j.JobID, j.JobName, jt.JobTypeName,  j.DueDate  FROM JOB j JOIN JOB_TYPE jt ON j.JobTypeID = jt.JobTypeID")
    er.CheckErr(err)
    //fmt.Println("rows: ",strconv.Itoa(rows))
    var jid int
    var jname string
    var jtype string
    var jduedate time.Time
    fmt.Println("jobID", "JobName", "JobTypeName", "Duedate")
    for rows.Next() {
        err = rows.Scan(&jid, &jname, &jtype, &jduedate)
        er.CheckErr(err)
        fmt.Println(jid, jname, jtype, jduedate)
    }
    rows.Close()
}

