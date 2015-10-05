package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/archcra/srtAdjuster/helper"
)

func main() {
	// Parse the arguments
	// ref: https://gobyexample.com/command-line-flags
	pathname := flag.String("filename", "foo", "a string")
	offset := flag.Int("offset", 42, "an int of millisends")
	flag.Parse()

	// pathname := "./testData/tintin-partial-01.txt"
	filename := filepath.Base(*pathname)
	absPath := filepath.Dir(*pathname)
	fmt.Println("the file path is: " + absPath)
	newFilename := absPath + "/" + filename + ".adjusted"

	fileW, err := os.OpenFile(newFilename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer fileW.Close()

	// Read the file
	// ref: http://stackoverflow.com/questions/8757389/reading-file-line-by-line-in-go

	// offset := 3600000

	fileR, err := os.Open(*pathname)
	if err != nil {
		log.Fatal(err)
	}
	defer fileR.Close()

	scanner := bufio.NewScanner(fileR)
	for scanner.Scan() {

		textLine := scanner.Text()
		if isTimeLine(textLine) {
			adjusted := adjustTime(textLine, *offset)
			fmt.Printf("Adjust text is: %s", adjusted)
			if _, err = fileW.WriteString(adjusted + "\n"); err != nil {
				panic(err)
			}
		} else {
			if _, err = fileW.WriteString(textLine + "\n"); err != nil {
				panic(err)
			}
		}
		// fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func isTimeLine(line string) bool {
	fmt.Println("Now line text is: " + line)
	// 00:00:10,000 --> 00:00:20,000
	match, _ := regexp.MatchString("[0-9][0-9]:[0-9][0-9]:[0-9][0-9],[0-9][0-9][0-9] --> [0-9][0-9]:[0-9][0-9]:[0-9][0-9],[0-9][0-9][0-9]", line)
	if match {
		return true
	}
	return false
}
func adjustTime(line string, offset int) string {
	timestamps := strings.Split(line, " --> ")
	timestamp1 := helper.Str2milliseconds(timestamps[0])
	timestamp2 := helper.Str2milliseconds(timestamps[1])
	fmt.Printf("timestamps[1] is: %s and timestamp2 is %d. \n", timestamps[1], timestamp2)

	return helper.Milliseconds2str(timestamp1+offset) + " --> " +
		helper.Milliseconds2str(timestamp2+offset)
}
