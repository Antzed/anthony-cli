package lua_handle

import (
    er "github.com/Antzed/anthony-cli/error_handle"
     "github.com/yuin/gopher-lua"
)

func CallNoParamFunc(L *lua.LState, funcName string, returnNum int){
    err := L.CallByParam(lua.P{
        Fn:      L.GetGlobal(funcName), // name of Lua function
        NRet:    returnNum,                     // number of returned values
        Protect: true,                  // return err or panic
    });
    er.CheckErr(err)
}

func CallOneStrParamFunc(L *lua.LState, funcName string, returnNum int, param string){
    err := L.CallByParam(lua.P{
         Fn:      L.GetGlobal(funcName), // name of Lua function
         NRet:    returnNum,                     // number of returned values
         Protect: true,                  // return err or panic
    }, lua.LString(param));
     er.CheckErr(err)
}
