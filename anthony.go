package main


import (
  "fmt"
  "log"
  "os"
  "os/exec"
  "io/ioutil"
  _"embed"
  er "github.com/Antzed/anthony-cli/error_handle"
  th "github.com/Antzed/anthony-cli/task_handle"
  "github.com/Antzed/anthony-cli/db_handle"
  "github.com/Antzed/anthony-cli/lua_handle"
  "github.com/Antzed/anthony-cli/file_handle"
  "github.com/Antzed/anthony-cli/path_handle"
  "github.com/urfave/cli/v2"
  "strings"
  "errors"
  _"github.com/mattn/go-sqlite3"
  "time"
  "github.com/qeesung/image2ascii/convert"
  _ "image/jpeg"
  _ "image/png"
  "image"
  "bytes"
  _"github.com/faiface/beep"
  "github.com/faiface/beep/mp3"
  "github.com/faiface/beep/speaker"
  "strconv"
)


//go:embed aileen.jpg
var loveimage []byte



//var luaScript = `
 //function addProject(name)
  // os.execute("python3 ./GanTTY/main.py ./projects/" .. name)
 //end

 //function showTask()
   //  os.execute("tb")
 //end

 //function fishTank()
   // os.execute("cursetank")
 //end
//`



func main() {
  var board string
  projectPath  := path_handle.RootDir() + "/"
  app := &cli.App{
    Name: "anthony",
    Usage: "anthony's linux automation",
    Version: "0.2.0",
    EnableBashCompletion:true,
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
            //add task --job true --added this weeks job    
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
                        Usage: "confirm whether add task from job",
                    },
                    &cli.StringFlag{
                          Name: "due",
                          Value: "null",
                          Usage: "specify date of the job before pulling",
                    },
                },
                Action: func(c *cli.Context) error {
                    
                    //fmt.Println(c.String("due"))
                    //if --job true
                    if c.Bool("job") == true {
                        db := db_handle.OpenDB("job", projectPath)
                        db_handle.ExportJobFromWeek(c.String("due"))
                        db.Close()
                        fmt.Println("db closed")

                        jobtxtPath := projectPath + "job.txt"
                        file := file_handle.OpenFile(jobtxtPath)
                        defer file.Close()

                        var jobs []string
                        jobs = file_handle.ScansToList(file, jobs)
                        fmt.Println(jobs)
                        
                        
                        th.AddTask(jobs, board)

                    } else {
                        var input string = c.Args().First()
                        inputsplit := strings.Split(input, " and ")
                        th.AddTask(inputsplit, board)
                    }


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
                    //open a database instance
                    db := db_handle.OpenDB("job", projectPath)                   
                    //select foriegn key(return a id)
                    var fkSelectCondition = "JobTypeName = '"+c.String("type")+"'"
                    var jtid = db_handle.SelectForeignKey(db, "JobTypeID", "JOB_TYPE", fkSelectCondition)

                    //insert a row in table job using the id
                    insertInstruction := "INSERTINTO JOB(JobName, JobTypeID, DueDate) values('" +c.Args().First() + "', " + strconv.Itoa(jtid) + ", '" + c.String("due") + "')"
                    fmt.Println(insertInstruction)
                    db_handle.Insert(db, insertInstruction)
    
                    
                    //close database
                    db_handle.CloseDB(db)
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
                    L := lua_handle.InitScriptString(lua_handle.LuaScript)
                    defer L.Close()                    

                    lua_handle.CallOneStrParamFunc(L, "addProject", 0, input)
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
                    path:= projectPath + "projects"
                    files, err := ioutil.ReadDir(path)
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
                    L := lua_handle.InitScriptString(lua_handle.LuaScript)
                    defer L.Close()
                    lua_handle.CallNoParamFunc(L, "showTask", 0)
                    return nil
                },
            },
            {
                Name: "job",
                Usage: "show all the jobs",
                Flags: []cli.Flag{
                      &cli.StringFlag{
                          Name: "due",
                          Usage: "set job type",
                          Value: "null",
                      },
                },
                Action: func(c *cli.Context) error{
                    db := db_handle.OpenDB("job", projectPath)
                    if c.String("due") == "null" {
                        db_handle.ShowJob(db)
                    } else{
                        db_handle.ShowJobByDueday(db, c.String("due"))
                    }
                    
                    
                    //if c.Bool("job") == true {
                      //   db := db_handle.OpenDB("job", "./")
                         //db_handle.ExportJobFromWeek("null")
                    //}
                    db_handle.CloseDB(db)
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
                    er.CheckErr(err)
                    fmt.Println("deleted project", c.Args().First())
                    return nil
                },
            },
            {
                Name: "job",
                Usage: "delete existing jobs",
                Flags: []cli.Flag{
                    &cli.StringFlag{
                        Name: "name",
                        Usage: "specify job names",
                        Required: true,
                    },
                    &cli.StringFlag{
                        Name: "due",
                        Usage: "specify job due date",
                        Required: true,
                    },
                 },
                Action: func(c *cli.Context) error{
                    db := db_handle.OpenDB("job", projectPath)
                    db_handle.DeleteJob(db, c.String("name"), c.String("due"))
                    fmt.Println("deleted line")
                    db.Close()
                    fmt.Println("cloased database")
                    return nil
                },

            },
        },
      },
      {
        Name: "update",
        Usage: "update values",
        Subcommands: []*cli.Command{
            {
                Name: "job",
                Usage: "update job's column",
                Flags: []cli.Flag{
                    &cli.IntFlag{
                        Name: "target",
                        Usage: "specify target row id",
                        Required: true,
                    },
                    &cli.StringFlag{
                        Name: "column",
                        Usage: "specifiy target column",
                        Required: true,
                    },
                    &cli.StringFlag{
                        Name: "to",
                        Usage: "specify the value that is going to be changed to",
                        Required: true,
                    },
                },
                Action: func(c *cli.Context) error{
                    db := db_handle.OpenDB("job", projectPath)
                    db_handle.UpdateJobColumn(db, c.Int("target"), c.String("column"), c.String("to"))
                    db.Close()
                    fmt.Println("database closed")
                    return nil
                },

            },
        },
      },
      {
        Name: "love",
        Usage: "to love",
        Action: func(c *cli.Context) error{
            convertOptions := convert.DefaultOptions
            convertOptions.Ratio = 0.25
            converter := convert.NewImageConverter()
            img, _, err := image.Decode(bytes.NewReader(loveimage))
            if err != nil {
                log.Fatalln(err)
            }
            //var imagename = c.Args().First() + ".jpg"
            fmt.Print(converter.Image2ASCIIMatrix(img, &convertOptions))
            return nil
        },
      },
      {
          Name: "poopshart",
          Usage: "poopshart",
          Action: func(c *cli.Context) error{
             convertOptions := convert.DefaultOptions
             convertOptions.Ratio = 0.25
             converter := convert.NewImageConverter()
             //img, _, err := image.Decode(bytes.NewReader(loveimage))
             //if err != nil {
               //  log.Fatalln(err)
             //}
             fmt.Print(converter.ImageFile2ASCIIString("poopshart.png", &convertOptions))
             f, err := os.Open("20 second raunchy fart (slowed + reverb)-AEIqCtImdsI.mp3")
	         if err != nil {
                 log.Fatal(err)
             }
             streamer, format, err := mp3.Decode(f)
	        if err != nil {
		        log.Fatal(err)
	        }
            defer streamer.Close()
            speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
            speaker.Play(streamer)
            select {}
            return nil
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
            fmt.Println("projectPath is", projectPath)
        } 
        return nil
    },

  }
  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
