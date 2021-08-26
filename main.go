package main

import (
	"github.com/ianling/gopher-lua-example/scripting"
)

func main() {
	L := scripting.InitializeLua()
	defer L.Close()

	err := L.DoFile("main.lua")
	if err != nil {
		panic("error running lua script: " + err.Error())
	}
}
