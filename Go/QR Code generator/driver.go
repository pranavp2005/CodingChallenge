package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// read data to encode from a text file
	fileName := os.	Args[1]
	data := ReadFile(fileName)

	

}


func ReadFile(fileName string) []byte{
 
     
    // The ioutil package contains inbuilt
    // methods like ReadFile that reads the 
    // filename and returns the contents.
    data, err := os.ReadFile("test.txt")
    if err != nil {
        log.Panicf("failed reading data from file: %s", err)
    }
    fmt.Printf("\nFile Name: %s", fileName)
    fmt.Printf("\nSize: %d bytes", len(data))
    fmt.Printf("\nData: %s", data)
 
}