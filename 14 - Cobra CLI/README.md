# Using Cobra to build Golang CLI apps

## Installation

```bash
go get -u github.com/spf13/cobra/cobra
```

To use in go app

```go
import "github.com/spf13/cobra"
```

## Usage

- Cobra is built on a structure of commands, arguments & flags.
- Commands represent actions, Args are things and Flags are modifiers for those actions.

- Pattern: `APPNAME VERB NOUN --ADJECTIVE` or `APPNAME COMMAND ARG --FLAG`

Example: `hugo server --port=8080` -> Command is `server` and `port` is the flag

- General file structure
    
    ```
    AppName/
    |
    |_ cmd/
    |   |
    |   |_ cmds_go_here.go
    |   |_ ...
    |
    |_ main.go
    ```


## Building a weather api

### Initializing project

- we can use `cobra-cli` a tool to set up the workspace

```bash
go install github.com/spf13/cobra-cli@latest

go mod init learn/cobra/weather

cobra-cli init weather_app -a fastbyt3
```

- Now we have the project structure set up automatically


### Adding commands

```bash
cobra-cli add fetch
```

`fetch.go` is created inside the `cmd` dir 

- By default we have a variable `fetchCmd` that executes a function when fetch is passed in CLI
- In some cases like here we might need some flags to do some operation. Flags can be passed in `init()` method.

- To create flags:

```
fetchCmd.PersistentFlags().StringVarP(
    VAR_TO_STORE,
    EXPANDED_FORM,
    SHORT_FORM,
    nil string,
    DESCRIPTION
)
```

- To make a flag REQUIRED, `rootCmd.MarkFlagRequired("full_form_arg")`

- Cobra also has in-built validators of args like min args, max, etc. Refer [https://cobra.dev/#positional-and-custom-arguments](Cobra docs on arg validators)

- It is possible to set a cmd to be a child cmd using cobra-cli

    ```
    cobra-cli create child -p parentCmd
    ```
    
    ```bash
    weather parent child
    ```

## Configuring generator

modify `~/.cobra.yaml`

---
