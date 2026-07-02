math.randomseed(os.time())

request = function()
    local r = math.random()
    local delay = 0
    
    if r < 0.8 then
        delay = math.random(10, 50)
    else
        delay = math.random(500, 1500)
    end
    
    local path = "/work?delay=" .. delay .. "ms"
    return wrk.format("GET", path)
end
