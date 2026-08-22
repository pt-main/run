-- @!
print("starting")

-- @build
-- #depends *.go
print("test")

-- @
print("runing")
if #get_args() > 0 then
    run(get_args()[1]) 
end
print("end")