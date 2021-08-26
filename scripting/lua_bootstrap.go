package scripting

import (
	"github.com/ianling/gopher-lua-example/common"
	"github.com/ianling/gopher-lua-example/dog"
	lua "github.com/yuin/gopher-lua"
)

// luaTypes is a slice of all the LuaTypeExport objects we've defined in other packages for making Go things available
// in Lua.
var luaTypes = []common.LuaTypeExport{
	dog.LuaTypeExport,
}

// registerType takes a LuaTypeExport
func registerType(L *lua.LState, luaTypeExport common.LuaTypeExport) {
	typeMetatable := L.NewTypeMetatable(luaTypeExport.Name)
	L.SetGlobal(luaTypeExport.Name, typeMetatable)

	// static attributes
	L.SetField(typeMetatable, "new", L.NewFunction(luaTypeExport.ConstructorFunc))

	// methods
	L.SetField(typeMetatable, "__index", L.SetFuncs(L.NewTable(), luaTypeExport.Methods))
}

func InitializeLua() *lua.LState {
	L := lua.NewState()

	for _, luaType := range luaTypes {
		registerType(L, luaType)
	}

	return L
}
