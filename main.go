package main

import (
	// "archive/zip"
	"fmt"
	"os"
	"zypper/zip"
)

func main() {
	file, err := os.Open("test.zip")
	if err != nil {
		panic("Unable to open test zip file")
	}
	defer file.Close()

	// z, _ := zip.OpenReader("test.zip")
	// for _, f := range z.File {
	// 	fmt.Printf("%+v", f)
	// }
	fmt.Printf("%+v", zip.ReadZip(file))
	fmt.Println("Hello World")
}
