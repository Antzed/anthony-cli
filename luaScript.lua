function addProject(name)
  os.execute("python3 ./GanTTY/main.py ./projects/" .. name)
end

function showTask()
    os.execute("tb")
end

