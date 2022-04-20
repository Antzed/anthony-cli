package db_handle

import (
    "os"
    //"os/exec"
    "fmt"
    "time"
    er "github.com/Antzed/anthony-cli/error_handle"
    "github.com/Antzed/anthony-cli/lua_handle"

)

func ExportJob(){
    if _, err := os.Stat("./job.txt"); err == nil {
        e := os.Remove("job.txt")
        er.CheckErr(e)
    }

    //var luaScript = lua_handle.luaScript
    
    L := lua_handle.InitScriptString(lua_handle.LuaScript)
    defer L.Close()

    lua_handle.CallNoParamFunc(L, "exportJob", 0)
    
}

func ExportJobFromWeek(){
    if _, err := os.Stat("./job.txt"); err == nil {
         e := os.Remove("job.txt")
         er.CheckErr(e)
     }

     L := lua_handle.InitScriptString(lua_handle.LuaScript)
     defer L.Close()
     currentTime := time.Now()
     //currentDate := curerntTime.Date()
     aWeekFromNow := currentTime.AddDate(0,0,7)
     //aWeekFromNowDate := aWeekFromNow.Date()
     fmt.Println(currentTime, aWeekFromNow)
     var timeRange string =  "'" + currentTime.Format("2006-01-02") + "'" + " AND " + "'" + aWeekFromNow.Format("2006-01-02") + "'"
     fmt.Println(timeRange)

     lua_handle.CallOneStrParamFunc(L, "exportJobThisWeek", 0, timeRange)
}
