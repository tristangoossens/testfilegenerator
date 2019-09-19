package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("fakepackage.go")
	if err != nil {
		log.Fatalf("Error opening file: '%s'", err)
	}
	scanner := bufio.NewScanner(file)
	if !isPackageMain(file, scanner) {
		findFunctions(file, scanner)
	} else {
		findFunctions(file, scanner)
	}
}

func isPackageMain(file io.Reader, scanner *bufio.Scanner) bool {
	var retval bool
	for scanner.Scan() {
		xline := strings.Split(scanner.Text(), " ")
		if xline[0] == "package" && xline[1] == "main" {
			retval = true
			break
		} else {
			retval = false
		}
	}
	return retval
}

func findFunctions(file io.Reader, scanner *bufio.Scanner) {
	for scanner.Scan() {
		xline := strings.Split(scanner.Text(), " ")
		if xline[0] == "func" {
			xword := strings.Split(xline[1], "")
			if strings.ToUpper(xword[0]) == xword[0] {
				fmt.Println("IS upper:", xword[0])
			} else {
				fmt.Println("ISN'T upper:", xword[0])
			}
		}
	}
}
