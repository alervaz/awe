# Welcome to AWE

An awesome cli maker made in golang to the world

```
go get -u github.com/alervaz/awe
```
Start by creating a cmd/rootCmd.go file
```go
package cmd

import "github.com/alervaz/awe"

var App = &awe.App{}

var rootCmd = &awe.Command{
	Short:       "Hello",
	Description: `Hi and hello`,
}

func init() {
	App.Root = rootCmd
}

```

Then pass it into the InitializeCli
```go
func main() {
    awe.InitializeCli(cmd.App)
}
```

Define a command
```go
command := &Command{
    Short: "My short description",
    Description: "A long description bla bla...",
    Run: func(c *awe.Command) {
        //This will run when you go toa route
    },
    Flags: awe.Flags{
        "name": "default value",
        //Or it can be nil
        "nil": nil,
        //Can be bool, int, float32, float64, etc
    },
}
```

Then add the Command  to the list
```go
command.Add(App)
```


