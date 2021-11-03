package main

import (
	"log"
	"os"
)

func basic() {
	log.Println("First log statement!!")
	// OP: 2021/11/03 10:44:11 First log statement!!
}

func WriteToFile() {
	file, err := os.OpenFile("./logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.Println("Writing to file!")
}

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func InitCustomLogger() {
	file, err := os.OpenFile("./logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func customLogger() {
	InitCustomLogger()

	InfoLogger.Println("Inside customLogger function")
	WarningLogger.Println("Please look into this function it looks incomplete")
	ErrorLogger.Println("Arggh I QUIT!")
}

func main() {
	// basic()
	// WriteToFile()

	customLogger()
}
