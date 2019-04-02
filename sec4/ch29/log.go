package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func LogginSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.SetOutput(multiLogFile)
}

func main() {

	LogginSettings("test.log")
	_, err := os.Open("fugahoge")

	if err != nil {
		log.Fatalln("Exit", err)
	}

	log.Println("loggin!")
	log.Printf("%T %v", "test", "test")
	log.Fatalln("error!!")

	// Fatallnを使用すると exit1 つまりそこで終了する
	fmt.Println("ok!")
}
