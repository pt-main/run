package generation

func GetRuntime() string {
	return `---@type string[]
local changed_list = changed()

---@class Task
---@field deps string[]
---@field func fun()

---@class Tasker
---@field tasks table<string, Task>

local tasker = {
    tasks = {}
}

---@param deps string[]
---@param changed string[]
---@return boolean
function tasker.has_any_dep_changed(patterns, changed)
    for _, pat in ipairs(patterns) do
        for _, ch in ipairs(changed) do
            if match_pattern(pat, ch) then
                return true
            end
        end
    end
    return false
end

---@param deps string[]
---@param name string
---@param func fun()
function tasker.add(deps, name, func)
    tasker.tasks[name] = {
        deps = deps,
        func = func
    }
end

---@param name string
function tasker.run(name)
    if tasker.tasks[name] == nil then
        error("Has not task: " .. name)
        return
    end
    local task = tasker.tasks[name]
    if #task.deps == 0 then
        task.func()
        return
    end
    if tasker.has_any_dep_changed(task.deps, changed_list) then
        task.func()
    end
end

-- for simple call in scripts
function run(name) 
    tasker.run(name)
end`
}
