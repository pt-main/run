package lua

import (
	"os"
	"path/filepath"

	"github.com/bmatcuk/doublestar/v4"
	lua "github.com/yuin/gopher-lua"
)

func NewTalLuaState(changedFiles, args []string) *lua.LState {
	L := lua.NewState()
	L.SetGlobal("changed", L.NewFunction(func(L *lua.LState) int {
		relPaths := makeRelativePaths(changedFiles)
		tbl := L.NewTable()
		for i, p := range relPaths {
			tbl.RawSetInt(i+1, lua.LString(p))
		}
		L.Push(tbl)
		return 1
	}))

	L.SetGlobal("get_args", L.NewFunction(func(L *lua.LState) int {
		relPaths := makeRelativePaths(args)
		tbl := L.NewTable()
		for i, p := range relPaths {
			tbl.RawSetInt(i+1, lua.LString(p))
		}
		L.Push(tbl)
		return 1
	}))

	L.SetGlobal("match_pattern", L.NewFunction(func(L *lua.LState) int {
		pat := L.CheckString(1)
		str := L.CheckString(2)
		ok, err := doublestar.Match(pat, str)
		if err != nil {
			L.Push(lua.LFalse)
			return 1
		}
		L.Push(lua.LBool(ok))
		return 1
	}))
	return L
}

func makeRelativePaths(absPaths []string) []string {
	cwd, err := os.Getwd()
	if err != nil {
		return absPaths
	}
	rel := make([]string, 0, len(absPaths))
	for _, p := range absPaths {
		relPath, err := filepath.Rel(cwd, p)
		if err == nil {
			relPath = filepath.ToSlash(relPath)
			rel = append(rel, relPath)
		} else {
			rel = append(rel, p)
		}
	}
	return rel
}
