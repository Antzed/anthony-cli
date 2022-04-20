package file_handle

import (
    "bufio"
    "fmt"
    "os"
     er "github.com/Antzed/anthony-cli/error_handle"
)

func ScansToList(file *os.File, jobs []string) []string {
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        jobs = append(jobs, scanner.Text())
    }
    fmt.Println(jobs)
    fmt.Println("scanner done")
    err2 := scanner.Err()
    er.CheckErr(err2)
    return jobs
}
