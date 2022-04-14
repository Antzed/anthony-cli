package lua_handle

import (
    er "github.com/Antzed/anthony-cli/error_handle"
    "github.com/yuin/gopher-lua"
)

func InitScriptString(luaScript string)(L *lua.LState){
    L = lua.NewState()
    err := L.DoString(luaScript)
    er.CheckErr(err)
    return L
}
