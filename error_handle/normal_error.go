package error_handle
import (
    "fmt"
    "os"
)

type DBInvalidInstructionError struct{}

func (dbErr *DBInvalidInstructionError) Error() string {
    return "Invalid instruciton, try type in the right syntax"
}
    
func CheckErr(err error) {
     if err != nil {
         switch e := err.(type) {
         case *DBInvalidInstructionError :
             fmt.Println("boooooooo")
             os.Exit(1)
         default:
             panic(e)
         }
     }   
}
