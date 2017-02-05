package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/codegangsta/cli"
)

type config struct {
	Package     string
	Name        string
	UpperName   string
	FirstLetter string
	Example     string
	dir         string
}

func Generate() cli.Command {
	return cli.Command{
		Name:  "generate",
		Usage: "Generate new Gobot adaptors, drivers, and projects",
		Action: func(c *cli.Context) {
			valid := false
			for _, s := range []string{"adaptor", "driver", "project"} {
				if s == c.Args().First() {
					valid = true
				}
			}
			if valid == false {
				fmt.Println("Invalid/no subcommand supplied.\n")
				fmt.Println("Usage:")
				fmt.Println(" gobot generate adaptor <name> [package] # generate a new Gobot adaptor")
				fmt.Println(" gobot generate driver  <name> [package] # generate a new Gobot driver")
				fmt.Println(" gobot generate project <name> [package] # generate a new Gobot project")
				return
			}

			if len(c.Args()) < 2 {
				fmt.Println("Please provide a one word name.")
				return
			}

			name := strings.ToLower(c.Args()[1])
			packageName := name
			if len(c.Args()) > 2 {
				packageName = strings.ToLower(c.Args()[2])
			}
			upperName := strings.ToUpper(string(name[0])) + string(name[1:])

			cfg := config{
				Package:     packageName,
				UpperName:   upperName,
				Name:        name,
				FirstLetter: string(name[0]),
				dir:         ".",
			}

			switch c.Args().First() {
			case "adaptor":
				if err := generateAdaptor(cfg); err != nil {
					fmt.Println(err)
				}
			case "driver":
				if err := generateDriver(cfg); err != nil {
					fmt.Println(err)
				}
			case "project":
				pwd, err := os.Getwd()
				if err != nil {
					fmt.Println(err)
					return
				}
				dir := pwd + "/gobot-" + cfg.Name
				fmt.Println("Creating", dir)
				if err := os.MkdirAll(dir, 0700); err != nil {
					fmt.Println(err)
					return
				}
				cfg.dir = dir

				examplesDir := dir + "/examples"
				fmt.Println("Creating", examplesDir)
				if err := os.MkdirAll(examplesDir, 0700); err != nil {
					fmt.Println(err)
					return
				}

				if err := generateProject(cfg); err != nil {
					fmt.Println(err)
				}
			}
		},
	}
}

func generate(c config, file string, tmpl string) error {
	fileLocation := c.dir + "/" + file
	fmt.Println("Creating", fileLocation)

	f, err := os.Create(fileLocation)
	defer f.Close()
	if err != nil {
		return err
	}

	t, err := template.New("").Parse(tmpl)
	if err != nil {
		return err
	}

	return t.Execute(f, c)
}

func generateDriver(c config) error {
	if err := generate(c, c.Name+"_driver.go", driver()); err != nil {
		return err
	}

	return generate(c, c.Name+"_driver_test.go", driverTest())
}

func generateAdaptor(c config) error {
	if err := generate(c, c.Name+"_adaptor.go", adaptor()); err != nil {
		return err
	}

	return generate(c, c.Name+"_adaptor_test.go", adaptorTest())
}

func generateProject(c config) error {
	if err := generateDriver(c); err != nil {
		return err
	}
	if err := generateAdaptor(c); err != nil {
		return err
	}

	dir := c.dir
	exampleDir := dir + "/examples"
	c.dir = exampleDir

	if err := generate(c, "main.go", example()); err != nil {
		return err
	}

	c.dir = dir

	if ex, err := ioutil.ReadFile(exampleDir + "/main.go"); err != nil {
		return err
	} else {
		c.Example = string(ex)
	}

	return generate(c, "README.md", readme())
}

func adaptor() string {
	return `package {{.Package}}

import (
	"github.com/hybridgroup/gobot"
)

var _ gobot.Adaptor = (*{{.UpperName}}Adaptor)(nil)

type {{.UpperName}}Adaptor struct {
	name string
}

func New{{.UpperName}}Adaptor(name string) *{{.UpperName}}Adaptor {
	return &{{.UpperName}}Adaptor{
		name: name,
	}
}

func ({{.FirstLetter}} *{{.UpperName}}Adaptor) Name() string { return {{.FirstLetter}}.name }

func ({{.FirstLetter}} *{{.UpperName}}Adaptor) Connect() []error { return nil }

func ({{.FirstLetter}} *{{.UpperName}}Adaptor) Finalize() []error { return nil }

func ({{.FirstLetter}} *{{.UpperName}}Adaptor) Ping() string { return "pong" }
`
}

func driver() string {
	return `package {{.Package}}

import (
	"time"

	"github.com/hybridgroup/gobot"
)

var _ gobot.Driver = (*{{.UpperName}}Driver)(nil)

const Hello string = "hello"

type {{.UpperName}}Driver struct {
	name string
	connection gobot.Connection
	interval time.Duration
	halt chan bool
	gobot.Eventer
	gobot.Commander
}

func New{{.UpperName}}Driver(a *{{.UpperName}}Adaptor, name string) *{{.UpperName}}Driver {
	{{.FirstLetter}} := &{{.UpperName}}Driver{
		name: name,
		connection: a,
		interval: 500*time.Millisecond,
		halt: make(chan bool, 0),
    Eventer:    gobot.NewEventer(),
    Commander:  gobot.NewCommander(),
	}

	{{.FirstLetter}}.AddEvent(Hello)

	{{.FirstLetter}}.AddCommand(Hello, func(params map[string]interface{}) interface{} {
		return {{.FirstLetter}}.Hello()
	})

	return {{.FirstLetter}}
}

func ({{.FirstLetter}} *{{.UpperName}}Driver) Name() string { return {{.FirstLetter}}.name }

func ({{.FirstLetter}} *{{.UpperName}}Driver) Connection() gobot.Connection {
	return {{.FirstLetter}}.connection
}

func ({{.FirstLetter}} *{{.UpperName}}Driver) adaptor() *{{.UpperName}}Adaptor {
	return {{.FirstLetter}}.Connection().(*{{.UpperName}}Adaptor)
}

func ({{.FirstLetter}} *{{.UpperName}}Driver) Hello() string {
	return "hello from " + {{.FirstLetter}}.Name() + "!"
}

func ({{.FirstLetter}} *{{.UpperName}}Driver) Ping() string {
	return {{.FirstLetter}}.adaptor().Ping()
}

func ({{.FirstLetter}} *{{.UpperName}}Driver) Start() []error {
	go func() {
		for {
			gobot.Publish({{.FirstLetter}}.Event(Hello), {{.FirstLetter}}.Hello())

			select {
			case <- time.After({{.FirstLetter}}.interval):
			case <- {{.FirstLetter}}.halt:
				return
			}
		}
	}()
	return nil
}

func ({{.FirstLetter}} *{{.UpperName}}Driver) Halt() []error {
	{{.FirstLetter}}.halt <- true
	return nil
}

`
}

func example() string {
	return `
package main

import (
  "../"
  "fmt"
  "time"

  "github.com/hybridgroup/gobot"
)

func main() {
  gbot := gobot.NewGobot()

  conn := {{.Package}}.New{{.UpperName}}Adaptor("conn")
  dev := {{.Package}}.New{{.UpperName}}Driver(conn, "dev")

  work := func() {
    gobot.On(dev.Event({{.Package}}.Hello), func(data interface{}) {
      fmt.Println(data)
    })

    gobot.Every(1200*time.Millisecond, func() {
      fmt.Println(dev.Ping())
    })
  }

  robot := gobot.NewRobot(
    "robot",
    []gobot.Connection{conn},
    []gobot.Device{dev},
    work,
  )

  gbot.AddRobot(robot)
  gbot.Start()
}
`
}

func driverTest() string {
	return `package {{.Package}}

import (
	"testing"
	"time"

	"github.com/hybridgroup/gobot"
)

func Test{{.UpperName}}Driver(t *testing.T) {
	d := New{{.UpperName}}Driver(New{{.UpperName}}Adaptor("conn"), "dev")

	gobot.Assert(t, d.Name(), "dev")
	gobot.Assert(t, d.Connection().Name(), "conn")

	ret := d.Command(Hello)(nil)
	gobot.Assert(t, ret.(string), "hello from dev!")

	gobot.Assert(t, d.Ping(), "pong")

	gobot.Assert(t, len(d.Start()), 0)

	<-time.After(d.interval)

	sem := make(chan bool, 0)

	gobot.On(d.Event(Hello), func(data interface{}) {
		sem <- true
	})

	select {
	case <-sem:
	case <-time.After(600 * time.Millisecond):
		t.Errorf("Hello Event was not published")
	}

	gobot.Assert(t, len(d.Halt()), 0)

	gobot.On(d.Event(Hello), func(data interface{}) {
		sem <- true
	})

	select {
	case <-sem:
		t.Errorf("Hello Event should not publish after Halt")
	case <-time.After(600 * time.Millisecond):
	}
}

`
}

func adaptorTest() string {
	return `package {{.Package}}

import (
	"testing"

	"github.com/hybridgroup/gobot"
)

func Test{{.UpperName}}Adaptor(t *testing.T) {
	a := New{{.UpperName}}Adaptor("tester")

	gobot.Assert(t, a.Name(), "tester")

	gobot.Assert(t, len(a.Connect()), 0)

	gobot.Assert(t, a.Ping(), "pong")

	gobot.Assert(t, len(a.Connect()), 0)

	gobot.Assert(t, len(a.Finalize()), 0)
}
`
}

func readme() string {
	return `# {{.Package}}

Gobot (http://gobot.io/) is a framework for robotics and physical computing using Go

This repository contains the Gobot adaptor and driver for {{.Package}}.

For more information about Gobot, check out the github repo at
https://github.com/hybridgroup/gobot

## Installing
` + "```bash\ngo get path/to/repo/{{.Package}}\n```" + `

## Using
` + "```go{{.Example}}\n```" + `

## Connecting

Explain how to connect to the device here...

## License

Copyright (c) 2015 <Your Name Here>. Licensed under the <Insert license here> license.
`
}
