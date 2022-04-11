package task_handle

import (
    "fmt"
    "os/exec"
    er "github.com/Antzed/anthony-cli/error_handle"
)

func AddTask(inputsplit []string, board string){
     if board != "My Board" {
         var atboard string = "@" + board
         for _, s := range inputsplit {
             fmt.Println("added task: ", s, " at board: ", board)
             cmd := exec.Command("tb", "-t", atboard, s)
             err := cmd.Run()

             er.CheckErr(err)
             fmt.Println("done")
         }
     } else {
         for _, s := range inputsplit {
             fmt.Println("added task: ", s)
             cmd := exec.Command("tb", "-t" ,s)
             err := cmd.Run()

             er.CheckErr(err)
             fmt.Println("done1")

         }
     }
 }
