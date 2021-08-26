package dog

import (
	"github.com/ianling/gopher-lua-example/common"
	lua "github.com/yuin/gopher-lua"
)

// these two variables set up our Go struct for use in Lua scripts
var luaTypeExportName = "dog" // this sets the global name used to access this type from Lua
var LuaTypeExport = common.LuaTypeExport{
	Name:            luaTypeExportName,
	ConstructorFunc: newDogLua,
	Methods: map[string]lua.LGFunction{
		"name": nameGetterSetter, // from Lua: dog:name() -- returns the dog's Name field
	},
}

// newDogLua is the constructor for Dog objects in Lua scripts.
// Usage: dog.new(name)
// Example: my_dog = dog.new("Rufus")
func newDogLua(L *lua.LState) int {
	dog := &Dog{
		Name: L.CheckString(1), // first positional arg from Lua function call
	}

	ud := L.NewUserData()
	ud.Value = dog

	L.SetMetatable(ud, L.GetTypeMetatable(luaTypeExportName))
	L.Push(ud)

	return 1
}

// checkDog is a Go utility function that checks whether the first lua argument is a *LUserData representing a *Dog and
// returns this *Dog.
// This allows us to reliably translate a Lua Dog to a Go Dog.
func checkDog(L *lua.LState) *Dog {
	ud := L.CheckUserData(1)

	if vv, ok := ud.Value.(*Dog); ok {
		return vv
	}

	L.ArgError(1, "dog expected")

	return nil
}

// nameGetterSetter is a combined getter and setter for Dog's Name field.
// Examples from Lua:
// Getter: my_dog:name()
// Setter: my_dog:name("Barney")
func nameGetterSetter(L *lua.LState) int {
	dog := checkDog(L)

	// setter
	if L.GetTop() == 2 {
		dog.Name = L.CheckString(2)

		return 0
	}

	// getter
	L.Push(lua.LString(dog.Name))

	return 1
}
