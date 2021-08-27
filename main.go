package main

import (
	"fmt"
	"github.com/ianling/gopher-lua-example/dog"
	"github.com/ianling/gopher-lua-example/scripting"
	lua "github.com/yuin/gopher-lua"
)

func main() {
	L := scripting.InitializeLua()
	defer L.Close()

	// run the script
	err := L.DoFile("main.lua")
	if err != nil {
		panic("error running lua script: " + err.Error())
	}

	// retrieve the my_dog object from the Lua script
	myDog, err := dog.FromLua(L.GetGlobal("my_dog").(*lua.LUserData))
	if err != nil {
		panic(fmt.Sprintf("failed to retrieve dog from Lua state: %s", err.Error()))
	}

	// run the Speak function that we set in the Lua script
	fmt.Println("asking the dog to speak from Go...")
	myDog.Speak()
}
