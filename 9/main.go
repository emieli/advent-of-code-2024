package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// https://adventofcode.com/2024/day/9
// Disk map shows the layout of files and free space on disk. Digits alternatve between
// length of the file and length of free space.
// A disk map of 12345 gives the following layout on disk: 0..111....22222
//
//	1 = 0 (first file ID, one block long)
//	2 = .. (two blocks of free space)
//	3 = 111 (three blocks of file ID 1)
//	4 = .... (four blocks of free space)
//	5 = 22222 (five blocks of file ID 3)
//
// Move file blocks one at a time from end of the disk to the leftmost free space block.
// This process turns the output into "022111222" which is much short than "0..111....22222"
func main() {

	diskMap, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	diskLayout := getDiskLayout(diskMap)
	fmt.Println(diskLayout)

}

func getDiskLayout(diskMap []byte) string {

	var diskLayout string
	var fileIdCounter int

	for i, byte := range diskMap {

		var character string
		var output string

		length, err := strconv.Atoi(string(byte))
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(length)

		if i%2 == 0 {
			// Even number, this is a file
			character = strconv.Itoa(fileIdCounter)
			fileIdCounter++
		} else {
			// Odd number, Free space
			character = "."
		}

		for l := 0; l < length; l++ {
			output += character
		}
		diskLayout += output

	}

	return diskLayout
}
