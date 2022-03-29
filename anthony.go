package main


import (
  "fmt"
  "log"
  "os"
  "os/exec"
  "github.com/urfave/cli/v2"
)

func main() {
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
                Action: func(c *cli.Context) error {
                    fmt.Println("added task: ", c.Args().First())
                    cmd := exec.Command("tb -c " + c.Args().First())
                    stdout, err := cmd.Output()
                    if err != nil {
                        fmt.Println(err.Error())
                    }
                    // Print the output
                    fmt.Println(string(stdout))
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
