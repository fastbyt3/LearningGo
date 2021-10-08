package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	// "path/filepath"
)

const usage = `
Proper Usage:

todo [option] [args]

OPTIONS:
	a [task]	-- add a new task
	ls 			-- list all tasks
	curr [sno]	-- change task to currently doing status
	done [sno]	-- change task to finished status
`
const fileLoc = "C:\\Users\\aldee\\.todo"

// const fileLoc = "/home/fastbyte/.todo"

func ChkToDoFile() {
	if _, err := os.Stat(fileLoc); os.IsNotExist(err) {
		// TODO : create a file - .todo in home dir
		if _, e := os.Create(fileLoc); e != nil {
			log.Fatal(e)
		}
		fmt.Println("Created a file ✔")
	}
}

func addTask(task string) {

}

func main() {

	// Print "Usage" if no arg is given
	allArgs := os.Args[0:]
	if len(allArgs) == 1 {
		fmt.Println(usage)
		return
	}

	ChkToDoFile()

	option := allArgs[1]

	if option == "a" {
		// addTask()
	} else if option == "ls" {
		// listTask()
	} else if strings.ContainsAny(option, "curr | done") {
		// changeStatus(option)
	} else {
		fmt.Println("[X] Invalid argument supplied!!")
		fmt.Println(usage)
		return
	}
}
