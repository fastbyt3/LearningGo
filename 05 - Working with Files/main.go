package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// func <name>(<params>) <ret type> {}
func checkIfExists(fileName string) bool {
	if _, err := os.Stat(fileName); err == nil {
		fmt.Println("File exists!!")
	} else {
		fmt.Println("File doesn't exist")
		return false
	}
	// fmt.Println(os.Stat(fileName))
	return true
}

func readFile(fileName string) {
	contentBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}

	content := string(contentBytes)
	fmt.Println(content)
}

func overwriteFile(fileName string) {
	if err := os.WriteFile(fileName, []byte("This overwrites the entire file"), 0666); err != nil {
		log.Fatal(err)
	}
}

func main() {
	fileName := "DoIExist"
	exists := checkIfExists(fileName)
	if exists == true {
		overwriteFile(fileName)
		readFile(fileName)
	}

}
