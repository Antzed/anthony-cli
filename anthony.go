package main


import (
  "fmt"
  "log"
  "os"
  "os/exec"
  "github.com/urfave/cli/v2"
  "strings"
  "github.com/yuin/gopher-lua"
  "io/ioutil"
  "errors"
  "database/sql"
  _"github.com/mattn/go-sqlite3"
  "time"
)

func main() {
  var board string
  app := &cli.App{
    Name: "anthony",
    Usage: "anthony's linux automation",

    Flags: []cli.Flag {
      &cli.StringFlag{
        Name: "lang",
        Value: "english",
        Usage: "language for the greeting",
      },
    },
    Commands: []*cli.Command{
      {
        Name: "fishtank",
        Usage: "show a fish tank",
        Action: func(c *cli.Context) error {
            exec.Command("cursetank").Output()
            return nil
        },
      },
      {
        Name:    "initialize",
        Aliases: []string{"init"},
        Usage:   "initilize enviroment",
        Action:  func(c *cli.Context) error {
            cmd := exec.Command("npm", "install", "--global","taskbook")
            out, err := cmd.Output()
            //err := cmd.Run()
            if err != nil{
                panic(err)
            }
            fmt.Printf(string(out))
            return nil
        },
      },
	
      {
        Name:    "add",
        Aliases: []string{"a"},
        Usage:   "add things",
	    Subcommands: []*cli.Command{
            //add task --job true --type "Individual Assignment" --due "2022-4-1"    
            {
                Name: "task",
                Usage: "add a new task in taskbook",
                Flags: []cli.Flag{
                    &cli.StringFlag{
                        Name: "board",
                        Value: "My Board",
                        Usage: "add a new task to a specific board",
                        Destination: &board,
                    },
                },
                Action: func(c *cli.Context) error {
                    var input string = c.Args().First()
                    inputsplit := strings.Split(input, " and ")
                    addTask(inputsplit, board)
                    return nil
                },
            },
            {
                Name: "job",
                Usage: "add new, official jobs to the database",
                Flags: []cli.Flag{
                     &cli.StringFlag{
                         Name: "type",
                         Usage: "set job type",
                         Required: true,
                     },
                     &cli.StringFlag{
                         Name: "due",
                         Usage: "set job type",
                         Required: true,
                     },
                 },
                 Action: func(c *cli.Context) error {
                    db, err := sql.Open("sqlite3", "./job.db")
                    checkErr(err)
                    fmt.Println("opened database")
                    checkErr(err)
                    var queryJobTypeID = "SELECT JobTypeID FROM JOB_TYPE WHERE JobTypeName = '" + c.String("type") + "'"
                    rows, err := db.Query(queryJobTypeID)
                    checkErr(err)
                    var jid int

                    for rows.Next() {
                        err = rows.Scan(&jid)
                        checkErr(err)
                    }
                    stmt, err := db.Prepare("INSERT INTO JOB(JobName, JobTypeID, DueDate) values(?,?,?)")
                    checkErr(err)
                    res, err := stmt.Exec(c.Args().First(), jid, c.String("due"))
                    checkErr(err)
                    id, err := res.LastInsertId()
                    checkErr(err)
                    fmt.Println(id)
                    db.Close()
                    fmt.Println("database closed")
                    return nil
                 },
            },
            {
                Name: "project",
                Usage: "add a new project with gantt chart",
                Action: func(c *cli.Context) error {
                    path := "projects"
                	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		                err := os.Mkdir(path, os.ModePerm)
		                if err != nil {
			                log.Println(err)
		                }
	                }
                    var input string = c.Args().First()
                    L := lua.NewState()
                    defer L.Close()
                    err := L.DoFile("luaScript.lua")
                    checkErr(err)
                    if err := L.CallByParam(lua.P{
		                Fn:      L.GetGlobal("addProject"), // name of Lua function
		                NRet:    0,                     // number of returned values
		                Protect: true,                  // return err or panic
	                },  lua.LString(input)); err != nil {
		                panic(err)
	                }
                    return nil
                },
            },
        },
    },
    

      {
        Name:    "show",
        Aliases: []string{"sw"},
        Usage:   "show things",
        Subcommands: []*cli.Command{
            {
                Name: "project",
                Usage: "list all the projects",
                Action: func(c *cli.Context) error {
                    files, err := ioutil.ReadDir("projects")
                    if err != nil {
                        log.Fatal(err)
                    }
 
                    for _, f := range files {
                        fmt.Println(f.Name())
                    }            
                    return nil
                },
            },
            {
                Name: "task",
                Usage: "list all task",
                Action: func(c *cli.Context) error{
                    L := lua.NewState()
                    defer L.Close()
                    err := L.DoFile("luaScript.lua")
                    checkErr(err)
                    //err != nil{
                        //panic(err)
                    //}
                    if err := L.CallByParam(lua.P{
                        Fn:      L.GetGlobal("showTask"), // name of Lua function
                        NRet:    0,                     // number of returned values
                        Protect: true,                  // return err or panic
                    }); err != nil {
                        panic(err)
                    }
                    return nil
                },
            },
            {
                Name: "job",
                Usage: "show all the jobs",
                Action: func(c *cli.Context) error{
                    db, err := sql.Open("sqlite3", "./job.db")
                    checkErr(err)
                    fmt.Println("opened database")
                    rows, err := db.Query("SELECT j.JobID, j.JobName, jt.JobTypeName, j.DueDate  FROM JOB j JOIN JOB_TYPE jt ON j.JobTypeID = jt.JobTypeID")
                    checkErr(err)
                    var jid int
                    var jname string
                    var jtype string
                    var jduedate time.Time

                    for rows.Next() {
                        err = rows.Scan(&jid, &jname, &jtype, &jduedate)
                        checkErr(err)
                        fmt.Println(jid, jname, jtype, jduedate)
                    }
                    rows.Close() 
                    db.Close()
                    return nil
                },
            },
        },
      },
      {
        Name: "delete",
        Usage: "delete things",
        Subcommands: []*cli.Command{
            {
                Name: "project",
                Usage: "delete existing projects",
                Action: func(c *cli.Context) error{
                    var input string = "./projects/" + c.Args().First()
                    err := os.Remove(input)
                    checkErr(err)
                    fmt.Println("deleted project", c.Args().First())
                    return nil
                },
            },
        },
      },
    },

    Action: func(c *cli.Context) error {
        name := "Anthony"
        if c.NArg() > 0 {
            name = c.Args().Get(0)
        }
        if c.String("lang") == "spanish" {
            fmt.Println("Hola", name)
        } else {
            fmt.Println("Hello", name)
        }
            return nil
    },

  }
  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func addTask(inputsplit []string, board string){
    if board != "My Board" {
        var atboard string = "@" + board
        for _, s := range inputsplit {
            fmt.Println("added task: ", s, " at board: ", board)
            cmd := exec.Command("tb", "-t", atboard, s)
            err := cmd.Run()

            checkErr(err)
            fmt.Println("done")
        }
    } else {
        for _, s := range inputsplit {
            fmt.Println("added task: ", s)
            cmd := exec.Command("tb", "-t" ,s)
            err := cmd.Run()

            checkErr(err)
            fmt.Println("done1")

        }
    }
}
