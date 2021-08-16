package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var SuccessLog *log.Logger
var ErrorLog *log.Logger

func init() {
	if !strings.HasSuffix(os.Args[0], ".test") {
		path, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}

		openLogfile, err := os.OpenFile(path+"/log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

		if err != nil {
			fmt.Println("Error opening file:", err)
			os.Exit(1)
		}

		SuccessLog = log.New(openLogfile, "Success Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)

		ErrorLog = log.New(openLogfile, "Error Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)
	}
}
