package extensions

import (
	"fmt"
	"strings"

	"github.com/tortitast/ja/config"
	"github.com/tortitast/ja/utils"
	"github.com/urfave/cli/v2"
	lua "github.com/yuin/gopher-lua"
)

func EvalExtension(c *cli.Context, commands []*cli.Command, name string) {
	files, err := utils.GetFilesWithExtension(config.ExtensionsDir, ".lua")
	utils.Must(err, "Error getting extensions")

	l := lua.NewState()

	registerGlobals(c, commands, l)

	found := false
	for _, file := range files {
		filename := utils.GetFilenameFromURL(file)

		if !strings.HasPrefix(filename, name) {
			continue
		}

		if name != strings.Replace(filename, ".lua", "", 1) {
			response := utils.Prompt(fmt.Sprintf("Running extension %s, are you sure?", filename))
			if response != "y" {
				return
			}
		}

		err = l.DoFile(file)
		utils.Must(err, "Error loading file: "+file)

		found = true
	}

	if !found {
		utils.Print(fmt.Sprintf("Extension %s not found", name), utils.Error)
	}
}

func registerGlobals(c *cli.Context, commands []*cli.Command, L *lua.LState) {
	L.SetGlobal("call", L.NewFunction(func(L *lua.LState) int {
		return call(c, commands, L)
	}))

	argscpy := c.Args().Slice()
	L.SetGlobal("args", L.NewFunction(func(L *lua.LState) int {
		return args(argscpy, L)
	}))
}

func call(c *cli.Context, commands []*cli.Command, L *lua.LState) int {
	name := L.ToString(1)

	for _, command := range commands {
		if command.Name == name {
			command.Run(c)
			return 1
		}
	}

	return 1
}

func args(args []string, L *lua.LState) int {
	table := L.NewTable()

	for _, arg := range args {
		table.Append(lua.LString(arg))
	}

	L.Push(table)
	return 1
}
