package db_handle

import (
    "os"
    //"os/exec"
    //"fmt"
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
