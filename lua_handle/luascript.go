package lua_handle

var LuaScript = `
  function addProject(name)
    os.execute("python3 ./GanTTY/main.py ./projects/" .. name)
  end

  function showTask()
      os.execute("tb")
  end

  function fishTank()
     os.execute("cursetank")
  end

  function exportJob()
     mode = ".mode list"
     instruction = "SELECT j.JobID, j.JobName, jt.           JobTypeName,   j.DueDate   FROM JOB j JOIN JOB_TYPE jt ON j.JobTypeID = jt.JobTypeID;"
     os.execute("sqlite3 job.db '.mode list' 'SELECT j.JobID, j.JobName, jt. JobTypeName, j.DueDate FROM JOB j JOIN JOB_TYPE jt ON j.JobTypeID = jt.JobTypeID;' > job.txt")
  end
  `


