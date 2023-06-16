package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type KeyValue struct {
	key   string
	value string
}

var (
	KeyNotFoundError = errors.New("key was not found")
	HashIndex        = make(map[string]int64)
)

func main() {
	for {
		fmt.Println("Please select an option:")
		fmt.Println("1. insert a new record")
		fmt.Println("2. get a record by a key")
		fmt.Println("3. quit")
		fmt.Println()

		var choice int
		fmt.Print("Your choice: ")
		_, err := fmt.Scanf("%d\n", &choice)
		if err != nil {
			fmt.Println("Invalid option selected! Please try again.")
			fmt.Println()
			continue
		}

		breakTheLoop := false
		switch choice {
		case 1:
			addToDbOption()
		case 2:
			getFromDbOption()
		case 3:
			breakTheLoop = true
		}

		if breakTheLoop {
			break
		}
	}

	fmt.Println("Good bye!")
}

func addToDbOption() {
	var key, value string
	fmt.Print("Please enter key: ")
	_, err := fmt.Scanf("%s\n", &key)
	if err != nil {
		fmt.Println("Something went wrong! Please start over.")
		fmt.Println()
		return
	}

	fmt.Print("Please enter value: ")
	scanner := bufio.NewScanner(os.Stdin)
	ok := scanner.Scan()
	if !ok {
		fmt.Println("Something went wrong! Please start over.")
		fmt.Println()
		return
	}

	value = scanner.Text()

	err = addToDb(key, value)

	if err != nil {
		fmt.Println("Something went wrong! Please start over.")
		fmt.Println()
	}

	fmt.Printf("(%s, %s) is successfully added to the database!\n\n", key, value)
}

func getFromDbOption() {
	var key string
	fmt.Print("Please enter key: ")
	_, err := fmt.Scanf("%s\n", &key)
	if err != nil {
		fmt.Println("Something went wrong! Please start over.")
		fmt.Println()
		return
	}

	value, err := getFromDb(key)

	if err != nil {
		fmt.Printf("%s. Please start over.\n\n", err.Error())
		return
	}

	fmt.Printf("%v\n\n", value)
}

func addToDb(key, value string) error {
	file, err := os.OpenFile("database", os.O_CREATE|os.O_WRONLY, 0755)
	defer file.Close()

	if err != nil {
		return err
	}

	// getFromDb expect to have \n at the end of the last record
	// todo: think if that needs to be changed ?
	s := fmt.Sprintf("%s,%s\n", key, value)

	stat, err := file.Stat()

	if err != nil {
		return err
	}

	var currentPosition int64

	currentPosition, err = file.Seek(stat.Size(), 0)
	if err != nil {
		return err
	}

	_, err = file.WriteString(s)

	HashIndex[key] = currentPosition

	if err != nil {
		return err
	}

	return nil
}

func getFromDb(key string) (value string, err error) {
	file, err := os.OpenFile("database", os.O_RDONLY, 0755)
	defer file.Close()

	offset, ok := HashIndex[key]

	if !ok {
		err = KeyNotFoundError
		return
	}

	_, err = file.Seek(offset, 0)

	if err != nil {
		return
	}

	reader := bufio.NewReader(file)

	line, err := reader.ReadString('\n')

	if err != nil {
		return
	}

	split := strings.Split(line, ",")

	value = split[1]
	value = strings.TrimRight(value, "\n")

	return
}
