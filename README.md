gopher-lua example
============
`main.lua` is the Lua script. The file names are all arbitrary.

The `dog` folder contains a very simple `Dog` struct that can be used in Go,
along with a file that defines methods that can be used from Lua.
The top of the `lua_compatibility.go` contains a little bit of glue that defines
which methods do what.

The `scripting` folder contains the weird glue code that facilitates the interoperability between Lua and Go.

Usage
-----
Running `go run main.go` will initialize all the Go and Lua glue stuff,
before finally executing the `main.lua` script:

```lua
my_dog = dog.new("Dingus")
print(my_dog:name())

my_dog:name("Bingus")
print(my_dog:name())
```

This script prints:

```
Dingus
Bingus
```

Adding another field to Dog
--------------
There is nothing special required to add another field to the Dog struct
in Go, you would simply add it to the struct definition as usual.

To export the new field to Lua, add a new getter/setter function to `dog/lua_compatibility.go`,
then add a new value to the `LuaTypeExport.Methods` map at the top of the file.

To allow setting this field in the Dog constructor, add the field to the struct initialization
at the top of the `newDogLua` function in `dog/lua_compatibility.go`.

Exporting a different struct
---------------------
The simplest way to export a different struct would be to copy the `dog` folder to use
as a base for the new struct, and then adjust the constructor and getters/setters according
to the new struct's fields.

The final step is to add the new struct's `LuaTypeExport` variable to the 
`luaTypes` slice at the top of `scripting/lua_bootstrap.go`.
After that, the new type should be accessible from Lua scripts.