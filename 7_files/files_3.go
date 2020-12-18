package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// functions
func main() {
	fmt.Println("_files_3()")

	filepath := "./teste.txt"

	// read the file by using buffer io
	// The Scan() method scans token by token, default token is ScanLines
	// ScanWords and other methods available, See http://golang.org/pkg/bufio/
	f, _ := os.Open(filepath)
	// defer is a nifty bit of magic which automatically runs
	// the command before exiting in this case close file
	// good practice to defer right after opening
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// returns text prior to token
		line := scanner.Text()
		fmt.Println(line)
	}

	// check if any errors occurred
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Write file by creating and then writing string by string
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalln("error creating file: ", err)
	}
	defer f.Close()

	for _, str := range []string{"apple", "banana", "carrot"} {
		bytes, err := f.WriteString(str)
		if err != nil {
			log.Fatalln("error writing string: ", err)
		}
		fmt.Printf("Wrote %d bytes to file\n", bytes)
	}
}
