package main

import (
	"fmt"
	"os"
)

type KeyValue struct {
	key   string
	value string
}

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
			var key, value string
			fmt.Print("Please enter key: ")
			_, err := fmt.Scanf("%s\n", &key)
			if err != nil {
				fmt.Println("Something went wrong! Please start over.")
				fmt.Println()
			}

			fmt.Print("Please enter value: ")
			_, err = fmt.Scanf("%s\n", &value)
			if err != nil {
				fmt.Println("Something went wrong! Please start over.")
				fmt.Println()
			}

			//todo: add to the db
			addToDb(key, value)

			fmt.Printf("(%s, %s) is successfully added to the database!\n\n", key, value)
		// todo: prompt for a key-value pair
		case 2:
		//todo: prompt for a key
		case 3:
			breakTheLoop = true
		}

		if breakTheLoop {
			break
		}
	}

	fmt.Println("Good bye!")
}

func addToDb(key, value string) {
	file, err := os.OpenFile("database", os.O_CREATE|os.O_APPEND, 0755)
	defer file.Close()

	if err != nil {
		panic("Can't open database file!\n" + err.Error())
	}

	s := fmt.Sprintf("\n%s,%s", key, value)

	stat, err := file.Stat()

	if err != nil {
		panic("Can't get file info!\n" + err.Error())
	}

	if stat.Size() == 0 {
		s = fmt.Sprintf("%s,%s", key, value)
	}

	_, err = file.WriteString(s)

	if err != nil {
		panic("Can't write to database file!\n" + err.Error())
	}
}
