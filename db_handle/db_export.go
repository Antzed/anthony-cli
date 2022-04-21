package db_handle

import (
    "os"
    //"os/exec"
    "fmt"
    "time"
    er "github.com/Antzed/anthony-cli/error_handle"
    "github.com/Antzed/anthony-cli/lua_handle"
    "github.com/Antzed/anthony-cli/path_handle"

)
var projectPath = path_handle.RootDir()
var jobtxtPath = projectPath + "/job.txt"
func ExportJob(){
    if _, err := os.Stat(jobtxtPath); err == nil {
        e := os.Remove(jobtxtPath)
        er.CheckErr(e)
    }

    //var luaScript = lua_handle.luaScript
    
    L := lua_handle.InitScriptString(lua_handle.LuaScript)
    defer L.Close()

    lua_handle.CallNoParamFunc(L, "exportJob", 0)
    
}

func ExportJobFromWeek(due string){
    if _, err := os.Stat(jobtxtPath); err == nil {
         e := os.Remove(jobtxtPath)
         er.CheckErr(e)
     }

     L := lua_handle.InitScriptString(lua_handle.LuaScript)
     //fmt.Println("due is " + due)
     defer L.Close()

     var currentTime time.Time
     if(due != "null"){
        var err error
        currentTime, err = time.Parse("2006-01-02", due)
        fmt.Println("changed to specified duedate")
        er.CheckErr(err)
     } else {
        currentTime = time.Now()
     }
     aWeekFromNow := currentTime.AddDate(0,0,7)
     fmt.Println(currentTime, aWeekFromNow)
     var timeRange string =  "'" + currentTime.Format("2006-01-02") + "'" + " AND " + "'" + aWeekFromNow.Format("2006-01-02") + "'"
     fmt.Println(timeRange)

     lua_handle.CallOneStrParamFunc(L, "exportJobThisWeek", 0, timeRange)
}


