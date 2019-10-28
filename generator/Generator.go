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
	file, err := os.Open("fakepackage.go") // Enter your package file name here!
	if err != nil {
		log.Fatalf("Error opening file: '%s'", err)
	}

	isPackage, packageName := isPackageMain(file)
	if isPackage {
		log.Fatalln("Error: got package main, expected package [package name]")
	} else {
		file.Seek(0, 0)
		// fmt.Println("Slice of functions:", findFunctions(file))
		generateTestFile(packageName, findFunctions(file))
	}
}

func isPackageMain(file io.Reader) (bool, string) {
	scanner := bufio.NewScanner(file)
	var retbool bool
	var retpackage string
	for scanner.Scan() {
		xline := strings.Split(scanner.Text(), " ")
		if len(xline) > 0 {
			if len(xline) != 1 {
				if xline[0] == "package" && xline[1] == "main" {
					retbool = true
					break
				} else if xline[0] == "package" {
					retpackage = xline[1]
					retbool = false
				}
			} else if xline[0] == "package" {
				log.Fatalln("Error: package not declared")
			}
		}
	}
	return retbool, retpackage
}

func findFunctions(file io.Reader) []string {
	scanner := bufio.NewScanner(file)
	var xfunc []string
	for scanner.Scan() {
		xline := strings.Split(scanner.Text(), " ")
		if xline[0] == "func" {
			xword := strings.Split(xline[1], "")
			if strings.ToUpper(xword[0]) == xword[0] {
				// fmt.Println("IS upper:", xline[1])
				newline := strings.Join(strings.Split(xline[1], "(")[:1], "")
				xfunc = append(xfunc, newline)
			} else {
				// fmt.Println("ISN'T upper:", xline[1])
			}
		}
	}
	return xfunc
}

func generateTestFile(packagename string, functions []string) {
	filename := packagename + "_test.go"
	_, err := os.Stat(filename)
	if err != nil {
		generateNewTestFile(filename, packagename, functions)
	} else {
		fmt.Printf("File %s already exists, do you want to overwrite it? WARNING: THIS ACTION IS IRREVERSIBLE\nYes or No? ", filename)
		scanner := bufio.NewScanner(os.Stdin)
	Loop:
		for scanner.Scan() {
			switch strings.ToLower(scanner.Text()) {
			case "yes":
				overwriteOldTestFile(filename, packagename, functions)
				break Loop
			case "no":
				log.Fatalln("Quitting program: chose no")
			default:
				log.Fatalln("Quitting program: invalid command")
			}
		}
	}
}

func generateNewTestFile(filename, packagename string, functions []string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Error creating file \"%s\": %v", filename, err)
	}
	imports := `
        
import (
    "testing"
)
`

	file.WriteString(fmt.Sprintf("package %s %s", packagename, imports))

	for _, v := range functions {
		file.WriteString("\nfunc Test" + v + "(t *testing.Testing){\n\n}\n")
	}

	for _, v := range functions {
		file.WriteString("\nfunc Example" + v + "(){\n\n}\n")
	}

	for _, v := range functions {
		file.WriteString("\nfunc Benchmark" + v + "(b *testing.B){\n\tfor i := 0; i < b.N; i++ {\n\t\t" + v + "()\t//Enter the values that your function needs between the parentheses\n\t}\n}")
	}
	fmt.Printf("Succesfully generated new test file called: %s\n", filename)
}

func overwriteOldTestFile(filename, packagename string, functions []string) {
	file, err := os.OpenFile(filename, os.O_RDWR, 0755)
	if err != nil {
		log.Fatalf("Error opening %s", err)
	}
	imports := `
        
import (
    "testing"
)
`
	file.WriteString(fmt.Sprintf("package %s %s", packagename, imports))

	for _, v := range functions {
		file.WriteString("\nfunc Test" + v + "(t *testing.Testing){\n\n}\n")
	}

	for _, v := range functions {
		file.WriteString("\nfunc Example" + v + "(){\n\n}\n")
	}

	for _, v := range functions {
		file.WriteString("\nfunc Benchmark" + v + "(b *testing.B){\n\tfor i := 0; i < b.N; i++ {\n\t\t" + v + "()\t//Enter the values that your function needs between the parentheses\n\t}\n}")
	}
	fmt.Printf("Succesfully generated new test file called: %s\n", filename)
}
