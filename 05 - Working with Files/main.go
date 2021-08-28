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
		fmt.Println("Does the File exist : YES")
	} else {
		fmt.Println("Does the File exist : NO")
		return false
	}
	// fmt.Println(os.Stat(fileName))
	return true
}

func readFile(fileName string) {
	contentBytes, err := ioutil.ReadFile(fileName)

	fmt.Printf("\nFile contents: \n")

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

func appendToFile(fileName string) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err := f.WriteString("\nAppending text\n"); err != nil {
		log.Fatal(err)
	}
}

func main() {
	fileName := "DoIExist"
	exists := checkIfExists(fileName)
	if exists == true {
		overwriteFile(fileName)
		appendToFile(fileName)
		readFile(fileName)
	}

}
