-- === CONFIGURATION ===
local script_file = script_path("go-build_go-build.py")
local args = get_args()
-- =====================

local function get_python()
    local function check(cmd)
        local f = io.popen(cmd .. " --version 2>&1")
        if f then
            local out = f:read("*a")
            f:close()
            return out:match("Python") ~= nil
        end
        return false
    end
    if check("python3") then return "python3" end
    if check("python")   then return "python"  end
    return nil
end

local function escape(arg)
    if arg:match("[ \t\"']") then
        return '"' .. arg:gsub('"', '\\"') .. '"'
    end
    return arg
end

local python = get_python()
if not python then
    io.stderr:write("Error: Python interpreter not found\n")
    os.exit(1)
end

local cmd = python .. " " .. escape(script_file)
for _, a in ipairs(args) do
    cmd = cmd .. " " .. escape(a)
end

local result = os.execute(cmd)
os.exit(result or 0)