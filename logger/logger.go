package logger

import (
	"fmt"
	"log"
	"os"
)

var SuccessLog *log.Logger
var ErrorLog *log.Logger

func init() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// path, err := os.Getwd()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(path)
	path := "/home/eduardo/projects/unico/"

	openLogfile, err := os.OpenFile(path+"/logs/log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	SuccessLog = log.New(openLogfile, "Success Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)

	ErrorLog = log.New(openLogfile, "Error Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)
}
