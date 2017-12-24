package main

import "log"
import "os"
import "net/http"

func init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	var port = os.Args[1]
	log.Println("port : " + port)

	currentDirectory, currentDirError := os.Getwd()
	if currentDirError != nil {
		log.Fatalln("Current Driectory invalid")
	}

	log.Println(currentDirectory)

	http.Handle("/", http.FileServer(http.Dir(currentDirectory)))
	log.Println("web server start")
	http.ListenAndServe(":"+port, nil)
}
