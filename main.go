package main

import (
	"fmt"
	"os"

	"marinladovic.com/go/workshop/fileutils"
)

func main() {
	rootPath, _ := os.Getwd()
	filePath := rootPath + "/files/test.txt"

	c, err := fileutils.ReadTextFile(filePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println(c)

		newContent := fmt.Sprintf("Original: %v\nDouble the Original: %v", c, c+c)
		fileutils.WriteToFile(filePath+".output.txt", newContent)
	}
}
