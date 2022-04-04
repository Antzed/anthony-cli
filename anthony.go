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
        Name:    "complete",
        Aliases: []string{"c"},
        Usage:   "complete a task on the list",
        Action:  func(c *cli.Context) error {
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
                },
                Action: func(c *cli.Context) error {
                    var input string = c.Args().First()
                    inputsplit := strings.Split(input, " and ")
                    if board != "My Board" {
                        
                        fmt.Println("added task: ", c.Args().First(), " at board: ", board)
                        
                        var atboard string = "@" + board
                        for _, s := range inputsplit {
                            cmd := exec.Command("tb", "-t", atboard, s)
                            err := cmd.Run()

                            if err != nil {
                                log.Fatal(err)
                            }
                            fmt.Println("done")
                        }
                    } else {
                        for _, s := range inputsplit {
                            fmt.Println("added task: ", s) 
                            cmd := exec.Command("tb", "-t" ,s)
                            err := cmd.Run()

                            if err != nil {
                                log.Fatal(err)
                            }
                            fmt.Println("done1")
                            
                        }
                    }
                    return nil
                },
            },
            {
                Name: "project",
                Usage: "add a new project with gantt chart",
                Action: func(c *cli.Context) error {
                    var input string = c.Args().First()
                    L := lua.NewState()
                    defer L.Close()
                    if err := L.DoFile("luaScript.lua"); err != nil{
                        panic(err)
                    }
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
