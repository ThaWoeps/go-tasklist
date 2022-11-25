package main

import (
    "fmt"
    "os"
	"log"
	"errors"
)

func checkFileExists(filePath string) bool{ //used to check if tasksFile is there
	_, error := os.Stat(filePath)
	//return !os.IsNotExist(err)
	return !errors.Is(error, os.ErrNotExist)
}

func createFile(){ //create the file we will use to store out tasks
    emptyFile, err := os.Create("tasksFile.txt")
    if err != nil {
        log.Fatal(err)
    }
    log.Println(emptyFile)
    emptyFile.Close()	
}

func main(){
	fmt.Println("Checking if the tasksFile.txt exists")
	var filePath string = "tasksFile.txt"

	isFileExist := checkFileExists(filePath)

	if isFileExist {
		fmt.Println("file exist")
	} else {

		fmt.Println("creating tasksFile.txt")
		createFile()
	}

}