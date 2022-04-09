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
  //"database/sql"
  //_"github.com/mattn/go-sqlite3"
)

func main() {
  var board string
  var isJob bool
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
                    &cli.BoolFlag{
                        Name: "job",
                        Value: false,
                        Usage: "want to urn it to job and store in databse",
                        Destination: &isJob,
                    },
                },
                Action: func(c *cli.Context) error {
                    var input string = c.Args().First()
                    inputsplit := strings.Split(input, " and ")
                    if isJob == true {
                        //db, err := sql.Open("sqlite3", "./job.db")
                        //checkErr(err)
                        fmt.Println("opened database")
                        //db.Close()
                        //fmt.Println("and then the database closed")
                    } else {
                        fmt.Println("else")
                    }
                    if board != "My Board" {
                        
                        fmt.Println("added task: ", c.Args().First(), " at board: ", board)
                        
                        var atboard string = "@" + board
                        for _, s := range inputsplit {
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
                    //db.Close()
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
