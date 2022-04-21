package file_handle

import (
    "bufio"
    "fmt"
    "os"
    "strings"
     er "github.com/Antzed/anthony-cli/error_handle"
)

func ScansToList(file *os.File, jobs []string) []string {
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        temp := scanner.Text()
        temp = strings.Replace(temp, "|", " is a ", 1)
        temp = strings.Replace(temp, "|", " and is due ", 1)
        fmt.Println("temp is ", temp)
        jobs = append(jobs, temp)
    }
    fmt.Println(jobs)
    fmt.Println("scanner done")
    err2 := scanner.Err()
    er.CheckErr(err2)
    return jobs
}

