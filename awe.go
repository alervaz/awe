package awe

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Command struct {
	Cmd         string
	Short       string
	Description string
	Flags       Flags
	Run         func(*Command)
}

type App struct {
	Root     *Command
	Commands []Command
}

type Flags map[string]any

func (c Command) Add(app *App) {
	app.Commands = append(app.Commands, c)
}

func (c *Command) Execute() {
	if c.Run != nil {
		c.Run(c)
	}
}

func InitializeCli(app *App) {
	args := os.Args
	if len(args) == 1 || strings.HasPrefix(args[1], "--") {
		if len(args) > 1 {
			for i, arg := range args {
				if !strings.HasPrefix(arg, "--") {
					continue
				}
				trimmedArg := strings.Trim(arg, "--")
				_, ok := app.Root.Flags[trimmedArg]

				if ok {
					if i+2 > len(args) {
						log.Println("Flag with value is not provided")
						continue
					}
					value := args[i+1]
					app.Root.Flags[trimmedArg] = value
				}
			}
		}
		if app.Root != nil {
			fmt.Println(app.Root.Short)
			fmt.Println(app.Root.Description)
			fmt.Println("")

			fmt.Println("Flags: ")
			for flag := range app.Root.Flags {
				fmt.Println(" --" + flag)
			}
			fmt.Println(" -h")
			fmt.Println("")

		}
		if len(app.Commands) > 0 {
			fmt.Println("Commands: Cmds")
		}
		for _, cmd := range app.Commands {
			if cmd.Cmd == "" {
				continue
			}
			fmt.Printf("  %s: %s\n", cmd.Cmd, cmd.Short)
			fmt.Println("")
		}
		app.Root.Execute()
		return
	}
	exists := false
	for _, cmd := range app.Commands {
		if cmd.Cmd != args[1] {
			continue
		}
		exists = true
		isHelp := false
		for i, arg := range args {
			if arg == "-h" {
				isHelp = true
				fmt.Println(cmd.Short)
				fmt.Println(cmd.Description)
				fmt.Println("")
				fmt.Println("Flags:")
				if len(app.Root.Flags) != 0 {
					for flag := range cmd.Flags {
						fmt.Println(" --" + flag)
					}
				}
				fmt.Println(" -h")
			}

			if !strings.HasPrefix(arg, "--") {
				continue
			}
			trimmedArg := strings.Trim(arg, "--")
			_, ok := cmd.Flags[trimmedArg]

			if ok {
				if i+2 > len(args) {
					log.Println("Flag with value is not provided")
					continue
				}
				value := args[i+1]
				cmd.Flags[trimmedArg] = value
			}
		}

		if !isHelp {
			cmd.Execute()
		}
	}

	if !exists {
		log.Println("Endpoint does not exist")
	}
}
