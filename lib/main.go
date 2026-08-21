package runlib

import (
	"fmt"
	"log"
	"path/filepath"
	"sync"
	"sync/atomic"

	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/tycl"
	"github.com/pt-main/tycl/format"
	"github.com/pt-main/tycl/shared"
	"github.com/pt-main/tycl/utils"
	lua "github.com/yuin/gopher-lua"
)

func GetCfg() (*shared.Config, error) {
	file, err := utils.OpenF(ConfigDirConfigPath())
	if err != nil {
		return nil, err
	}
	var errI core.ErrorInterface
	cfg, errI := tycl.Process(file, tyclContract, true)
	if errI != nil {
		return cfg, fmt.Errorf(format.FormatError(errI))
	}
	return cfg, nil
}

func NewLuaState(args []string) *lua.LState {
	L := lua.NewState()

	L.SetGlobal("script_path", L.NewFunction(func(L *lua.LState) int {
		name := L.CheckString(1)
		path := filepath.Join(ConfigDirBasePath(), name)
		L.Push(lua.LString(path))
		return 1
	}))

	L.SetGlobal("get_arg", L.NewFunction(func(L *lua.LState) int {
		idx := L.CheckInt(1)
		if idx < 1 || idx > len(args) {
			L.Push(lua.LNil)
		} else {
			L.Push(lua.LString(args[idx-1]))
		}
		return 1
	}))

	L.SetGlobal("get_args", L.NewFunction(func(L *lua.LState) int {
		tbl := L.NewTable()
		for i, arg := range args {
			tbl.RawSetInt(i+1, lua.LString(arg))
		}
		L.Push(tbl)
		return 1
	}))

	L.SetGlobal("run_script", L.NewFunction(func(L *lua.LState) int {
		name := L.CheckString(1)
		var scriptArgs []string
		top := L.GetTop()
		for i := 2; i <= top; i++ {
			arg := L.Get(i)
			if str, ok := arg.(lua.LString); ok {
				scriptArgs = append(scriptArgs, string(str))
			} else {
				scriptArgs = append(scriptArgs, L.ToStringMeta(arg).String())
			}
		}

		cfg, err := GetCfg()
		if err != nil {
			L.RaiseError("failed to load config: %v", err)
			return 0
		}

		if err := RunScript(cfg, name, scriptArgs); err != nil {
			L.RaiseError("failed to run script %q: %v", name, err)
			return 0
		}
		return 0
	}))

	var (
		activeScripts int32
		wg            sync.WaitGroup
	)

	L.SetGlobal("run_script_parallel", L.NewFunction(func(L *lua.LState) int {
		name := L.CheckString(1)
		var scriptArgs []string
		top := L.GetTop()
		for i := 2; i <= top; i++ {
			arg := L.Get(i)
			if str, ok := arg.(lua.LString); ok {
				scriptArgs = append(scriptArgs, string(str))
			} else {
				scriptArgs = append(scriptArgs, L.ToStringMeta(arg).String())
			}
		}

		cfg, err := GetCfg()
		if err != nil {
			L.RaiseError("failed to load config: %v", err)
			return 0
		}

		atomic.AddInt32(&activeScripts, 1)
		wg.Add(1)

		go func() {
			defer wg.Done()
			defer atomic.AddInt32(&activeScripts, -1)
			defer func() {
				if r := recover(); r != nil {
					log.Printf("script %q panicked: %v", name, r)
				}
			}()

			if err := RunScript(cfg, name, scriptArgs); err != nil {
				log.Printf("script %q failed: %v", name, err)
			}
		}()
		return 0
	}))

	L.SetGlobal("wait", L.NewFunction(func(L *lua.LState) int {
		wg.Wait()
		return 0
	}))

	return L
}
