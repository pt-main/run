-- === CONFIGURATION ===
-- FALLBACK template: runs the original file directly with its shebang.
local script_file = script_path("calc_.__calc.py")
local args = get_args()
-- =====================

local function escape(arg)
    if arg:match("[ \t\"']") then
        return '"' .. arg:gsub('"', '\\"') .. '"'
    end
    return arg
end

local cmd = "sh " .. escape(script_file)
for _, a in ipairs(args) do
    cmd = cmd .. " " .. escape(a)
end

local result = os.execute(cmd)
os.exit(result or 0)
