package lua_handle

import (
    "github.com/Antzed/anthony-cli/path_handle"
)

var projectPath = path_handle.RootDir() + "/"
var jobdbPath = projectPath + "job.db"
var LuaScript = `
  function addProject(name)
    os.execute("python3 `+ projectPath + `GanTTY/main.py `+ projectPath + `projects/" .. name)
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
     os.execute("sqlite3 `+ jobdbPath +` '.mode list' 'SELECT j.JobID, j.JobName, jt. JobTypeName, j.DueDate FROM JOB j JOIN JOB_TYPE jt ON j.JobTypeID = jt.JobTypeID;' > `+ projectPath + `job.txt")
  end

  function exportJobThisWeek(timerange)
    instruction = "SELECT j.JobName, jt.JobTypeName,   j.DueDate   FROM JOB j JOIN JOB_TYPE jt ON j.JobTypeID = jt.JobTypeID WHERE DueDate BETWEEN ".. timerange ..";"
    os.execute("sqlite3 `+ jobdbPath + ` \".mode list\" \"".. instruction .. "\" > ` + projectPath +`job.txt")

  end
  `


