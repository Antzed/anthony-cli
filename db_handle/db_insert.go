package db_handle

import (
    "fmt"
    "database/sql"
    _"github.com/mattn/go-sqlite3"
    er "github.com/Antzed/anthony-cli/error_handle"
    //"strings"
    //"regexp"
    //"time"
)

func InsertJob(db *sql.DB, jobName string, jtid int, dueDate string){
    stmt, err := db.Prepare("INSERT INTO JOB(JobName, JobTypeID, DueDate) values(?,?,?)")
    er.CheckErr(err)
    res, err := stmt.Exec(jobName, jtid, dueDate)
    er.CheckErr(err)
    id, err := res.LastInsertId()
    er.CheckErr(err)
    fmt.Println(id)
    fmt.Println("sucessfule inserted job")
}

//usage: db_handle.Insert(db, "INSERT INTO JOB(JobName, JobTypeID, DueDate) values('jobname','jobetype','duedate'))
//need to put 'db *sql.DB' back to arguments
func Insert(db *sql.DB, instruction string) {

    res, err := db.Exec(instruction)
    er.CheckErr(err.(DBInvalidInstructionError))
    id, err := res.LastInsertId()
    er.CheckErr(err)
    fmt.Println(id)
    fmt.Println("sucessfule inserted job")

}
