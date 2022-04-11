package error_handle

func checkErr(err error) {
     if err != nil {
         panic(err)
     }
}
