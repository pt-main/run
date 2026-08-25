-- @build
-- #depends main.go
print_colored("[?RD]building[?RT]\n")
shell("echo 'End of building'")

-- @runTest
run("help")

-- @
if #get_args() > 0 then
    script(get_args()[1]) 
end