-- @test
print("test")
run_cli("--lm -r test")
run_cli("tal run main.task.lua test2")

-- @test2
print("test2")
script("test3")
script("test4")

-- @test3
-- #depends main.go
print("1!")

-- @test4
-- #depends merged.md
print("2!")

-- @
if #get_args() > 0 then
    script(get_args()[1]) 
end

update()