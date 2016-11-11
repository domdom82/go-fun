package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
)

var (
	txtFileName = "txtFile.txt"
	binFileName = "binFile.bin"
)

// Playground app for file i/o with Go
func main() {

	// Let's see what we can do.
	// Make a simple text file maybe?
	writeTxt()

	// That was easy. Now maybe write a binary file and read it into a buffer?
	writeBin()

	buffer := make([]byte, 512)

	readBin(&buffer)

	// Print the buffer
	fmt.Println("Buffer:")
	fmt.Println(buffer)

	defer func() {
		// Delete all files finally.
		deleteFile(txtFileName)
		deleteFile(binFileName)
	}()

}

func writeTxt() {
	txtFile, err := os.Create(txtFileName)
	if err != nil {
		log.Fatal("Could not create file ", txtFileName)
	}

	// Close file at the end
	defer txtFile.Close()

	// This writer stuff feels like Java. Only there it would be twice as long.
	writer := bufio.NewWriter(txtFile)

	_, err = writer.WriteString("This couldn't be easier!")

	// Flush this text at the end
	defer writer.Flush()

	if err != nil {
		log.Fatal("Couldn't write stuff to ", txtFileName)
	}

	// We should now have a file with default umask
	// but we want to make sure so chmod it
	os.Chmod(txtFileName, 0600)
}

func writeBin() {
	binFile, err := os.Create(binFileName)
	if err != nil {
		log.Fatal("Could not create binary file ", binFileName)
	}

	defer binFile.Close()

	binWriter := bufio.NewWriter(binFile)

	defer binWriter.Flush()

	// write some random stuff
	bytes := 0
	for ; bytes < 512; bytes++ {
		i := byte(255.0 * rand.Float32())
		err = binWriter.WriteByte(i)
		if err != nil {
			log.Fatal("Could not write binary to ", binFileName)
		}
	}

	fmt.Printf("Wrote %d bytes successfully.\n", bytes)

	os.Chmod(binFileName, 0600)
}

func readBin(buffer *[]byte) {
	binFile, err := os.Open(binFileName)
	if err != nil {
		log.Fatal("Could not open binary file ", binFileName)
	}

	defer binFile.Close()

	binReader := bufio.NewReader(binFile)

	// read file into buffer
	bytes, err := binReader.Read(*buffer)

	if err != nil {
		log.Fatal("Could not read binary file ", binFileName, " and read ", bytes, " bytes")
	}

	fmt.Printf("Read %d bytes successfully.\n", bytes)
}

func deleteFile(fileName string) {
	err := os.Remove(fileName)
	if err != nil {
		log.Print("Could not delete file ", fileName)
	}
}
