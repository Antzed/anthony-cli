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

func SelectJobByDueday(db *sql.DB, dueDay string){
    var instruction = "SELECT j.JobID, j.JobName, jt.JobTypeName,  j.DueDate       FROM JOB j JOIN JOB_TYPE jt ON j.JobTypeID = jt.JobTypeID WHERE j.DueDate = '" + dueDay + "'"
    rows, err := db.Query(instruction)
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


