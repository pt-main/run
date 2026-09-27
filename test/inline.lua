print("LUA: inline script body")
local args = get_args()
for i, a in ipairs(args) do
  print("LUA: arg " .. i .. " = " .. tostring(a))
end
print("LUA: done")
