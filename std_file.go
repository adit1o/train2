package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func createNewFile(filePath, content string) error {

	// create or replace file
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)

	// validation
	if err != nil {
		return errors.New("failed create file")
	}

	defer f.Close()
	f.WriteString(content)

	return nil
}

func readFile(filePath string) string {
	output := ""

	f, err := os.OpenFile("adit.md", os.O_RDONLY, 0666)
	defer f.Close()

	if err != nil {
		return err.Error()
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		output += scanner.Text() + "\n"
	}

	return output

}

func appendFile(filePath string, content string) error {

	// create or replace file
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0666)
	defer f.Close()
	if err != nil {
		return err
	}

	if _, err := f.Write([]byte("\n" + content)); err != nil {
		return err
	}

	return nil
}

func main() {
	// content := readFile("adit.md")
	// fmt.Println(content)

	err := appendFile("adit.md", "hello world")
	if err != nil {
		fmt.Println(err)
	}
}
