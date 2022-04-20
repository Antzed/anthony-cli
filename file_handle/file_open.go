package file_handle

import (
    "os"
    er "github.com/Antzed/anthony-cli/error_handle"

)

func OpenFile(filepath string)(file *os.File){
    file, err := os.Open(filepath)
    er.CheckErr(err)
    return file
}
