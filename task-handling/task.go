package task

import (
    "fmt"
    "os/exec"
    "github.com/anthony-cli/error-handle/error_handle"
)

func addTask(inputsplit []string, board string){
     if board != "My Board" {
         var atboard string = "@" + board
         for _, s := range inputsplit {
             fmt.Println("added task: ", s, " at board: ", board)
             cmd := exec.Command("tb", "-t", atboard, s)
             err := cmd.Run()

             error_handle.checkErr(err)
             fmt.Println("done")
         }
     } else {
         for _, s := range inputsplit {
             fmt.Println("added task: ", s)
             cmd := exec.Command("tb", "-t" ,s)
             err := cmd.Run()

             checkErr(err)
             fmt.Println("done1")

         }
     }
 }
